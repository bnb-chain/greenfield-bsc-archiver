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
	BlockHash   common.Hash `gorm:"NOT NULL;type:BINARY(32);index:idx_block_hash,prefix_len:32"`
	BlockNumber uint64      `gorm:"NOT NULL;uniqueIndex:idx_block_number;index:idx_block_number_status,priority:2"`

	BundleName string `gorm:"NOT NULL"`
	Status     Status `gorm:"index:idx_block_number_status,priority:1"`
}

func (*Block) TableName() string {
	return "block"
}
