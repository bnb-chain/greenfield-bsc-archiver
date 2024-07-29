package util

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// StringToUint64 converts string to uint64
func StringToUint64(str string) (uint64, error) {
	ui64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return ui64, nil
}

// StringToInt64 converts string to int64
func StringToInt64(str string) (int64, error) {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i64, nil
}

func SplitByComma(str string) []string {
	str = strings.TrimSpace(str)
	strArr := strings.Split(str, ",")
	var trimStr []string
	for _, item := range strArr {
		if len(strings.TrimSpace(item)) > 0 {
			trimStr = append(trimStr, strings.TrimSpace(item))
		}
	}
	return trimStr
}

func JoinWithComma(slice []string) string {
	return strings.Join(slice, ",")
}

func GenerateHash(hexStr string) ([]byte, error) {
	if !strings.HasPrefix(hexStr, "0x") {
		hexStr = "0x" + hexStr
	}
	bz, err := hexutil.Decode(hexStr)
	if err != nil {
		return nil, nil
	}
	hash := sha256.New()
	hash.Write(bz)
	return hash.Sum(nil), nil
}

// Uint64ToString coverts uint64 to string
func Uint64ToString(u uint64) string {
	return strconv.FormatUint(u, 10)
}

// Int64ToString coverts uint64 to string
func Int64ToString(u int64) string {
	return strconv.FormatInt(u, 10)
}

// HexToUint64 converts hex string to uint64
func HexToUint64(hexStr string) (uint64, error) {
	intValue, err := strconv.ParseUint(hexStr, 0, 64)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

// Int64ToHex converts int64 to hex string
func Int64ToHex(int64 int64) string {
	return "0x" + strconv.FormatInt(int64, 16)
}

// Uint64ToHex converts uint64 to hex string
func Uint64ToHex(uint64 uint64) string {
	return "0x" + strconv.FormatUint(uint64, 16)
}

// Uint64ToHexNonce converts uint64 to hex string
func Uint64ToHexNonce(uint64 uint64) string {
	return fmt.Sprintf("0x%016x", uint64)
}

// IsHexHash verifies whether a string can represent a valid hex hash
func IsHexHash(hash string) bool {
	match, _ := regexp.MatchString(`^0x[0-9a-fA-F]{64}$`, hash)
	return match
}

// Uint64ToHexPtr safely converts a *uint64 to a *string
func Uint64ToHexPtr(u *uint64) *string {
	if u == nil {
		return nil
	}
	s := Uint64ToHex(*u)
	return &s
}

func HexutilBigToBigInt(hexBig *hexutil.Big) *big.Int {
	if hexBig == nil {
		return nil
	}
	return hexBig.ToInt()
}

func Uint64ToBigInt(u hexutil.Uint64) *big.Int {
	return new(big.Int).SetUint64(uint64(u))
}
