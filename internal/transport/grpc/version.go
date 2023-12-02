package grpc_api

import (
	"context"
	"time"

	pb "github.com/Red-Sock/rscli_example/pkg/api"
)

type version struct {
	pb.UnimplementedRscliExampleAPIServer
}

func (v version) Version(ctx context.Context, request *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Took: uint32(time.Since(request.ClientTimestamp.AsTime()))}, nil
}
