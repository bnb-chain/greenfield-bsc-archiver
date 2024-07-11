package cmn

import (
	"context"
	"fmt"
	"io"
	"testing"
)

func TestSPClient_GetBundleObject(t *testing.T) {
	spClient, err := NewSPClient("https://gnfd-testnet-sp3.nodereal.io")
	if err != nil {
		t.Fatal(err)
	}
	objectContent, err := spClient.GetBundleObject(context.Background(), "bsc-blob", "blobs_s40202240_e40202439")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	objectBytes, err := io.ReadAll(objectContent)
	t.Log(string(objectBytes))
}
