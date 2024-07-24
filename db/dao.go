package db

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type BlockDao interface {
	BlockDB
	BundleDB
}

type BlockSvcDB struct {
	db *gorm.DB
}

func NewBlockSvcDB(db *gorm.DB) BlockDao {
	return &BlockSvcDB{
		db,
	}
}

type BlockDB interface {
	GetBlock(blockNumber uint64) (*Block, error)
	GetBlockByHash(blockHash common.Hash) (*Block, error)
	GetBlockByRoot(root string) (*Block, error)
	GetLatestProcessedBlock() (*Block, error)
	GetEarliestUnverifiedBlock() (*Block, error)
	GetEarliestVerifiedBlock() (*Block, error)
	UpdateBlockStatus(block uint64, status Status) error
	UpdateBlocksStatus(startBlock, endBlock uint64, status Status) error
	SaveBlock(block *Block) error
}

func (d *BlockSvcDB) GetBlock(blockNumber uint64) (*Block, error) {
	block := Block{}
	err := d.db.Model(Block{}).Where("block_number = ?", blockNumber).Take(&block).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (d *BlockSvcDB) GetBlockByHash(blockHash common.Hash) (*Block, error) {
	block := Block{}
	err := d.db.Model(Block{}).Where("block_hash = ?", blockHash).Take(&block).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (d *BlockSvcDB) GetBlockByRoot(root string) (*Block, error) {
	block := Block{}
	err := d.db.Model(Block{}).Where("root = ?", root).Take(&block).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &block, nil
}

func (d *BlockSvcDB) GetLatestProcessedBlock() (*Block, error) {
	block := Block{}
	err := d.db.Model(Block{}).Order("block_number desc").Take(&block).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (d *BlockSvcDB) GetEarliestUnverifiedBlock() (*Block, error) {
	block := Block{}
	err := d.db.Model(Block{}).Where("status = ?", Processed).Order("block_number asc").Take(&block).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (d *BlockSvcDB) GetEarliestVerifiedBlock() (*Block, error) {
	block := Block{}
	err := d.db.Model(Block{}).Where("status = ?", Verified).Order("block_number desc").Take(&block).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (d *BlockSvcDB) UpdateBlockStatus(blockNumber uint64, status Status) error {
	return d.db.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(Block{}).Where("block_number = ?", blockNumber).Updates(
			Block{Status: status}).Error
	})
}

func (d *BlockSvcDB) UpdateBlocksStatus(startBlock, endBlock uint64, status Status) error {
	return d.db.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(Block{}).Where("block_number >= ? and block_number <= ?", startBlock, endBlock).Updates(
			Block{Status: status}).Error
	})
}

type BundleDB interface {
	GetBundle(name string) (*Bundle, error)
	GetLatestFinalizingBundle() (*Bundle, error)
	CreateBundle(*Bundle) error
	UpdateBundleStatus(bundleName string, status InnerBundleStatus) error
}

func (d *BlockSvcDB) GetBundle(name string) (*Bundle, error) {
	bundle := Bundle{}
	err := d.db.Model(Bundle{}).Where("name = ?", name).Take(&bundle).Error
	if err != nil {
		return nil, err
	}
	return &bundle, nil
}

func (d *BlockSvcDB) GetLatestFinalizingBundle() (*Bundle, error) {
	bundle := Bundle{}
	err := d.db.Model(Bundle{}).Where("status = ? and calibrated = false", Finalizing).Order("id desc").Take(&bundle).Error
	if err != nil {
		return nil, err
	}
	return &bundle, nil
}

func (d *BlockSvcDB) CreateBundle(b *Bundle) error {
	return d.db.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Create(b).Error
		if err != nil && MysqlErrCode(err) == ErrDuplicateEntryCode {
			return nil
		}
		return err
	})
}

func (d *BlockSvcDB) UpdateBundleStatus(bundleName string, status InnerBundleStatus) error {
	return d.db.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(Bundle{}).Where("name = ?", bundleName).Updates(
			Bundle{Status: status}).Error
	})
}

func (d *BlockSvcDB) SaveBlock(block *Block) error {
	return d.db.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Save(block).Error
		if err != nil && MysqlErrCode(err) != ErrDuplicateEntryCode {
			return err
		}
		return nil
	})
}

func AutoMigrateDB(db *gorm.DB) {
	var err error
	if err = db.AutoMigrate(&Bundle{}); err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&Block{}); err != nil {
		panic(err)
	}
}
