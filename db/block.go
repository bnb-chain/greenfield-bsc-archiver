package db

import "github.com/ethereum/go-ethereum/common"

type Status int

const (
	Processed Status = 0
	Verified  Status = 1 // each block's blobs will be verified by the post-verification process
	Skipped   Status = 2
)

type Block struct {
	Id          int64
	BlockHash   common.Hash `gorm:"NOT NULL"`
	BlockNumber uint64      `gorm:"NOT NULL;index:idx_block_number"`

	BundleName string `gorm:"NOT NULL"`
	Status     Status `gorm:"index:idx_block_status"`
}

func (*Block) TableName() string {
	return "block"
}
