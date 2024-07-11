package syncer

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cosmossdk.io/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"

	"greeenfield-bsc-archiver/db"
	"greeenfield-bsc-archiver/external/cmn"
	"greeenfield-bsc-archiver/logging"
	"greeenfield-bsc-archiver/metrics"
	"greeenfield-bsc-archiver/types"
	"greeenfield-bsc-archiver/util"
)

const VerifyPauseTime = 90 * time.Second

var (
	ErrVerificationFailed = errors.New("verification failed")
)

// Verify is used to verify the block uploaded to bundle service is indeed in Greenfield, and the integrity.
// In the cases:
//  1. a recorded finalized bundle lost in bundle service
//  2. SP can't seal the object (probably won't seal it anymore)
//  3. verification on a specified block failed
//
// a new bundle should be re-uploaded.
func (s *BlockIndexer) verify() error {
	verifyBlock, err := s.blockDao.GetEarliestUnverifiedBlock()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logging.Logger.Debugf("found no unverified block in DB")
			time.Sleep(VerifyPauseTime)
			return nil
		}
		return err
	}
	bundleName := verifyBlock.BundleName

	// check if the bundle has been submitted to bundle service
	bundle, err := s.blockDao.GetBundle(bundleName)
	if err != nil {
		return err
	}
	if bundle.Status == db.Finalizing {
		logging.Logger.Debugf("the bundle has not been submitted to bundle service yet, bundleName=%s", bundleName)
		time.Sleep(VerifyPauseTime)
		return nil
	}

	bundleStartBlockID, bundleEndBlockID, err := types.ParseBundleName(bundleName)
	if err != nil {
		return err
	}
	verifyBlockID := verifyBlock.BlockNumber

	// validate the bundle info at the start block of a bundle
	if verifyBlockID == bundleStartBlockID || !s.DetailedIntegrityCheckEnabled() {
		// the bundle is recorded finalized in DB, validate the bundle is sealed onchain
		bundleInfo, err := s.bundleClient.GetBundleInfo(s.getBucketName(), bundleName)
		if err != nil {
			if err != cmn.ErrorBundleNotExist {
				logging.Logger.Errorf("failed to get bundle info, bundleName=%s", bundleName)
				return err
			}
			if err = s.blockDao.UpdateBlocksStatus(bundleStartBlockID, bundleEndBlockID, db.Verified); err != nil {
				return err
			}
			if err = s.blockDao.UpdateBundleStatus(bundleName, db.Sealed); err != nil {
				return err
			}
			return nil
		}
		// the bundle is not sealed yet
		if bundleInfo.Status == BundleStatusFinalized || bundleInfo.Status == BundleStatusCreatedOnChain {
			if bundle.CreatedTime > 0 && time.Now().Unix()-bundle.CreatedTime > s.config.GetReUploadBundleThresh() {
				logging.Logger.Infof("the bundle %s is not sealed and exceed the re-upload threshold %d ", bundleName, s.config.GetReUploadBundleThresh())
				return s.reUploadBundle(bundleName)
			}
			return nil
		}
	}

	// if the detailed integrity check is disabled, verify the bundle integrity
	if !s.DetailedIntegrityCheckEnabled() {
		err = s.verifyBundleIntegrity(bundleName, bundleStartBlockID, bundleEndBlockID)
		if err != nil {
			logging.Logger.Errorf("failed to verify bundle integrity, bundleName=%s, err=%s", bundleName, err.Error())
			if errors.Is(err, ErrVerificationFailed) {
				return s.reUploadBundle(bundleName)
			}
			return err
		}
		return nil
	}

	//if verifyBlock.BlobCount == 0 {
	//	if err = s.blockDao.UpdateBlockStatus(verifyBlockID, db.Verified); err != nil {
	//		logging.Logger.Errorf("failed to update block status, block_id=%d err=%s", verifyBlockID, err.Error())
	//		return err
	//	}
	//	if verifyBlockID == bundleEndBlockID {
	//		logging.Logger.Debugf("update bundle status to sealed, name=%s , block_id %d ", bundleName, verifyBlockID)
	//		if err = s.blockDao.UpdateBundleStatus(bundleName, db.Sealed); err != nil {
	//			logging.Logger.Errorf("failed to update bundle status to sealed, name=%s , block_id %d ", bundleName, verifyBlockID)
	//			return err
	//		}
	//	}
	//	return nil
	//}

	// get block from BSC again
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	block, err := s.client.BlockByNumber(ctx, math.NewUint(verifyBlockID).BigInt())
	if err != nil {
		logging.Logger.Errorf("failed to get blob at block_id=%d, err=%s", verifyBlockID, err.Error())
		return err
	}

	// get block meta from DB
	blockMetas, err := s.blockDao.GetBlock(verifyBlockID)
	if err != nil {
		return err
	}

	//if len(blockMetas) != len(block) {
	//	logging.Logger.Errorf("found blob number mismatch at block_id=%d, bundleName=%s, expected=%d, actual=%d", verifyBlockID, bundleName, len(sideCars), len(blobMetas))
	//	return s.reUploadBundle(bundleName)
	//}

	err = s.verifyBlobsAtBlock(verifyBlockID, block, blockMetas, bundleName)
	if err != nil {
		if errors.Is(err, ErrVerificationFailed) {
			return s.reUploadBundle(bundleName)
		}
		return err
	}
	if err = s.blockDao.UpdateBlockStatus(verifyBlockID, db.Verified); err != nil {
		logging.Logger.Errorf("failed to update block status to verified, block_id=%d err=%s", verifyBlockID, err.Error())
		return err
	}
	metrics.VerifiedBlockIDGauge.Set(float64(verifyBlockID))
	if bundleEndBlockID == verifyBlockID {
		logging.Logger.Debugf("update bundle status to sealed, name=%s , block_id=%d ", bundleName, verifyBlockID)
		if err = s.blockDao.UpdateBundleStatus(bundleName, db.Sealed); err != nil {
			logging.Logger.Errorf("failed to update bundle status to sealed, name=%s, block_id %d ", bundleName, verifyBlockID)
			return err
		}
	}
	logging.Logger.Infof("successfully verify at block block_id=%d ", verifyBlockID)
	return nil
}

