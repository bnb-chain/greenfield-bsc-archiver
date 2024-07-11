package syncer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cosmossdk.io/math"
	"gorm.io/gorm"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"greeenfield-bsc-archiver/config"
	"greeenfield-bsc-archiver/db"
	"greeenfield-bsc-archiver/external"
	"greeenfield-bsc-archiver/external/cmn"
	"greeenfield-bsc-archiver/logging"
	"greeenfield-bsc-archiver/metrics"
	"greeenfield-bsc-archiver/types"
)

const (
	BundleStatusFinalized      = 1
	BundleStatusCreatedOnChain = 2
	BundleStatusSealedOnChain  = 3

	LoopSleepTime = 10 * time.Millisecond
	BSCPauseTime  = 3 * time.Second

	ETHPauseTime         = 90 * time.Second
	RPCTimeout           = 20 * time.Second
	MonitorQuotaInterval = 5 * time.Minute
)

type curBundleDetail struct {
	name            string
	startBlockID    uint64
	finalizeBlockID uint64
}

type BlockIndexer struct {
	blockDao     db.BlockDao
	client       external.IClient
	bundleClient *cmn.BundleClient
	chainClient  *cmn.ChainClient
	config       *config.SyncerConfig
	bundleDetail *curBundleDetail
	spClient     *cmn.SPClient
	params       *cmn.VersionedParams
}

func NewBlockIndexer(
	blockDao db.BlockDao,
	cfg *config.SyncerConfig,
) *BlockIndexer {
	pkBz, err := hex.DecodeString(cfg.PrivateKey)
	if err != nil {
		panic(err)
	}
	bundleClient, err := cmn.NewBundleClient(cfg.BundleServiceEndpoints[0], cmn.WithPrivateKey(pkBz))
	if err != nil {
		panic(err)
	}
	chainClient, err := cmn.NewChainClient(cfg.GnfdRpcAddr)
	if err != nil {
		panic(err)
	}
	bs := &BlockIndexer{
		blockDao:     blockDao,
		bundleClient: bundleClient,
		chainClient:  chainClient,
		config:       cfg,
	}
	bs.client = external.NewClient(cfg)
	if cfg.MetricsConfig.Enable && len(cfg.MetricsConfig.SPEndpoint) > 0 {
		spClient, err := cmn.NewSPClient(cfg.MetricsConfig.SPEndpoint)
		if err != nil {
			panic(err)
		}
		bs.spClient = spClient
	}
	return bs
}

func (s *BlockIndexer) StartLoop() {
	go func() {
		// nextBlockID defines the block number (BSC)
		nextBlockID, err := s.getNextBlockNum()
		if err != nil {
			panic(err)
		}
		err = s.LoadProgressAndResume(nextBlockID)
		if err != nil {
			panic(err)
		}
		syncTicker := time.NewTicker(LoopSleepTime)
		for range syncTicker.C {
			if err = s.sync(); err != nil {
				logging.Logger.Error(err)
				continue
			}
		}
	}()
	go func() {
		verifyTicket := time.NewTicker(LoopSleepTime)
		for range verifyTicket.C {
			if err := s.verify(); err != nil {
				logging.Logger.Error(err)
				continue
			}
		}
	}()
	go s.monitorQuota()
}

