package cmn

import (
	"context"
	"fmt"
	"io"
	"testing"
)

func TestSPClient_GetBundleObject(t *testing.T) {
	spClient, err := NewSPClient("https://gnfd-testnet-sp2.bnbchain.org")
	if err != nil {
		t.Fatal(err)
	}
	objectContent, err := spClient.GetBundleObject(context.Background(), "bsc-historical-blocks", "blocks_s901_e1000")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	objectBytes, err := io.ReadAll(objectContent)
	t.Log(string(objectBytes))
}