// verifyBundleIntegrity is used to verify the integrity of a bundle by comparing the checksums of the re-constructed bundle object and the on-chain object.
// If the checksums are not equal, the bundle will be re-uploaded, and the re-uploaded bundle will be verified as well, until the verification is successful.
func (s *BlockIndexer) verifyBundleIntegrity(bundleName string, bundleStartBlockID, bundleEndBlockID uint64) error {
	// recreate the bundle for the block range
	verifyBundleName := bundleName + "_verify"
	_, err := os.Stat(s.getBundleDir(verifyBundleName))
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(s.getBundleDir(verifyBundleName)), os.ModePerm)
		if err != nil {
			return err
		}
	}
	defer os.RemoveAll(s.getBundleDir(verifyBundleName))

	for bi := bundleStartBlockID; bi <= bundleEndBlockID; bi++ {
		logging.Logger.Infof("start to get blob from block_id=%d", bi)
		ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
		defer cancel()
		var block *ethtypes.Block
		block, err = s.client.BlockByNumber(ctx, math.NewUint(bi).BigInt())
		if err != nil {
			return err
		}
		if err = s.writeBlockToFile(bi, verifyBundleName, &types.Block{
			Header: block.Header(),
			Body:   block.Body(),
		}); err != nil {
			return err
		}
	}
	bundleObject, _, err := cmn.BundleObjectFromDirectory(s.getBundleDir(verifyBundleName))
	if err != nil {
		return err
	}
	logging.Logger.Infof("successfully bundle object from dir, name=%s", verifyBundleName)

	storageParams, err := s.GetParams()
	if err != nil {
		return err
	}
	maxSegSize, err := util.StringToInt64(storageParams.MaxSegmentSize)
	if err != nil {
		return err
	}
	// compute the integrity hash
	expectCheckSums, _, err := util.ComputeIntegrityHashSerial(bundleObject, maxSegSize, storageParams.RedundantDataChunkNum, storageParams.RedundantParityChunkNum)
	if err != nil {
		return err
	}
	// get object from chain
	onChainBundleObject, err := s.chainClient.GetObjectMeta(context.Background(), s.getBucketName(), bundleName)
	if err != nil {
		logging.Logger.Errorf("failed to get object from chain, bucketName = %s, bundleName=%s, err=%s", s.getBucketName(), bundleName, err.Error())
		return err
	}
	if len(expectCheckSums) != len(onChainBundleObject.Checksums) {
		logging.Logger.Errorf("found checksum number mismatch")
		return ErrVerificationFailed
	}
	// compare the checksum
	for i, expectCheckSum := range expectCheckSums {
		encodedChecksum := base64.StdEncoding.EncodeToString(expectCheckSum)
		if !strings.EqualFold(encodedChecksum, onChainBundleObject.Checksums[i]) {
			logging.Logger.Errorf("found checksum mismatch")
			return ErrVerificationFailed
		}
	}
	// update the status
	if err = s.blockDao.UpdateBlocksStatus(bundleStartBlockID, bundleEndBlockID, db.Verified); err != nil {
		return err
	}
	metrics.VerifiedBlockIDGauge.Set(float64(bundleEndBlockID))
	if err = s.blockDao.UpdateBundleStatus(bundleName, db.Sealed); err != nil {
		return err
	}
	logging.Logger.Infof("successfully verify bundle=%s, start_block_id=%d, end_block_id =%d ", bundleName, bundleStartBlockID, bundleEndBlockID)
	return nil
}

