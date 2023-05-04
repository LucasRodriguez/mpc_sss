package sss

import (
	"github.com/hashicorp/vault/shamir"
	"math/big"
	
)

func SplitSecret(secret int, n int, k int) ([][]byte, error) {
	secretBigInt := big.NewInt(int64(secret))
	shares, err := shamir.Split(secretBigInt.Bytes(), n, k)
	if err != nil {
		return nil, err
	}
	return shares, nil
}