func (s *BlockIndexer) sync() error {
	var (
		blockID uint64
		err     error
		block   *ethtypes.Block
	)
	blockID, err = s.getNextBlockNum()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()

	finalizedBlockNum, err := s.client.GetFinalizedBlockNum(context.Background())
	if err != nil {
		return err
	}
	if int64(blockID) >= int64(finalizedBlockNum) {
		time.Sleep(BSCPauseTime)
		return nil
	}

	ctx, cancel = context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	block, err = s.client.BlockByNumber(ctx, math.NewUint(blockID).BigInt())
	if err != nil {
		return err
	}

	// syncer only store block header & body
	blockInfo := &types.Block{
		Header: block.Header(),
		Body:   block.Body(),
	}

	bundleName := s.bundleDetail.name
	err = s.process(bundleName, blockID, blockInfo)
	if err != nil {
		return err
	}

	blockToSave, err := s.toBlock(block, blockID, bundleName)
	if err != nil {
		return err
	}

	err = s.blockDao.SaveBlock(blockToSave)
	if err != nil {
		logging.Logger.Errorf("failed to save block(h=%d), err=%s", blockToSave.BlockNumber, err.Error())
		return err
	}
	metrics.SyncedBlockIDGauge.Set(float64(blockID))
	logging.Logger.Infof("saved block(block_id=%d) to DB \n", blockID)
	return nil
}

func (s *BlockIndexer) process(bundleName string, blockID uint64, block *types.Block) error {
	var err error
	// create a new bundle in local.
	if blockID == s.bundleDetail.startBlockID {
		if err = s.createLocalBundleDir(); err != nil {
			logging.Logger.Errorf("failed to create local bundle dir, bundle=%s, err=%s", bundleName, err.Error())
			return err
		}
	}

	if err = s.writeBlockToFile(blockID, bundleName, block); err != nil {
		return err
	}
	if blockID == s.bundleDetail.finalizeBlockID {
		err = s.finalizeCurBundle(bundleName)
		if err != nil {
			return err
		}
		logging.Logger.Infof("finalized bundle, bundle_name=%s, bucket_name=%s\n", bundleName, s.getBucketName())
		// init next bundle
		startBlockID := blockID + 1
		endBlockID := blockID + s.getCreateBundleInterval()
		s.bundleDetail = &curBundleDetail{
			name:            types.GetBundleName(startBlockID, endBlockID),
			startBlockID:    startBlockID,
			finalizeBlockID: endBlockID,
		}
	}
	return nil
}

func (s *BlockIndexer) getBucketName() string {
	return s.config.BucketName
}

func (s *BlockIndexer) getCreateBundleInterval() uint64 {
	return s.config.GetCreateBundleInterval()
}

func (s *BlockIndexer) getNextBlockNum() (uint64, error) {
	latestProcessedBlock, err := s.blockDao.GetLatestProcessedBlock()
	if err != nil {
		return 0, fmt.Errorf("failed to get latest polled block from db, error: %s", err.Error())
	}
	latestPolledBlockNumber := latestProcessedBlock.BlockNumber
	nextBlockID := s.config.StartBlock
	if nextBlockID <= latestPolledBlockNumber {
		nextBlockID = latestPolledBlockNumber + 1
	}
	return nextBlockID, nil
}

