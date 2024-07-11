package types

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prysmaticlabs/prysm/v5/api/server/structs"
)

// GeneralSideCar is a general sidecar struct for both BSC and ETH
type GeneralSideCar struct {
	structs.Sidecar
	TxIndex int64  `json:"tx_index,omitempty"`
	TxHash  string `json:"tx_hash,omitempty"`
}

type Block struct {
	Header *types.Header `json:"-"`
	Body   *types.Body   `json:"-"`
}
