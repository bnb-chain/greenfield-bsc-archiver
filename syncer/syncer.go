package syncer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"cosmossdk.io/math"
	"gorm.io/gorm"
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

	DBRetryTimes = 3

	LoopSleepTime = 10 * time.Millisecond
	WaitSleepTime = 100 * time.Millisecond
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
	blockDao             db.BlockDao
	client               external.IClient
	bundleClient         *cmn.BundleClient
	chainClient          *cmn.ChainClient
	config               *config.SyncerConfig
	bundleDetail         *curBundleDetail
	spClient             *cmn.SPClient
	params               *cmn.VersionedParams
	blocks               map[uint64]*types.RpcBlock
	blocksLock           sync.Mutex
	blockHeightLock      sync.Mutex
	indexBlockHeight     uint64
	processorBlockHeight uint64
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
		blocks:       make(map[uint64]*types.RpcBlock),
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

func (b *BlockIndexer) StartConcurrentSync() {
	var wg sync.WaitGroup
	for i := 0; i < b.config.ConcurrencyLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				var blockID uint64
				b.blockHeightLock.Lock()
				blockID = b.indexBlockHeight
				b.indexBlockHeight++
				b.blockHeightLock.Unlock()

				for {
					err := b.sync(blockID)
					if err != nil {
						logging.Logger.Error("Failed to sync: ", err)
						time.Sleep(time.Second)
						continue
					}
					break
				}
				logging.Logger.Infof("Successfully fetched and incremented block ID to %d", blockID)
			}
		}()
	}
	wg.Wait()
}

func (b *BlockIndexer) StartLoop() {
	go func() {
		// nextBlockID defines the block number (BSC)
		var nextBlockID uint64
		var err error
		for retries := 0; retries < DBRetryTimes; retries++ {
			nextBlockID, err = b.getNextBlockNum()
			if err == nil {
				break
			}
			logging.Logger.Errorf("Failed to get next block number: %v. Retrying... (%d/3)", err, retries+1)
			time.Sleep(time.Second * 2) // Wait for 2 seconds before retrying
		}
		if err != nil {
			panic(err)
		}
		b.processorBlockHeight = nextBlockID
		b.indexBlockHeight = nextBlockID
		err = b.LoadProgressAndResume(nextBlockID)
		if err != nil {
			panic(err)
		}
		for {
			b.StartConcurrentSync()
			time.Sleep(LoopSleepTime)
		}
	}()
	go func() {
		syncTicker := time.NewTicker(LoopSleepTime)
		for range syncTicker.C {
			blockID := b.processorBlockHeight
			b.blocksLock.Lock()
			block, exists := b.blocks[blockID]
			b.blocksLock.Unlock()
			for !exists {
				time.Sleep(WaitSleepTime)
				block, exists = b.blocks[blockID]
			}
			b.blocksLock.Lock()
			delete(b.blocks, blockID)
			b.blocksLock.Unlock()
			b.processorBlockHeight++
			var err error
			for err != nil {
				if err = b.process(b.bundleDetail.name, blockID, block); err != nil {
					logging.Logger.Error(err)
					continue
				}
			}

		}
	}()
	go func() {
		verifyTicket := time.NewTicker(LoopSleepTime)
		for range verifyTicket.C {
			if err := b.verify(); err != nil {
				logging.Logger.Error(err)
				continue
			}
		}
	}()
	go b.monitorQuota()
}

func (b *BlockIndexer) sync(blockID uint64) error {
	var (
		err      error
		rpcBlock *types.RpcBlock
	)
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()

	finalizedBlockNum, err := b.client.GetFinalizedBlockNum(ctx)
	if err != nil {
		return err
	}
	if int64(blockID) >= int64(finalizedBlockNum) {
		time.Sleep(BSCPauseTime)
		return nil
	}

	ctx, cancel = context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()

	rpcBlock, err = b.client.GetBlockByNumber(ctx, math.NewUint(blockID).BigInt())
	if err != nil {
		return err
	}

	b.blocksLock.Lock()
	b.blocks[blockID] = rpcBlock
	b.blocksLock.Unlock()
	return nil
}

func (b *BlockIndexer) process(bundleName string, blockID uint64, block *types.RpcBlock) error {
	var err error
	// create a new bundle in local.
	if blockID == b.bundleDetail.startBlockID {
		if err = b.createLocalBundleDir(); err != nil {
			logging.Logger.Errorf("failed to create local bundle dir, bundle=%s, err=%s", bundleName, err.Error())
			return err
		}
	}

	if err = b.writeBlockToFile(blockID, bundleName, block); err != nil {
		return err
	}
	if blockID == b.bundleDetail.finalizeBlockID {
		err = b.finalizeCurBundle(bundleName)
		if err != nil {
			return err
		}
		logging.Logger.Infof("finalized bundle, bundle_name=%s, bucket_name=%s\n", bundleName, b.getBucketName())
		// init next bundle
		startBlockID := blockID + 1
		endBlockID := blockID + b.getCreateBundleInterval()
		b.bundleDetail = &curBundleDetail{
			name:            types.GetBundleName(startBlockID, endBlockID),
			startBlockID:    startBlockID,
			finalizeBlockID: endBlockID,
		}
	}

	blockToSave, err := b.toBlock(block, blockID, bundleName)
	if err != nil {
		return err
	}

	err = b.blockDao.SaveBlock(blockToSave)
	if err != nil {
		logging.Logger.Errorf("failed to save block(h=%d), err=%s", blockToSave.BlockNumber, err.Error())
		return err
	}
	metrics.SyncedBlockIDGauge.Set(float64(blockID))
	logging.Logger.Infof("saved block(block_id=%d) to DB \n", blockID)
	return nil
}

