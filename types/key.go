package types

import (
	"fmt"
	"strconv"
	"strings"
)

const RootLength = 32

func GetBlockName(block uint64) string {
	return fmt.Sprintf("block_h%d", block)
}

func GetBundleName(startBlock, endBlock uint64) string {
	return fmt.Sprintf("blocks_s%d_e%d", startBlock, endBlock)
}

//func ParseBlobName(blobName string) (slot uint64, index uint64, err error) {
//	parts := strings.Split(blobName, "_")
//	slot, err = strconv.ParseUint(parts[1][1:], 10, 64)
//	if err != nil {
//		return
//	}
//	index, err = strconv.ParseUint(parts[2][1:], 10, 64)
//	if err != nil {
//		return
//	}
//	return
//}

func ParseBundleName(bundleName string) (startBlock, endBlock uint64, err error) {
	parts := strings.Split(bundleName, "_")
	startBlock, err = strconv.ParseUint(parts[1][1:], 10, 64)
	if err != nil {
		return
	}
	endBlock, err = strconv.ParseUint(parts[2][1:], 10, 64)
	if err != nil {
		return
	}
	return
}
