package main

import (
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/vault/shamir"
	"github.com/LucasRodriguez/mpc_sss/server"
	"github.com/LucasRodriguez/mpc_sss/client"
)

func main() {
	nodes := 3
	secrets := []int{1, 2}
	n := 3
	k := 2
	address := "localhost:50051"

	sum, receivedSums, err := runMPC(nodes, secrets, n, k, address)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Received sums: %v\n", receivedSums)
}

func runMPC(nodes int, secrets []int, n int, k int, address string) (int, [][][]byte, error) {
    // Start the server in a separate goroutine
    go func() {
        if err := server.RunServer(address); err != nil {
            fmt.Printf("Error running server: %v", err)
        }
    }()

    // Give the server some time to start
    time.Sleep(3 * time.Second)

    // Generate shares for each secret
	allShares := make([][][]byte, len(secrets))
	for i, secret := range secrets {
		secretBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(secretBytes, uint32(secret+1000)) // Add a fixed value to pad the secrets
		shares, err := shamir.Split(secretBytes, n, k)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to split secret %d: %w", secret, err)
		}
		allShares[i] = shares
	}
    // Distribute shares to nodes
    nodeShares := make([][][]byte, nodes)
    for i := range nodeShares {
		fmt.Printf("Node %d shares: %v\n", i+1, secrets)
        nodeShares[i] = make([][]byte, len(secrets))
        for j := range allShares {
            nodeShares[i][j] = allShares[j][i]
        }
    }

    // Each node sends its shares to the server
    receivedSums := make([][][]byte, nodes)
    var wg sync.WaitGroup
    wg.Add(nodes)
    for i, shares := range nodeShares {
        go func(i int, shares [][]byte) {
            defer wg.Done()
            fmt.Printf("Sending shares: %v\n", shares) // <- Print shares being sent
            sums, err := client.SendSharesToServer(address, shares, 30*time.Second)
            if err != nil {
                fmt.Printf("failed to send shares to server: %v\n", err)
            }
            receivedSums[i] = sums
        }(i, shares)
    }
    wg.Wait()

    fmt.Printf("Received sums: %v\n", receivedSums) // <- Print received sums

    // Flatten the received sums
    combinedShares := make([][]byte, 0)
    for _, sums := range receivedSums {
        for _, sum := range sums {
            combinedShares = append(combinedShares, sum)
        }
    }
 // Collect the received sums by secret
    secretSums := make([][][]byte, len(secrets))
    for i := range secretSums {
        secretSums[i] = make([][]byte, 0)
    }

    for _, sums := range receivedSums {
        for i, sum := range sums {
            secretSums[i%len(secrets)] = append(secretSums[i%len(secrets)], sum)
        }
    }

    fmt.Printf("Secret sums: %v\n", secretSums) // <- Print secret sums

    // Combine the received sums and sum the reconstructed secrets
    totalSum := 0
    for i, secretSum := range secretSums {
        combinedSecret, err := shamir.Combine(secretSum)
        if err != nil {
            return 0, nil, fmt.Errorf("failed to combine received sums for secret %d: %w", i, err)
        }

        // Correct the padding and conversion process for the combined secret
        secretBuf := make([]byte, 4)
        copy(secretBuf[4-len(combinedSecret):], combinedSecret)
        secret := int(binary.BigEndian.Uint32(secretBuf))
        secret -= 1000 // Subtract the padding value added earlier

        totalSum += secret
    }

    return totalSum, receivedSums, nil

}