package client

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"

	pb "github.com/LucasRodriguez/mpc_sss/proto"
)

func SendSharesToServer(address string, shares [][]byte) ([]byte, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewMPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.ComputeSum(ctx, &pb.ComputeSumRequest{Shares: shares})
	if err != nil {
		return nil, errors.New("failed to send shares to server")
	}

	return resp.GetResult(), nil
}