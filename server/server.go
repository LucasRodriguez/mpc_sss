package server

import (
	"context"
	"net"
	"log"
	"fmt"

	"github.com/hashicorp/vault/shamir"
	"google.golang.org/grpc"

	pb "github.com/LucasRodriguez/mpc_sss/proto"
)

type server struct {
	pb.UnimplementedMPCServer
}

func (s *server) ComputeSum(ctx context.Context, in *pb.ComputeSumRequest) (*pb.ComputeSumResponse, error) {
	shares := in.GetShares()

	// Find the maximum length of shares
	maxLen := 0
	for _, share := range shares {
		if len(share) > maxLen {
			maxLen = len(share)
		}
	}

	// Pad shares with leading zeros to match the maximum length
	paddedShares := make([][]byte, len(shares))
	for i, share := range shares {
		paddedShares[i] = make([]byte, maxLen)
		copy(paddedShares[i][maxLen-len(share):], share)
	}

	result, err := shamir.Combine(paddedShares)

	if err != nil {
		return nil, fmt.Errorf("failed to combine shares: %w", err)
	}

	return &pb.ComputeSumResponse{Result: result}, nil
}

func RunServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("Failed to listen on %s: %v", address, err)
		return err
	}

	s := grpc.NewServer()
	pb.RegisterMPCServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve gRPC server: %v", err)
		return err
	}

	return nil
}