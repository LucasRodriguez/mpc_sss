package server

import (
	"context"
	"errors"
	"net"

	"github.com/hashicorp/vault/shamir"
	"google.golang.org/grpc"

	pb "github.com/LucasRodriguez/mpc_sss/proto"
)

type server struct {
	pb.UnimplementedMPCServer
}

func (s *server) ComputeSum(ctx context.Context, in *pb.ComputeSumRequest) (*pb.ComputeSumResponse, error) {
	shares := in.GetShares()
	result, err := shamir.Combine(shares)

	if err != nil {
		return nil, errors.New("failed to combine shares")
	}

	return &pb.ComputeSumResponse{Result: result}, nil
}

func RunServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterMPCServer(s, &server{})
	return s.Serve(lis)
}