func (b *BlockIndexer) getBucketName() string {
	return b.config.BucketName
}

func (b *BlockIndexer) getCreateBundleInterval() uint64 {
	return b.config.GetCreateBundleInterval()
}

func (b *BlockIndexer) getNextBlockNum() (uint64, error) {
	latestProcessedBlock, err := b.blockDao.GetLatestProcessedBlock()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get latest polled block from db, error: %s", err.Error())
	}

	latestPolledBlockNumber := latestProcessedBlock.BlockNumber
	nextBlockID := b.config.StartBlock
	if nextBlockID <= latestPolledBlockNumber {
		nextBlockID = latestPolledBlockNumber + 1
	}
	return nextBlockID, nil
}

// createLocalBundleDir creates an empty dir to hold block files among a range of blocks, the block info in this dir will be assembled into a bundle and uploaded to bundle service
func (b *BlockIndexer) createLocalBundleDir() error {
	bundleName := b.bundleDetail.name
	_, err := os.Stat(b.getBundleDir(bundleName))
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(b.getBundleDir(bundleName)), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return b.blockDao.CreateBundle(
		&db.Bundle{
			Name:        b.bundleDetail.name,
			Status:      db.Finalizing,
			CreatedTime: time.Now().Unix(),
		})
}

func (b *BlockIndexer) finalizeBundle(bundleName, bundleDir, bundleFilePath string) error {
	err := b.bundleClient.UploadAndFinalizeBundle(bundleName, b.getBucketName(), bundleDir, bundleFilePath)
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
	return b.blockDao.UpdateBundleStatus(bundleName, db.Finalized)
}

func (b *BlockIndexer) finalizeCurBundle(bundleName string) error {
	return b.finalizeBundle(bundleName, b.getBundleDir(bundleName), b.getBundleFilePath(bundleName))
}

func (b *BlockIndexer) writeBlockToFile(blockNumber uint64, bundleName string, block *types.RpcBlock) error {
	blockName := types.GetBlockName(blockNumber)
	file, err := os.Create(b.getBlockPath(bundleName, blockName))
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

func (b *BlockIndexer) getBundleDir(bundleName string) string {
	return fmt.Sprintf("%s/%s/", b.config.TempDir, bundleName)
}

func (b *BlockIndexer) getBlockPath(bundleName, blockName string) string {
	return fmt.Sprintf("%s/%s/%s", b.config.TempDir, bundleName, blockName)
}

func (b *BlockIndexer) getBundleFilePath(bundleName string) string {
	return fmt.Sprintf("%s/%s.bundle", b.config.TempDir, bundleName)
}

func (b *BlockIndexer) LoadProgressAndResume(nextBlockID uint64) error {
	var (
		startBlockID uint64
		endBlockID   uint64
		err          error
	)
	finalizingBundle, err := b.blockDao.GetLatestFinalizingBundle()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		// There is no pending(finalizing) bundle, start a new bundle. e.g. a bundle includes
		// block 0-9 when the block interval is config to 10
		startBlockID = nextBlockID
		endBlockID = nextBlockID + b.getCreateBundleInterval() - 1
	} else {
		// resume
		startBlockID, endBlockID, err = types.ParseBundleName(finalizingBundle.Name)
		if err != nil {
			return err
		}

		// might no longer need to process the bundle even-thought it is not finalized if the user set the config to skip it.
		if nextBlockID > endBlockID {
			err = b.blockDao.UpdateBlocksStatus(startBlockID, endBlockID, db.Skipped)
			if err != nil {
				logging.Logger.Errorf("failed to update blocks status, startBlockID=%d, endBlockID=%d", startBlockID, endBlockID)
				return err
			}
			logging.Logger.Infof("the config block number %d is larger than the recorded bundle end block %d, will resume from the config block", nextBlockID, endBlockID)
			if err = b.blockDao.UpdateBundleStatus(finalizingBundle.Name, db.Deprecated); err != nil {
				return err
			}
			startBlockID = nextBlockID
			endBlockID = nextBlockID + b.getCreateBundleInterval() - 1
		}

	}
	b.bundleDetail = &curBundleDetail{
		name:            types.GetBundleName(startBlockID, endBlockID),
		startBlockID:    startBlockID,
		finalizeBlockID: endBlockID,
	}
	return nil
}

func (b *BlockIndexer) toBlock(block *types.RpcBlock, blockNumber uint64, bundleName string) (*db.Block, error) {
	blockReturn := &db.Block{
		BlockHash:   block.Hash,
		BlockNumber: blockNumber,
		BundleName:  bundleName,
	}
	return blockReturn, nil
}

func (b *BlockIndexer) BSCChain() bool {
	return b.config.Chain == config.BSC
}

func (b *BlockIndexer) GetParams() (*cmn.VersionedParams, error) {
	if b.params == nil {
		params, err := b.chainClient.GetParams(context.Background())
		if err != nil {
			logging.Logger.Errorf("failed to get params, err=%s", err.Error())
			return nil, err
		}
		b.params = params
	}
	return b.params, nil

}
