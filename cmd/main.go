package main

import (
	"fmt"

	"github.com/LucasRodriguez/mpc_sss/mpc"
)

func main() {
	address := "localhost:50051"
	secrets := []int{5, 7, 3}
	n := 5
	k := 3

	results,_, err := mpc.RunMPCExample(address, secrets, n, k)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Results: %v\n", results)
}