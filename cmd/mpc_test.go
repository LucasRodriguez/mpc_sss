package main

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/LucasRodriguez/mpc_sss/mpc"
	"github.com/hashicorp/vault/shamir"
)

func TestMPCExample(t *testing.T) {
	address := "localhost:8080"
	secrets := []int{42, 123, 987}
	n := 7
	k := 3

	_, allShares, err := mpc.RunMPCExample(address, secrets, n, k)
	if err != nil {
		t.Fatalf("Error running MPC example: %v", err)
	}

	fmt.Printf("Iteration index: 0\n")
	fmt.Printf("Shares to use for reconstruction:\n")
	for i := 0; i < k; i++ {
		share := allShares[0][i]
		fmt.Printf("%0x\n", share)
	}

	for i := 0; i < len(secrets); i++ {
		secretShares := make([][]byte, n)
		for j := 0; j < n; j++ {
			secretShares[j] = allShares[i][j]
		}

		fmt.Printf("\nSecret %d\n", i)
		fmt.Printf("Secret shares:\n")
		for _, share := range secretShares {
			fmt.Printf("%0x\n", share)
		}

		secretBytes, err := shamir.Combine(secretShares)
		if err != nil {
			t.Fatalf("Error combining shares for secret %d: %v", i, err)
		}

		fmt.Printf("Combined secret bytes: %0x\n", secretBytes)

		secret := bytesToInt(secretBytes)
		fmt.Printf("Reconstructed secret %d: %d\n", i, secret)

		if secret != secrets[i] {
			t.Errorf(
				"Reconstructed secret %d does not match original secret. Expected %d, got %d",
				i,
				secrets[i],
				secret,
			)
		} else {
			fmt.Printf("Reconstructed secret %d matches the original secret\n", i)
		}
	}
}

func bytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(append(make([]byte, 4-len(b[:len(b)-1])), b...)))
}

func intToBytes(n int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	return buf
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