package server

import (
	"fmt"
	data "github.com/rusystem/product-data/pkg/gen/data/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	dataServer data.DataServer
}

func New(dataServer data.DataServer) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		dataServer: dataServer,
	}
}

func (s *Server) Run(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	data.RegisterDataServer(s.grpcServer, s.dataServer)

	if err := s.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() func() {
	return s.grpcServer.GracefulStop
}
