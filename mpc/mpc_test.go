package mpc_test

import (
	"fmt"
	"testing"
	"time"    
	"bytes" // Add this line

	"github.com/LucasRodriguez/mpc_sss/mpc"
	"github.com/LucasRodriguez/mpc_sss/client"
)

func TestMPCExample(t *testing.T) {
	secrets := []int{1, 22, 983}
	n := 8
	k := 6
	address := "localhost:12345"
	timeout := 5 * time.Second

	_, allShares, err := mpc.RunMPCExample(address, secrets, n, k)
	if err != nil {
		t.Fatalf("Error running MPC example: %v", err)
	}

	for i := range secrets {
		secretShares := allShares[i]

		fmt.Printf("\nSecret %d\n", i)
		fmt.Printf("Secret shares:\n")
		for _, share := range secretShares {
			fmt.Printf("%0x\n", share)
		}

		combinedSecret, err := client.SendSharesToServer(address, secretShares, timeout)
		if err != nil {
			t.Fatalf("Error combining shares for secret %d: %v", i, err)
		}

		fmt.Printf("Combined secret bytes: %0x\n", combinedSecret)

		// Remove the extra byte at the end before converting to int
		combinedSecret = combinedSecret[:len(combinedSecret)-1]

		reconstructedSecret := mpc.BytesToInt(combinedSecret)
		fmt.Printf("Reconstructed secret %d: %d\n", i, reconstructedSecret)

		if reconstructedSecret != secrets[i] {
			t.Errorf(
				"Reconstructed secret %d does not match the original secret. Expected %d, got %d",
				i,
				secrets[i],
				reconstructedSecret,
			)
		} else {
			fmt.Printf("Reconstructed secret %d matches the original secret\n", i)
		}
	}
}

func TestIntToBytes(t *testing.T) {
	tests := []struct {
		n        int
		expected []byte
	}{
		{1, []byte{0, 0, 0, 1}},
		{255, []byte{0, 0, 0, 255}},
		{65535, []byte{0, 0, 255, 255}},
	}

	for _, test := range tests {
		result := mpc.IntToBytes(test.n)
		if !bytes.Equal(result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestRunMPCExampleError(t *testing.T) {
    address := "wrong-address"
    secrets := []int{1, 22, 983}
    n := -1 // Set n to an invalid value
    k := 6

    _, _, err := mpc.RunMPCExample(address, secrets, n, k)
    if err == nil {
        t.Fatal("Expected an error, but got nil")
    }
}

func TestSendSharesToServerError(t *testing.T) {
    address := "wrong-address"
    secrets := []int{1, 22, 983}
    n := 8
    k := 6

    _, _, err := mpc.RunMPCExample(address, secrets, n, k)
    if err == nil {
        t.Fatal("Expected an error, but got nil")
    }
}
