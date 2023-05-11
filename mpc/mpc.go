package mpc

import (
	"math/big"
	"encoding/binary"
	"time"
	"log"
	"github.com/LucasRodriguez/mpc_sss/server"
	"github.com/LucasRodriguez/mpc_sss/client"
	"github.com/hashicorp/vault/shamir"
)

func RunMPCExample(address string, secrets []int, n int, k int) ([]int, [][][]byte, error) {
	go func() {
		if err := server.RunServer(address); err != nil {
			log.Printf("Error running server: %v", err)
		}
	}()

	// Give the server some time to start
	time.Sleep(3 * time.Second)

	allShares := make([][][]byte, len(secrets))
    for i, secret := range secrets {
        secretBytes := append(big.NewInt(int64(secret)).Bytes(), byte(i))
        shares, err := shamir.Split(secretBytes, n, k)
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

        resultBytes, err := client.SendSharesToServer(address, partyShares, time.Second)
        if err != nil {
            log.Printf("Error sending shares to server: %v", err)
            return nil, nil, err
        }

        result := new(big.Int).SetBytes(resultBytes[:len(resultBytes)-1])
        results[i] = int(result.Int64())
    }

    return results, allShares, nil
}

func BytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(append(make([]byte, 4-len(b)), b...)))
}

func IntToBytes(n int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	return buf
}