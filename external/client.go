package external

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"greeenfield-bsc-archiver/config"
)

const BSCBlockConfirmNum = 3

type IClient interface {
	GetFinalizedBlockNum(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, int2 *big.Int) (*types.Block, error)
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

func (c *Client) BlockByNumber(ctx context.Context, int2 *big.Int) (*types.Block, error) {
	return c.ethClient.BlockByNumber(ctx, int2)
}
