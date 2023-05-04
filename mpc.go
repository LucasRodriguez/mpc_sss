package mpc_sss

import (
	"fmt"
	"math/big"

	"github.com/LucasRodriguez/mpc_sss/client"
	"github.com/LucasRodriguez/mpc_sss/server"
)

func RunMPCExample(address string, secrets []int, n int, k int) ([]int, error) {
	go func() {
		if err := server.RunServer(address); err != nil {
			panic(err)
		}
	}()

	// Give the server some time to start
	time.Sleep(1 * time.Second)

	allShares := make([][][]byte, len(secrets))
	for i, secret := range secrets {
		shares, err := splitSecret(secret, n, k)
		if err != nil {
			return nil, fmt.Errorf("error splitting secret %d: %w", i, err)
		}
		allShares[i] = shares
	}

	results := make([]int, n)
	for i := 0; i < n; i++ {
		partyShares := make([][]byte, len(secrets))
		for j := 0; j < len(secrets); j++ {
			partyShares[j] = allShares[j][i]
		}

		resultBytes, err := client.SendSharesToServer(address, partyShares)
		if err != nil {
			return nil, fmt.Errorf("error sending shares to server: %w", err)
		}

		result := new(big.Int).SetBytes(resultBytes)
		results[i] = int(result.Int64())
	}

	return results, nil
}