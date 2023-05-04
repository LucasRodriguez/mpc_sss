package client

import (
	"context"
	"time"
	"log"

	"google.golang.org/grpc"

	pb "github.com/LucasRodriguez/mpc_sss/proto"
)

func SendSharesToServer(address string, shares [][]byte) ([]byte, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to dial gRPC server: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := pb.NewMPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.ComputeSum(ctx, &pb.ComputeSumRequest{Shares: shares})
	if err != nil {
		log.Printf("Failed to send shares to server: %v", err)
		return nil, err
	}

	return resp.GetResult(), nil
}