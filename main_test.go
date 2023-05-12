package main

import (
	"testing"
)

func TestRunMPC(t *testing.T) {
	nodes := 3
	secrets := []int{1, 2}
	n := 3
	k := 2
	address := "localhost:50051"

	sum, receivedSums, err := runMPC(nodes, secrets, n, k, address)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expectedSum := 3
	if sum != expectedSum {
		t.Errorf("expected: %d, actual: %d", expectedSum, sum)
	}

	for _, receivedSum := range receivedSums {
		if len(receivedSum) != 5 {
			t.Errorf("%q should have 4 item(s), but has %d", receivedSum, len(receivedSum))
		}
	}
}