func (s *BlockIndexer) verifyBlobsAtBlock(blockID uint64, block *ethtypes.Block, blockMetas *db.Block, bundleName string) error {
	// get block from bundle service
	blockFromBundle, err := s.bundleClient.GetObject(s.getBucketName(), bundleName, types.GetBlockName(blockID))
	if err != nil {
		if errors.Is(err, cmn.ErrorBundleObjectNotExist) {
			logging.Logger.Errorf("the bundle object not found in bundle service, object=%s", types.GetBlockName(blockID))
			return ErrVerificationFailed
		}
		return err
	}

	if block.Hash() != blockMetas.BlockHash {
		logging.Logger.Errorf("found db block mismatch")
		return ErrVerificationFailed
	}

	actualBlockHash, err := util.GenerateHash(blockFromBundle)
	if err != nil {
		return err
	}

	blockInfo := &types.Block{
		Header: block.Header(),
		Body:   block.Body(),
	}
	blockJson, err := json.Marshal(blockInfo)
	if err != nil {
		logging.Logger.Errorf("failed to marshal block to JSON", "err", err.Error())
		return err
	}
	expectedBlockHash, err := util.GenerateHash(string(blockJson))
	if err != nil {
		return err
	}
	if !bytes.Equal(actualBlockHash, expectedBlockHash) {
		logging.Logger.Errorf("found block mismatch")
		return ErrVerificationFailed
	}

	return nil
}

func (s *BlockIndexer) reUploadBundle(bundleName string) error {
	if err := s.blockDao.UpdateBundleStatus(bundleName, db.Deprecated); err != nil {
		return err
	}

	newBundleName := bundleName + "_calibrated_" + util.Int64ToString(time.Now().Unix())
	startBlockID, endBlockID, err := types.ParseBundleName(bundleName)
	if err != nil {
		return err
	}
	logging.Logger.Infof("creating new calibrated bundle %s", newBundleName)

	_, err = os.Stat(s.getBundleDir(newBundleName))
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(s.getBundleDir(newBundleName)), os.ModePerm)
		if err != nil {
			return err
		}
	}
	if err = s.blockDao.CreateBundle(&db.Bundle{
		Name:        newBundleName,
		Status:      db.Finalizing,
		Calibrated:  true,
		CreatedTime: time.Now().Unix(),
	}); err != nil {
		return err
	}
	for bi := startBlockID; bi <= endBlockID; bi++ {
		ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
		defer cancel()

		var block *ethtypes.Block
		block, err = s.client.BlockByNumber(ctx, math.NewUint(bi).BigInt())
		if err != nil {
			return err
		}

		if err = s.writeBlockToFile(bi, newBundleName, &types.Block{
			Header: block.Header(),
			Body:   block.Body(),
		}); err != nil {
			return err
		}

		blockMeta, err := s.blockDao.GetBlock(bi)
		if err != nil {
			return err
		}

		blockToSave, err := s.toBlock(block, bi, newBundleName)
		if err != nil {
			return err
		}

		blockToSave.Id = blockMeta.Id
		err = s.blockDao.SaveBlock(blockToSave)
		if err != nil {
			logging.Logger.Errorf("failed to save block(h=%d), err=%s", blockToSave.BlockNumber, err.Error())
			return err
		}

		logging.Logger.Infof("save calibrated block(block_id=%d)", bi)
	}
	if err = s.finalizeBundle(newBundleName, s.getBundleDir(newBundleName), s.getBundleFilePath(newBundleName)); err != nil {
		logging.Logger.Errorf("failed to finalized bundle, name=%s, err=%s", newBundleName, err.Error())
		return err
	}
	return nil
}

// DetailedIntegrityCheckEnabled returns whether the detailed integrity check on individual block is enabled, otherwise the
// integrity check will be done on the bundle level.
func (s *BlockIndexer) DetailedIntegrityCheckEnabled() bool {
	return s.config.EnableIndivBlockVerification
}
