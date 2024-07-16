package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"greeenfield-bsc-archiver/models"
	"greeenfield-bsc-archiver/util"
)

type Block struct {
	Hash        common.Hash      `json:"hash"       gencodec:"required"`
	ParentHash  common.Hash      `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address   `json:"miner"`
	Root        common.Hash      `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash      `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int         `json:"difficulty"       gencodec:"required"`
	Number      *big.Int         `json:"number"           gencodec:"required"`
	GasLimit    uint64           `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64           `json:"gasUsed"          gencodec:"required"`
	Time        uint64           `json:"timestamp"        gencodec:"required"`
	Extra       []byte           `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash      `json:"mixHash"`
	Nonce       types.BlockNonce `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash        `json:"parentBeaconBlockRoot" rlp:"optional"`
	Transactions     []*Transaction      `json:"transactions"`
	Uncles           []*Header           `json:"uncles"`
	Withdrawals      []*types.Withdrawal `json:"withdrawals"`

	TotalDifficulty *big.Int `json:"totalDifficulty,omitempty"  gencodec:"required"`
	Size            *big.Int `json:"size,omitempty"  gencodec:"required"`
}

type RealBlock struct {
	Hash        common.Hash      `json:"hash"       gencodec:"required"`
	ParentHash  common.Hash      `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address   `json:"miner"`
	Root        common.Hash      `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash      `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int         `json:"difficulty"       gencodec:"required"`
	Number      *big.Int         `json:"number"           gencodec:"required"`
	GasLimit    uint64           `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64           `json:"gasUsed"          gencodec:"required"`
	Time        uint64           `json:"timestamp"        gencodec:"required"`
	Extra       []byte           `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash      `json:"mixHash"`
	Nonce       types.BlockNonce `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`

	Transactions []*types.Transaction `json:"transactions"`
	Uncles       []*types.Header      `json:"uncles"`
	Withdrawals  []*types.Withdrawal  `json:"withdrawals"`

	TotalDifficulty *big.Int `json:"totalDifficulty,omitempty"  gencodec:"required"`
	Size            *big.Int `json:"size,omitempty"  gencodec:"required"`
}

type Header struct {
	Hash            common.Hash      `json:"hash"             gencodec:"required"`
	ParentHash      common.Hash      `json:"parentHash"       gencodec:"required"`
	UncleHash       common.Hash      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase        common.Address   `json:"miner"            gencodec:"required"`
	Root            common.Hash      `json:"stateRoot"        gencodec:"required"`
	TxHash          common.Hash      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash     common.Hash      `json:"receiptsRoot"     gencodec:"required"`
	Bloom           types.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty      *hexutil.Big     `json:"difficulty"       gencodec:"required"`
	TotalDifficulty *hexutil.Big     `json:"totalDifficulty,omitempty"  gencodec:"required"`
	Number          *hexutil.Big     `json:"number"           gencodec:"required"`
	GasLimit        hexutil.Uint64   `json:"gasLimit"         gencodec:"required"`
	GasUsed         hexutil.Uint64   `json:"gasUsed"          gencodec:"required"`
	Time            hexutil.Uint64   `json:"timestamp"        gencodec:"required"`
	Extra           hexutil.Bytes    `json:"extraData"        gencodec:"required"`
	MixDigest       common.Hash      `json:"mixHash"`
	Nonce           types.BlockNonce `json:"nonce"`

	Size hexutil.Uint64 `json:"size"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *hexutil.Big `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *hexutil.Uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *hexutil.Uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`
}

type Transaction struct {
	BlockHash        *common.Hash      `json:"blockHash"`
	BlockNumber      *string           `json:"blockNumber"`
	From             common.Address    `json:"from"`
	Gas              string            `json:"gas"`
	GasPrice         string            `json:"gasPrice"`
	GasFeeCap        *string           `json:"maxFeePerGas,omitempty"`
	GasTipCap        *string           `json:"maxPriorityFeePerGas,omitempty"`
	Hash             common.Hash       `json:"hash"`
	Input            string            `json:"input"`
	Nonce            string            `json:"nonce"`
	To               *common.Address   `json:"to"`
	TransactionIndex *string           `json:"transactionIndex"`
	Value            string            `json:"value"`
	Type             string            `json:"type"`
	Accesses         *types.AccessList `json:"accessList,omitempty"`
	ChainID          *string           `json:"chainId,omitempty"`
	V                string            `json:"v"`
	R                string            `json:"r"`
	S                string            `json:"s"`
}

