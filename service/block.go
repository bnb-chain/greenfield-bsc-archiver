package service

import (
	"greeenfield-bsc-archiver/cache"
	"greeenfield-bsc-archiver/config"
	"greeenfield-bsc-archiver/db"
	"greeenfield-bsc-archiver/external/cmn"
	"greeenfield-bsc-archiver/models"
)

const prefixHex = "0x"

type Block interface {
	GetBlockByBlockNumber(block uint64) ([]*models.Sidecar, error)
	GetBundledBlockByBlockNumber(block uint64, indices []int64) ([]*models.Sidecar, error)
}

type BlockService struct {
	blobDB       db.BlockDao
	bundleClient *cmn.BundleClient
	cacheService cache.Cache
	cfg          *config.ServerConfig
}

func NewBlockService(blobDB db.BlockDao, bundleClient *cmn.BundleClient, cache cache.Cache, config *config.ServerConfig) Block {
	return &BlockService{
		blobDB:       blobDB,
		bundleClient: bundleClient,
		cacheService: cache,
		cfg:          config,
	}
}

func (b BlockService) GetBlockByBlockNumber(blockNumber uint64) ([]*models.Sidecar, error) {
	//block, err := b.blobDB.GetBlock(blockNumber)
	//if err != nil {
	//	return nil, err
	//}
	////return b.GetBlobSidecarsByBlockNumOrSlot(block.Slot, indices)
	//bundleObject, err := b.bundleClient.GetObject(b.cfg.BucketName, block.BundleName, types.GetBlockName(blockNumber))
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (b BlockService) GetBundledBlockByBlockNumber(block uint64, indices []int64) ([]*models.Sidecar, error) {
	return nil, nil
}

//func (b BlobService) GetBlobSidecarsByBlockNumOrSlot(blockNumOrSlot uint64, indices []int64) ([]*models.Sidecar, error) {
//	var err error
//	blobs, found := b.cacheService.Get(util.Uint64ToString(blockNumOrSlot))
//	if found {
//		blobsFound := blobs.([]*models.Sidecar)
//		if len(indices) != 0 {
//			blobReturn := make([]*models.Sidecar, 0)
//			for _, idx := range indices {
//				if int(idx) >= len(blobsFound) {
//					return nil, fmt.Errorf("index %d out of bound, only %d blob at block %d", idx, len(blobsFound), blockNumOrSlot)
//				}
//				blobReturn = append(blobReturn, blobsFound[idx])
//			}
//			return blobReturn, nil
//		}
//		return blobsFound, nil
//	}
//
//	block, err := b.blobDB.GetBlock(blockNumOrSlot)
//	if err != nil {
//		return nil, err
//	}
//
//	var blobMetas []*db.Blob
//	if len(indices) == 0 {
//		blobMetas, err = b.blobDB.GetBlobByBlockID(blockNumOrSlot)
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		blobMetas, err = b.blobDB.GetBlobByBlockIDAndIndices(blockNumOrSlot, indices)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	sideCars := make([]*models.Sidecar, 0)
//	for _, meta := range blobMetas {
//		bundleObject, err := b.bundleClient.GetObject(b.cfg.BucketName, block.BundleName, meta.Name)
//		if err != nil {
//			return nil, err
//		}
//		var header *models.SidecarSignedBlockHeader
//		if b.cfg.Chain == config.ETH {
//			header = &models.SidecarSignedBlockHeader{
//				Message: &models.SidecarSignedBlockHeaderMessage{
//					BodyRoot:      fmt.Sprintf("%s%s", prefixHex, block.BodyRoot),
//					ParentRoot:    fmt.Sprintf("%s%s", prefixHex, block.ParentRoot),
//					StateRoot:     fmt.Sprintf("%s%s", prefixHex, block.StateRoot),
//					ProposerIndex: util.Uint64ToString(block.ProposerIndex),
//					Slot:          util.Uint64ToString(block.Slot),
//				},
//				Signature: fmt.Sprintf("%s%s", prefixHex, block.Signature),
//			}
//		}
//		sideCars = append(sideCars,
//			&models.Sidecar{
//				Blob:                        bundleObject,
//				Index:                       util.Int64ToString(int64(meta.Idx)),
//				KzgCommitmentInclusionProof: util.SplitByComma(meta.CommitmentInclusionProof),
//				KzgCommitment:               meta.KzgCommitment,
//				KzgProof:                    meta.KzgProof,
//				SignedBlockHeader:           header,
//				TxIndex:                     int64(meta.TxIndex),
//				TxHash:                      meta.TxHash,
//			})
//	}
//
//	// cache all blobs at a specified slot
//	if len(indices) == 0 {
//		b.cacheService.Set(util.Uint64ToString(blockNumOrSlot), sideCars)
//	}
//	return sideCars, nil
//}
