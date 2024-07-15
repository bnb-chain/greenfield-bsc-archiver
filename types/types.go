package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"greeenfield-bsc-archiver/models"
	"greeenfield-bsc-archiver/util"
)

//type Block struct {
//	ParentHash  common.Hash      `json:"parentHash"       gencodec:"required"`
//	UncleHash   common.Hash      `json:"sha3Uncles"       gencodec:"required"`
//	Coinbase    common.Address   `json:"miner"`
//	Root        common.Hash      `json:"stateRoot"        gencodec:"required"`
//	TxHash      common.Hash      `json:"transactionsRoot" gencodec:"required"`
//	ReceiptHash common.Hash      `json:"receiptsRoot"     gencodec:"required"`
//	Bloom       types.Bloom      `json:"logsBloom"        gencodec:"required"`
//	Difficulty  *big.Int         `json:"difficulty"       gencodec:"required"`
//	Number      *big.Int         `json:"number"           gencodec:"required"`
//	GasLimit    uint64           `json:"gasLimit"         gencodec:"required"`
//	GasUsed     uint64           `json:"gasUsed"          gencodec:"required"`
//	Time        uint64           `json:"timestamp"        gencodec:"required"`
//	Extra       []byte           `json:"extraData"        gencodec:"required"`
//	MixDigest   common.Hash      `json:"mixHash"`
//	Nonce       types.BlockNonce `json:"nonce"`
//
//	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
//	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`
//
//	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
//	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`
//
//	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
//	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`
//
//	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
//	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`
//
//	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
//	ParentBeaconRoot *common.Hash        `json:"parentBeaconBlockRoot" rlp:"optional"`
//	Transactions     []*Transaction      `json:"transactions"`
//	Uncles           []*Header           `json:"uncles"`
//	Withdrawals      []*types.Withdrawal `json:"withdrawals"`
//}

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
	ParentBeaconRoot *common.Hash         `json:"parentBeaconBlockRoot" rlp:"optional"`
	Transactions     []*types.Transaction `json:"transactions"`
	Uncles           []*types.Header      `json:"uncles"`
	Withdrawals      []*types.Withdrawal  `json:"withdrawals"`

	TotalDifficulty *big.Int `json:"totalDifficulty,omitempty"  gencodec:"required"`
	Size            *big.Int `json:"size,omitempty"  gencodec:"required"`
}

type Header struct {
	ParentHash  common.Hash      `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address   `json:"miner"`
	Root        common.Hash      `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash      `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty  *string          `json:"difficulty"       gencodec:"required"`
	Number      *string          `json:"number"           gencodec:"required"`
	GasLimit    uint64           `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64           `json:"gasUsed"          gencodec:"required"`
	Time        uint64           `json:"timestamp"        gencodec:"required"`
	Extra       []byte           `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash      `json:"mixHash"`
	Nonce       types.BlockNonce `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *string `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`
}

//type Header struct {
//	ParentHash  common.Hash      `json:"parentHash"       gencodec:"required"`
//	UncleHash   common.Hash      `json:"sha3Uncles"       gencodec:"required"`
//	Coinbase    common.Address   `json:"miner"`
//	Root        common.Hash      `json:"stateRoot"        gencodec:"required"`
//	TxHash      common.Hash      `json:"transactionsRoot" gencodec:"required"`
//	ReceiptHash common.Hash      `json:"receiptsRoot"     gencodec:"required"`
//	Bloom       types.Bloom      `json:"logsBloom"        gencodec:"required"`
//	Difficulty  *big.Int         `json:"difficulty"       gencodec:"required"`
//	Number      *big.Int         `json:"number"           gencodec:"required"`
//	GasLimit    uint64           `json:"gasLimit"         gencodec:"required"`
//	GasUsed     uint64           `json:"gasUsed"          gencodec:"required"`
//	Time        uint64           `json:"timestamp"        gencodec:"required"`
//	Extra       []byte           `json:"extraData"        gencodec:"required"`
//	MixDigest   common.Hash      `json:"mixHash"`
//	Nonce       types.BlockNonce `json:"nonce"`
//
//	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
//	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`
//
//	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
//	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`
//
//	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
//	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`
//
//	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
//	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`
//
//	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
//	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`
//}

