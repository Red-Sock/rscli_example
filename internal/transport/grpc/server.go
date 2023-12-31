package grpc_api

import (
	"context"
	"net"

	"github.com/godverv/matreshka/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Red-Sock/rscli_example/internal/config"
	pb "github.com/Red-Sock/rscli_example/pkg/api"
)

type GrpcServer struct {
	srv *grpc.Server

	networkType string
	address     string
}

func NewServer(cfg config.Config, server *api.GRPC) *GrpcServer {
	srv := grpc.NewServer()

	// Register your servers here
	pb.RegisterRscliExampleAPIServer(srv, &version{})
	return &GrpcServer{
		srv:         srv,
		networkType: "tcp",
		address:     "0.0.0.0:" + server.GetPortStr(),
	}
}

func (s *GrpcServer) Start(_ context.Context) error {
	lis, err := net.Listen(s.networkType, s.address)
	if err != nil {
		return errors.Wrapf(err, "error when tried to listen for %s, %s", s.networkType, s.address)
	}

	err = s.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "error serving grpc")
	}

	return nil
}

func (s *GrpcServer) Stop(_ context.Context) error {
	s.srv.GracefulStop()
	return nil
}