// createLocalBundleDir creates an empty dir to hold block files among a range of blocks, the block info in this dir will be assembled into a bundle and uploaded to bundle service
func (s *BlockIndexer) createLocalBundleDir() error {
	bundleName := s.bundleDetail.name
	_, err := os.Stat(s.getBundleDir(bundleName))
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(s.getBundleDir(bundleName)), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return s.blockDao.CreateBundle(
		&db.Bundle{
			Name:        s.bundleDetail.name,
			Status:      db.Finalizing,
			CreatedTime: time.Now().Unix(),
		})
}
func (s *BlockIndexer) finalizeBundle(bundleName, bundleDir, bundleFilePath string) error {
	err := s.bundleClient.UploadAndFinalizeBundle(bundleName, s.getBucketName(), bundleDir, bundleFilePath)
	if err != nil {
		if !strings.Contains(err.Error(), "Object exists") && !strings.Contains(err.Error(), "empty bundle") {
			return err
		}
	}
	err = os.RemoveAll(bundleDir)
	if err != nil {
		return err
	}
	err = os.Remove(bundleFilePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return s.blockDao.UpdateBundleStatus(bundleName, db.Finalized)
}

func (s *BlockIndexer) finalizeCurBundle(bundleName string) error {
	return s.finalizeBundle(bundleName, s.getBundleDir(bundleName), s.getBundleFilePath(bundleName))
}

func (s *BlockIndexer) writeBlockToFile(blockNumber uint64, bundleName string, block *types.Block) error {
	blockName := types.GetBlockName(blockNumber)
	file, err := os.Create(s.getBlockPath(bundleName, blockName))
	if err != nil {
		logging.Logger.Errorf("failed to create file, err=%s", err.Error())
		return err
	}
	defer file.Close()

	blockJson, err := json.Marshal(block)
	if err != nil {
		logging.Logger.Errorf("failed to marshal block to JSON", "err", err.Error())
		return err
	}

	_, err = file.Write(blockJson)
	if err != nil {
		logging.Logger.Errorf("failed to write JSON to file", "err", err.Error())
		return err
	}
	return nil
}

func (s *BlockIndexer) getBundleDir(bundleName string) string {
	return fmt.Sprintf("%s/%s/", s.config.TempDir, bundleName)
}

func (s *BlockIndexer) getBlockPath(bundleName, blockName string) string {
	return fmt.Sprintf("%s/%s/%s", s.config.TempDir, bundleName, blockName)
}

func (s *BlockIndexer) getBundleFilePath(bundleName string) string {
	return fmt.Sprintf("%s/%s.bundle", s.config.TempDir, bundleName)
}

func (s *BlockIndexer) LoadProgressAndResume(nextBlockID uint64) error {
	var (
		startBlockID uint64
		endBlockID   uint64
		err          error
	)
	finalizingBundle, err := s.blockDao.GetLatestFinalizingBundle()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		// There is no pending(finalizing) bundle, start a new bundle. e.g. a bundle includes
		// block 0-9 when the block interval is config to 10
		startBlockID = nextBlockID
		endBlockID = nextBlockID + s.getCreateBundleInterval() - 1
	} else {
		// resume
		startBlockID, endBlockID, err = types.ParseBundleName(finalizingBundle.Name)
		if err != nil {
			return err
		}

		// might no longer need to process the bundle even-thought it is not finalized if the user set the config to skip it.
		if nextBlockID > endBlockID {
			err = s.blockDao.UpdateBlocksStatus(startBlockID, endBlockID, db.Skipped)
			if err != nil {
				logging.Logger.Errorf("failed to update blocks status, startBlockID=%d, endBlockID=%d", startBlockID, endBlockID)
				return err
			}
			logging.Logger.Infof("the config block number %d is larger than the recorded bundle end block %d, will resume from the config block", nextBlockID, endBlockID)
			if err = s.blockDao.UpdateBundleStatus(finalizingBundle.Name, db.Deprecated); err != nil {
				return err
			}
			startBlockID = nextBlockID
			endBlockID = nextBlockID + s.getCreateBundleInterval() - 1
		}

	}
	s.bundleDetail = &curBundleDetail{
		name:            types.GetBundleName(startBlockID, endBlockID),
		startBlockID:    startBlockID,
		finalizeBlockID: endBlockID,
	}
	return nil
}

func (s *BlockIndexer) toBlock(block *ethtypes.Block, blockNumber uint64, bundleName string) (*db.Block, error) {
	var (
		blockReturn *db.Block
	)

	blockReturn = &db.Block{
		BlockHash:   block.Hash(),
		BlockNumber: blockNumber,
		BundleName:  bundleName,
	}
	return blockReturn, nil
}

func (s *BlockIndexer) BSCChain() bool {
	return s.config.Chain == config.BSC
}

func (s *BlockIndexer) GetParams() (*cmn.VersionedParams, error) {
	if s.params == nil {
		params, err := s.chainClient.GetParams(context.Background())
		if err != nil {
			logging.Logger.Errorf("failed to get params, err=%s", err.Error())
			return nil, err
		}
		s.params = params
	}
	return s.params, nil

}