func BuildBlock(block *RpcBlock) *models.Block {
	txs := make([]*models.Transaction, len(*block.Txs))
	for i, tx := range *block.Txs {
		txs[i] = &models.Transaction{
			From:     tx.From.Hex(),
			Gas:      tx.Gas.String(),
			GasPrice: tx.GasPrice.String(),
			Hash:     tx.Hash.String(),
			Input:    tx.Input.String(),
			Nonce:    tx.Nonce.String(),
			Value:    tx.Value.String(),
			Type:     tx.Type.String(),
			V:        tx.V.String(),
			R:        tx.R.String(),
			S:        tx.S.String(),
		}
		if tx.BlockHash != nil {
			blockHash := tx.BlockHash.Hex()
			txs[i].BlockHash = &blockHash
		}
		if tx.BlockNumber != nil {
			blockNumber := tx.BlockNumber.String()
			txs[i].BlockNumber = &blockNumber
		}
		if tx.GasFeeCap != nil {
			gasFeeCap := tx.GasFeeCap.String()
			txs[i].MaxFeePerGas = &gasFeeCap
		}
		if tx.GasTipCap != nil {
			gasTipCap := tx.GasTipCap.String()
			txs[i].MaxPriorityFeePerGas = &gasTipCap
		}
		if tx.To != nil {
			to := tx.To.String()
			txs[i].To = &to
		}
		if tx.TransactionIndex != nil {
			txIndex := tx.TransactionIndex.String()
			txs[i].TransactionIndex = &txIndex
		}

		if tx.Accesses != nil {
			accessList := make([]*models.AccessTuple, len(*tx.Accesses))
			for in, tuple := range *tx.Accesses {
				for j, storageKey := range tuple.StorageKeys {
					accessList[in].StorageKeys[j] = storageKey.Hex()
				}
				accessList[in].Address = tuple.Address.Hex()
			}

			txs[i].AccessList = accessList
		}
		if tx.ChainID != nil {
			chainID := tx.ChainID.String()
			txs[i].ChainID = &chainID
		}
	}

	var withdrawals []*models.Withdrawal
	if block.Withdrawals != nil {
		withdrawals := make([]*models.Withdrawal, len(*block.Withdrawals))
		for i, withdrawal := range *block.Withdrawals {
			withdrawals[i] = &models.Withdrawal{
				Index:     util.Uint64ToHex(withdrawal.Index),
				Validator: util.Uint64ToHex(withdrawal.Validator),
				Address:   withdrawal.Address.Hex(),
				Amount:    util.Uint64ToHex(withdrawal.Amount),
			}
		}
	}
	var uncles []*models.Header
	if block.Uncles != nil {
		uncles = make([]*models.Header, len(*block.Uncles))
		for i, uncle := range *block.Uncles {
			uncles[i] = &models.Header{
				ParentHash:       uncle.ParentHash.String(),
				Sha3Uncles:       uncle.UncleHash.String(),
				Miner:            uncle.Coinbase.String(),
				StateRoot:        uncle.Root.String(),
				TransactionsRoot: uncle.TxHash.String(),
				ReceiptsRoot:     uncle.ReceiptHash.String(),
				LogsBloom:        hexutil.Bytes(uncle.Bloom.Bytes()).String(),
				GasLimit:         uncle.GasLimit.String(),
				GasUsed:          uncle.GasUsed.String(),
				Timestamp:        uncle.Time.String(),
				ExtraData:        uncle.Extra.String(),
				MixHash:          uncle.MixDigest.String(),
				Nonce:            util.Uint64ToHex(uncle.Nonce.Uint64()),
			}
			if uncle.Difficulty != nil {
				difficulty := uncle.Difficulty.String()
				uncles[i].Difficulty = &difficulty
			}
			if uncle.Number != nil {
				number := uncle.Number.String()
				uncles[i].Number = &number
			}
			if uncle.BaseFee != nil {
				baseFeePerGas := uncle.BaseFee.String()
				uncles[i].BaseFeePerGas = &baseFeePerGas
			}
			if uncle.WithdrawalsHash != nil {
				withdrawalsHash := uncle.WithdrawalsHash.String()
				uncles[i].WithdrawalsRoot = &withdrawalsHash
			}
			if uncle.BlobGasUsed != nil {
				blobGasUsed := uncle.BlobGasUsed.String()
				uncles[i].BlobGasUsed = &blobGasUsed
			}
			if uncle.ExcessBlobGas != nil {
				excessBlobGas := uncle.ExcessBlobGas.String()
				uncles[i].ExcessBlobGas = &excessBlobGas
			}
			if uncle.ParentBeaconRoot != nil {
				parentBeaconRoot := uncle.ParentBeaconRoot.String()
				uncles[i].ParentBeaconBlockRoot = &parentBeaconRoot
			}
		}
	}

	blockInfo := &models.Block{
		Hash:             block.Hash.String(),
		ParentHash:       block.ParentHash.String(),
		Sha3Uncles:       block.UncleHash.String(),
		Miner:            block.Coinbase.String(),
		StateRoot:        block.Root.String(),
		TransactionsRoot: block.TxHash.String(),
		ReceiptsRoot:     block.ReceiptHash.String(),
		LogsBloom:        hexutil.Bytes(block.Bloom.Bytes()).String(),
		GasLimit:         block.GasLimit.String(),
		GasUsed:          block.GasUsed.String(),
		Timestamp:        block.Time.String(),
		ExtraData:        block.Extra.String(),
		MixHash:          block.MixDigest.String(),
		Nonce:            hexutil.Bytes(block.Nonce[:]).String(),

		Transactions: txs,
		Uncles:       uncles,
		Withdrawals:  withdrawals,

		Size: block.Size.String(),
	}
	if block.Difficulty != nil {
		difficulty := block.Difficulty.String()
		blockInfo.Difficulty = &difficulty
	}
	if block.Number != nil {
		number := block.Number.String()
		blockInfo.Number = &number
	}
	if block.BaseFee != nil {
		baseFeePerGas := block.BaseFee.String()
		blockInfo.BaseFeePerGas = &baseFeePerGas
	}
	if block.TotalDifficulty != nil {
		totalDifficulty := block.TotalDifficulty.String()
		blockInfo.TotalDifficulty = &totalDifficulty
	}
	if block.WithdrawalsHash != nil {
		withdrawalsHash := block.WithdrawalsHash.String()
		blockInfo.WithdrawalsRoot = &withdrawalsHash
	}
	if block.BlobGasUsed != nil {
		blobGasUsed := block.BlobGasUsed.String()
		blockInfo.BlobGasUsed = &blobGasUsed
	}
	if block.ExcessBlobGas != nil {
		excessBlobGas := block.ExcessBlobGas.String()
		blockInfo.ExcessBlobGas = &excessBlobGas
	}
	if block.ParentBeaconRoot != nil {
		parentBeaconRoot := block.ParentBeaconRoot.String()
		blockInfo.ParentBeaconBlockRoot = &parentBeaconRoot
	}
	return blockInfo
}

