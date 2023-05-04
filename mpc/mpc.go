package mpc

import (
	"math/big"
	"time"
	"log"
	"github.com/LucasRodriguez/mpc_sss/server"
	"github.com/LucasRodriguez/mpc_sss/client"
	"github.com/hashicorp/vault/shamir"
)

func RunMPCExample(address string, secrets []int, n int, k int) ([]int, [][][]byte, error){
	go func() {
		if err := server.RunServer(address); err != nil {
			log.Printf("Error running server: %v", err)
			panic(err)
		}
	}()

	// Give the server some time to start
	time.Sleep(1 * time.Second)

	allShares := make([][][]byte, len(secrets))
	for i, secret := range secrets {
		secretBigInt := big.NewInt(int64(secret))
		shares, err := shamir.Split(secretBigInt.Bytes(), n, k)
		if err != nil {
			log.Printf("Error splitting secret %d: %v", i, err)
        	return nil, nil, err
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
			log.Printf("Error sending shares to server: %v", err)
        	return nil, nil, err
		}

		result := new(big.Int).SetBytes(resultBytes)
		results[i] = int(result.Int64())
	}

    return results, allShares, nil
}