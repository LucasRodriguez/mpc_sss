package client

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	pb "github.com/LucasRodriguez/mpc_sss/proto"
)

// Define a test gRPC server to handle ComputeSum requests
type testServer struct {
	pb.UnimplementedMPCServer
	shouldFail bool
}

func (s *testServer) ComputeSum(ctx context.Context, in *pb.ComputeSumRequest) (*pb.ComputeSumResponse, error) {
	if s.shouldFail {
		return nil, fmt.Errorf("compute sum failed")
	}
	return &pb.ComputeSumResponse{Result: []byte{1, 2, 3}}, nil
}

func startTestServer(shouldFail bool) string {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatalf("Failed to start test server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMPCServer(grpcServer, &testServer{shouldFail: shouldFail})
	go grpcServer.Serve(lis)

	return lis.Addr().String()
}

func TestSendSharesToServer(t *testing.T) {
	// Start a test gRPC server
	testServerAddress := startTestServer(false)

	// Define your test case inputs and expected outputs
	testShares := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	}

	expectedResult := []byte{1, 2, 3}

	// Call the SendSharesToServer function
	result, err := SendSharesToServer(testServerAddress, testShares, time.Second)

	// Verify the result and error
	assert.NoError(t, err, "SendSharesToServer should not return an error")
	assert.Equal(t, expectedResult, result, "SendSharesToServer should return the expected result")
}

func TestSendSharesToServer_ConnectionFailure(t *testing.T) {
	// Pass an invalid address to trigger a connection failure
	invalidAddress := "invalid_address"

	testShares := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	}

	// Call the SendSharesToServer function
	result, err := SendSharesToServer(invalidAddress, testShares, time.Second)

	// Verify the result and error
	assert.Error(t, err, "SendSharesToServer should return an error when the connection fails")
	assert.Nil(t, result, "SendSharesToServer should return nil result when the connection fails")
}

func TestSendSharesToServer_ComputeSumFailure(t *testing.T) {
	// Start a test gRPC server that will fail the ComputeSum call
	testServerAddress := startTestServer(true)

	testShares := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	}

	// Call the SendSharesToServer function
	result, err := SendSharesToServer(testServerAddress, testShares, time.Second)

	// Verify the result and error
	assert.Error(t, err, "SendSharesToServer should return an error when the ComputeSum call fails")
	assert.Nil(t, result, "SendSharesToServer should return nil result when the ComputeSum call fails")
}