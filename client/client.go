package client

import (
	"context"
	"time"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/LucasRodriguez/mpc_sss/proto"
)

func SendSharesToServer(serverAddress string, shares [][]byte, timeout time.Duration) ([]byte, error) {
	// Set a custom dial timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewMPCClient(conn)

	req := &pb.ComputeSumRequest{Shares: shares}
	resp, err := client.ComputeSum(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("compute sum failed: %v", err)
	}

	return resp.Result, nil
}
