package grpc_adapter

import (
	"context"
	"fmt"
	"log"
	"net"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcAdapter struct {
	hellov1.HelloServiceServer

	grpcPort int
	server   *grpc.Server
}

func NewGrpcAdapter(grpcPort int) *grpcAdapter {
	return &grpcAdapter{
		grpcPort: grpcPort,
	}
}

func (adapter *grpcAdapter) Start() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.grpcPort))

	if err != nil {
		return err
	}

	adapter.server = grpc.NewServer()
	hellov1.RegisterHelloServiceServer(adapter.server, adapter)
	reflection.Register(adapter.server)

	return adapter.server.Serve(listen)

}

func (adapter *grpcAdapter) Stop(context context.Context) {
	select {
	case <-context.Done():
		if adapter.server != nil {
			log.Println("Context done, shutting down gRPC server...")
			adapter.server.Stop()
		}
	default:
		if adapter.server != nil {
			adapter.server.GracefulStop()
		}
	}
}