//type Transaction struct {
//	BlockHash        *common.Hash      `json:"blockHash"`
//	BlockNumber      *uint64           `json:"blockNumber"`
//	From             common.Address    `json:"from"`
//	Gas              uint64            `json:"gas"`
//	GasPrice         big.Int           `json:"gasPrice"`
//	GasFeeCap        *big.Int          `json:"maxFeePerGas,omitempty"`
//	GasTipCap        *big.Int          `json:"maxPriorityFeePerGas,omitempty"`
//	Hash             common.Hash       `json:"hash"`
//	Input            []byte            `json:"input"`
//	Nonce            uint64            `json:"nonce"`
//	To               *common.Address   `json:"to"`
//	TransactionIndex *uint64           `json:"transactionIndex"`
//	Value            big.Int           `json:"value"`
//	Type             uint64            `json:"type"`
//	Accesses         *types.AccessList `json:"accessList,omitempty"`
//	ChainID          *uint64           `json:"chainId,omitempty"`
//	V                big.Int           `json:"v"`
//	R                big.Int           `json:"r"`
//	S                big.Int           `json:"s"`
//}

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

func BuildBlock(bundleBlock *Block) *models.Block {
	txs := make([]*models.Transaction, len(bundleBlock.Transactions))
	for i, tx := range bundleBlock.Transactions {
		txs[i] = &models.Transaction{
			From: tx.From.Hex(),
			//Gas:      util.Uint64ToHex(tx.Gas),
			Gas:      tx.Gas,
			GasPrice: tx.GasPrice,
			Hash:     tx.Hash.String(),
			Input:    string(tx.Input),
			Nonce:    tx.Nonce,
			//Nonce:    util.Uint64ToHex(tx.Nonce),
			//Value:    tx.Value.String(),
			Value: tx.Value,
			Type:  tx.Type,
			//Type:  util.Uint64ToHex(tx.Type),

			//V: tx.V.String(),
			//R: tx.R.String(),
			//S: tx.S.String(),
			V: tx.V,
			R: tx.R,
			S: tx.S,
		}
		if tx.BlockHash != nil {
			blockHash := tx.BlockHash.Hex()
			txs[i].BlockHash = &blockHash
		}
		if tx.BlockNumber != nil {
			//blockNumber := util.Uint64ToHex(*tx.BlockNumber)
			//txs[i].BlockNumber = &blockNumber
			txs[i].BlockNumber = tx.BlockNumber
		}
		if tx.GasFeeCap != nil {
			//gasFeeCap := util.Uint64ToHex(tx.GasFeeCap.Uint64())
			//txs[i].MaxFeePerGas = &gasFeeCap
			txs[i].MaxFeePerGas = tx.GasFeeCap
		}
		if tx.GasTipCap != nil {
			//gasTipCap := tx.GasTipCap.String()
			//txs[i].MaxPriorityFeePerGas = &gasTipCap
			txs[i].MaxPriorityFeePerGas = tx.GasTipCap
		}
		if tx.To != nil {
			to := tx.To.String()
			txs[i].To = &to
		}
		if tx.TransactionIndex != nil {
			//txIndex := util.Uint64ToHex(*tx.TransactionIndex)
			//txs[i].TransactionIndex = &txIndex
			txs[i].TransactionIndex = tx.TransactionIndex
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
			//chainID := util.Uint64ToHex(*tx.ChainID)
			//txs[i].ChainID = &chainID
			txs[i].ChainID = tx.ChainID
		}
	}

	withdrawals := make([]*models.Withdrawal, len(bundleBlock.Withdrawals))
	for i, withdrawal := range bundleBlock.Withdrawals {
		withdrawals[i] = &models.Withdrawal{
			Index:     util.Uint64ToHex(withdrawal.Index),
			Validator: util.Uint64ToHex(withdrawal.Validator),
			Address:   withdrawal.Address.Hex(),
			Amount:    util.Uint64ToHex(withdrawal.Amount),
		}
	}

	uncles := make([]*models.Header, len(bundleBlock.Uncles))
	for i, uncle := range bundleBlock.Uncles {
		uncles[i] = &models.Header{
			ParentHash:       uncle.ParentHash.String(),
			Sha3Uncles:       uncle.UncleHash.String(),
			Miner:            uncle.Coinbase.String(),
			StateRoot:        uncle.Root.String(),
			TransactionsRoot: uncle.TxHash.String(),
			ReceiptsRoot:     uncle.ReceiptHash.String(),
			LogsBloom:        hexutil.Bytes(uncle.Bloom.Bytes()).String(),
			GasLimit:         util.Uint64ToHex(uncle.GasLimit),
			GasUsed:          util.Uint64ToHex(uncle.GasUsed),
			Timestamp:        util.Uint64ToHex(uncle.Time),
			ExtraData:        hexutil.Bytes(uncle.Extra).String(),
			MixHash:          uncle.MixDigest.String(),
			Nonce:            util.Uint64ToHex(uncle.Nonce.Uint64()),
		}
		if uncle.Difficulty != nil {
			//difficulty := uncle.Difficulty.String()
			//uncles[i].Difficulty = &difficulty
			uncles[i].Difficulty = uncle.Difficulty
		}
		if uncle.Number != nil {
			//number := uncle.Number.String()
			//uncles[i].Number = &number
			uncles[i].Number = uncle.Number
		}
		if uncle.BaseFee != nil {
			//baseFeePerGas := uncle.BaseFee.String()
			//uncles[i].BaseFeePerGas = &baseFeePerGas
			uncles[i].BaseFeePerGas = uncle.BaseFee
		}
		if uncle.WithdrawalsHash != nil {
			withdrawalsHash := uncle.WithdrawalsHash.String()
			uncles[i].WithdrawalsRoot = &withdrawalsHash
		}
		if uncle.BlobGasUsed != nil {
			blobGasUsed := util.Uint64ToHex(*uncle.BlobGasUsed)
			uncles[i].BlobGasUsed = &blobGasUsed
		}
		if uncle.ExcessBlobGas != nil {
			excessBlobGas := util.Uint64ToHex(*uncle.ExcessBlobGas)
			uncles[i].ExcessBlobGas = &excessBlobGas
		}
		if uncle.ParentBeaconRoot != nil {
			parentBeaconRoot := uncle.ParentBeaconRoot.String()
			uncles[i].ParentBeaconBlockRoot = &parentBeaconRoot
		}
	}

	blockInfo := &models.Block{
		Hash:             bundleBlock.Hash.String(),
		ParentHash:       bundleBlock.ParentHash.String(),
		Sha3Uncles:       bundleBlock.UncleHash.String(),
		Miner:            bundleBlock.Coinbase.String(),
		StateRoot:        bundleBlock.Root.String(),
		TransactionsRoot: bundleBlock.TxHash.String(),
		ReceiptsRoot:     bundleBlock.ReceiptHash.String(),
		LogsBloom:        hexutil.Bytes(bundleBlock.Bloom.Bytes()).String(),
		GasLimit:         util.Uint64ToHex(bundleBlock.GasLimit),
		GasUsed:          util.Uint64ToHex(bundleBlock.GasUsed),
		Timestamp:        util.Uint64ToHex(bundleBlock.Time),
		ExtraData:        hexutil.Bytes(bundleBlock.Extra).String(),
		MixHash:          bundleBlock.MixDigest.String(),
		Nonce:            hexutil.Bytes(bundleBlock.Nonce[:]).String(),

		Transactions: txs,
		Uncles:       uncles,
		Withdrawals:  withdrawals,

		Size: util.Uint64ToHex(bundleBlock.Size.Uint64()),
	}
	if bundleBlock.Difficulty != nil {
		difficulty := util.Uint64ToHex(bundleBlock.Difficulty.Uint64())
		blockInfo.Difficulty = &difficulty
	}
	if bundleBlock.Number != nil {
		number := util.Uint64ToHex(bundleBlock.Number.Uint64())
		blockInfo.Number = &number
	}
	if bundleBlock.BaseFee != nil {
		baseFeePerGas := util.Uint64ToHex(bundleBlock.BaseFee.Uint64())
		blockInfo.BaseFeePerGas = &baseFeePerGas
	}
	if bundleBlock.TotalDifficulty != nil {
		totalDifficulty := util.Uint64ToHex(bundleBlock.TotalDifficulty.Uint64())
		blockInfo.TotalDifficulty = &totalDifficulty
	}
	if bundleBlock.WithdrawalsHash != nil {
		withdrawalsHash := bundleBlock.WithdrawalsHash.String()
		blockInfo.WithdrawalsRoot = &withdrawalsHash
	}
	if bundleBlock.BlobGasUsed != nil {
		blobGasUsed := util.Uint64ToHex(*bundleBlock.BlobGasUsed)
		blockInfo.BlobGasUsed = &blobGasUsed
	}
	if bundleBlock.ExcessBlobGas != nil {
		excessBlobGas := util.Uint64ToHex(*bundleBlock.ExcessBlobGas)
		blockInfo.ExcessBlobGas = &excessBlobGas
	}
	if bundleBlock.ParentBeaconRoot != nil {
		parentBeaconRoot := bundleBlock.ParentBeaconRoot.String()
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
	Difficulty  hexutil.Big      `json:"difficulty"       gencodec:"required"`
	Number      hexutil.Big      `json:"number"           gencodec:"required"`
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
	BaseFee *hexutil.Big `json:"baseFeePerGas,omitempty"`

	Txs    *[]common.Hash `json:"transactions,omitempty"`
	Uncles *[]common.Hash `json:"uncles,omitempty"`
}
