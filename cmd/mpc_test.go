package main

import (
	"testing"
	"encoding/binary"	
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/LucasRodriguez/mpc_sss/mpc"
	"github.com/hashicorp/vault/shamir"
)

func TestMPCExample(t *testing.T) {
	address := "localhost:50051"
	secrets := []int{123, 456, 789}
	n := 5
	k := 3

	results, allShares, err := mpc.RunMPCExample(address, secrets, n, k)

	assert.NoError(t, err)

	for i, result := range results {
		// Try to reconstruct the secret using the result and k-1 shares from other secrets
		sharesToUse := make([][]byte, k)
		sharesToUse[0] = intToBytes(result)

		index := 0
		for j := 1; j < k; j++ {
			if index == i {
				index++
			}
			sharesToUse[j] = allShares[index][i]
			index++
		}

		fmt.Printf("Shares to use for reconstruction:\n")
		for _, share := range sharesToUse {
			fmt.Printf("%x\n", share)
		}

		reconstructedSecret, err := shamir.Combine(sharesToUse)
		assert.NoError(t, err)

		// Check if the reconstructed secret is in the original secrets
		assert.Contains(t, secrets, int(bytesToInt(reconstructedSecret)))
	}
}

func intToBytes(n int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	return buf
}

func bytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}

func padShares(shares [][]byte) [][]byte {
	maxLen := 0
	for _, share := range shares {
		if len(share) > maxLen {
			maxLen = len(share)
		}
	}

	paddedShares := make([][]byte, len(shares))
	for i, share := range shares {
		paddedShares[i] = make([]byte, maxLen)
		copy(paddedShares[i][maxLen-len(share):], share)
	}

	return paddedShares
}