type RpcBlock struct {
	Hash        common.Hash      `json:"hash"             gencodec:"required"`
	ParentHash  common.Hash      `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address   `json:"miner"            gencodec:"required"`
	Root        common.Hash      `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash      `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty  *hexutil.Big     `json:"difficulty"       gencodec:"required"`
	Number      *hexutil.Big     `json:"number"           gencodec:"required"`
	GasLimit    hexutil.Uint64   `json:"gasLimit"         gencodec:"required"`
	GasUsed     hexutil.Uint64   `json:"gasUsed"          gencodec:"required"`
	Time        hexutil.Uint64   `json:"timestamp"        gencodec:"required"`
	Extra       hexutil.Bytes    `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash      `json:"mixHash"`
	Nonce       types.BlockNonce `json:"nonce"`
	Size        hexutil.Uint64   `json:"size"`

	// belows are not belong to header
	TotalDifficulty *hexutil.Big `json:"totalDifficulty,omitempty"  gencodec:"required"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *hexutil.Big `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *hexutil.Uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *hexutil.Uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`

	Withdrawals *types.Withdrawals `json:"withdrawals,omitempty"`
	Txs         *[]Tx              `json:"transactions,omitempty"`
	Uncles      *[]Header          `json:"uncles,omitempty"`
}

type Tx struct {
	BlockHash        *common.Hash      `json:"blockHash"`
	BlockNumber      *hexutil.Uint64   `json:"blockNumber"`
	From             common.Address    `json:"from"`
	Gas              hexutil.Uint64    `json:"gas"`
	GasPrice         hexutil.Big       `json:"gasPrice"`
	GasFeeCap        *hexutil.Big      `json:"maxFeePerGas,omitempty"`
	GasTipCap        *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
	Hash             common.Hash       `json:"hash"`
	Input            hexutil.Bytes     `json:"input"`
	Nonce            hexutil.Uint64    `json:"nonce"`
	To               *common.Address   `json:"to"`
	TransactionIndex *hexutil.Uint64   `json:"transactionIndex"`
	Value            hexutil.Big       `json:"value"`
	Type             hexutil.Uint64    `json:"type"`
	Accesses         *types.AccessList `json:"accessList,omitempty"`
	ChainID          *hexutil.Uint64   `json:"chainId,omitempty"`
	V                hexutil.Big       `json:"v"`
	R                hexutil.Big       `json:"r"`
	S                hexutil.Big       `json:"s"`
}
