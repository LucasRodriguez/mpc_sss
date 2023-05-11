package main

import (
	"testing"
	"github.com/LucasRodriguez/mpc_sss/mpc"
	"github.com/stretchr/testify/assert"
)

func TestRunMPCExample(t *testing.T) {
	secrets := []int{1, 22, 983}
	n := 8
	k := 6

	results, allShares, err := mpc.RunMPCExample("localhost:50051", secrets, n, k)

	// Check for errors
	assert.NoError(t, err, "Error running MPC example")

	// Validate results
	expectedResultsLen := n
	assert.Equal(t, expectedResultsLen, len(results), "Expected %d results, but got %d", n, len(results))

	// Validate allShares
	expectedSharesLen := len(secrets)
	assert.Equal(t, expectedSharesLen, len(allShares), "Expected %d sets of shares, but got %d", len(secrets), len(allShares))

	for i, shares := range allShares {
		expectedSharesPerSecret := n
		assert.Equal(t, expectedSharesPerSecret, len(shares), "Expected %d shares for secret %d, but got %d", n, i, len(shares))
	}
}