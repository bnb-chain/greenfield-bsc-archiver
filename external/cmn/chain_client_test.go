package cmn

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainClient_GetObjectMeta(t *testing.T) {
	chainClient, err := NewChainClient("https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org")
	if err != nil {
		t.Fatal(err)
	}
	objectContent, err := chainClient.GetObjectMeta(context.Background(), "testnet-bsc-block", "blocks_s42499000_e42499999")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	assert.Nil(t, err)
	t.Log(objectContent)
}
