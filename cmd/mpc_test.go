package main

import (
	"testing"
	"github.com/LucasRodriguez/mpc_sss/mpc"
)

func TestRunMPCExample(t *testing.T) {
	secrets := []int{1, 22, 983}
	n := 8
	k := 6

	results, allShares, err := mpc.RunMPCExample("localhost:50051", secrets, n, k)
	if err != nil {
		t.Fatalf("Error running MPC example: %v", err)
	}

	// Validate results
	if len(results) != n {
		t.Errorf("Expected %d results, but got %d", n, len(results))
	}

	// Validate allShares
	if len(allShares) != len(secrets) {
		t.Errorf("Expected %d sets of shares, but got %d", len(secrets), len(allShares))
	}

	for i, shares := range allShares {
		if len(shares) != n {
			t.Errorf("Expected %d shares for secret %d, but got %d", n, i, len(shares))
		}
	}
}