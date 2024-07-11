package service

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"greeenfield-bsc-archiver/cache"
	"greeenfield-bsc-archiver/config"
	"greeenfield-bsc-archiver/db"
	"greeenfield-bsc-archiver/external/cmn"
	"greeenfield-bsc-archiver/models"
	"greeenfield-bsc-archiver/types"
)

type Block interface {
	GetBlockByBlockNumber(blockNumber uint64) (*models.Block, error)
	GetBundledBlockByBlockNumber(blockNumber uint64) ([]*models.Block, error)
	GetBlockByBlockHash(blockHash common.Hash) (*models.Block, error)
	GetLatestBlockNumber() (uint64, error)
	GetBundleStartBlockID(blockNumber uint64) (uint64, error)
}

type BlockService struct {
	blockDB      db.BlockDao
	bundleClient *cmn.BundleClient
	spClient     *cmn.SPClient
	cacheService cache.Cache
	cfg          *config.ServerConfig
}

func NewBlockService(blockDB db.BlockDao, bundleClient *cmn.BundleClient, spClient *cmn.SPClient, cache cache.Cache, config *config.ServerConfig) Block {
	return &BlockService{
		blockDB:      blockDB,
		bundleClient: bundleClient,
		spClient:     spClient,
		cacheService: cache,
		cfg:          config,
	}
}

func (b BlockService) GetBlockByBlockNumber(blockNumber uint64) (*models.Block, error) {
	block, err := b.blockDB.GetBlock(blockNumber)
	if err != nil {
		return nil, err
	}

	bundleObject, err := b.bundleClient.GetObject(b.cfg.BucketName, block.BundleName, types.GetBlockName(blockNumber))
	if err != nil {
		return nil, err
	}
	var blockInfo *models.Block
	err = json.Unmarshal([]byte(bundleObject), &blockInfo)
	return blockInfo, err
}

func (b BlockService) GetBundledBlockByBlockNumber(blockNumber uint64) ([]*models.Block, error) {
	block, err := b.blockDB.GetBlock(blockNumber)
	if err != nil {
		return nil, err
	}

	return b.spClient.GetBundleBlocks(context.Background(), b.cfg.BucketName, block.BundleName)
}

func (b BlockService) GetBlockByBlockHash(blockHash common.Hash) (*models.Block, error) {
	block, err := b.blockDB.GetBlockByHash(blockHash)
	if err != nil {
		return nil, err
	}

	bundleObject, err := b.bundleClient.GetObject(b.cfg.BucketName, block.BundleName, types.GetBlockName(block.BlockNumber))
	if err != nil {
		return nil, err
	}
	var blockInfo *models.Block
	err = json.Unmarshal([]byte(bundleObject), &blockInfo)
	return blockInfo, err
}

func (b BlockService) GetLatestBlockNumber() (uint64, error) {
	block, err := b.blockDB.GetLatestProcessedBlock()
	if err != nil {
		return 0, err
	}
	return block.BlockNumber, err
}

func (b BlockService) GetBundleStartBlockID(blockNumber uint64) (uint64, error) {
	block, err := b.blockDB.GetBlock(blockNumber)
	if err != nil {
		return 0, err
	}
	startBlock, _, err := types.ParseBundleName(block.BundleName)
	if err != nil {
		return 0, err
	}
	return startBlock, nil
}
