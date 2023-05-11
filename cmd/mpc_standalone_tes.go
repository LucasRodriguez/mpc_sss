package main

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/hashicorp/vault/shamir"
)

func TestMPCExample(t *testing.T) {
	secrets := []int{1, 22, 983}
	n := 8
	k := 6

	allShares := make([][][]byte, len(secrets))

	for i, secret := range secrets {
		secretBytes := intToBytes(secret)
		shares, err := shamir.Split(secretBytes, n, k)
		if err != nil {
			t.Fatalf("Error splitting secret %d: %v", i, err)
		}
		allShares[i] = shares
	}

	for i := 0; i < len(secrets); i++ {
		secretShares := allShares[i]

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
	return int(binary.BigEndian.Uint32(append(make([]byte, 4-len(b)), b...)))
}

func intToBytes(n int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	return buf
}