package external

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"greeenfield-bsc-archiver/config"
	types2 "greeenfield-bsc-archiver/types"
)

const BSCBlockConfirmNum = 3

type IClient interface {
	GetFinalizedBlockNum(ctx context.Context) (uint64, error)
	GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*types2.RpcBlock, error)
	BlockByNumber(ctx context.Context, blockNumber *big.Int) (*types.Block, error)
}

type Client struct {
	ethClient *ethclient.Client
	rpcClient *rpc.Client
	cfg       *config.SyncerConfig
}

func NewClient(cfg *config.SyncerConfig) IClient {
	ethClient, err := ethclient.Dial(cfg.RPCAddrs[0])
	if err != nil {
		panic("new eth client error")
	}
	cli := &Client{
		cfg:       cfg,
		ethClient: ethClient,
	}
	rpcClient, err := rpc.DialContext(context.Background(), cfg.RPCAddrs[0])
	if err != nil {
		panic("new rpc client error")
	}
	cli.rpcClient = rpcClient
	return cli
}

func (c *Client) GetFinalizedBlockNum(ctx context.Context) (uint64, error) {
	var head *types.Header
	if err := c.rpcClient.CallContext(ctx, &head, "eth_getFinalizedHeader", BSCBlockConfirmNum); err != nil {
		return 0, err
	}
	if head == nil || head.Number == nil {
		return 0, ethereum.NotFound
	}
	return head.Number.Uint64(), nil
}

func (c *Client) BlockByNumber(ctx context.Context, blockNumber *big.Int) (*types.Block, error) {
	return c.ethClient.BlockByNumber(ctx, blockNumber)
}

func (c *Client) GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*types2.RpcBlock, error) {
	var block *types2.RpcBlock
	if err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", hexutil.EncodeBig(blockNumber), true); err != nil {
		return nil, err
	}
	if block == nil || block.TotalDifficulty == nil {
		return nil, ethereum.NotFound
	}
	return block, nil
}
