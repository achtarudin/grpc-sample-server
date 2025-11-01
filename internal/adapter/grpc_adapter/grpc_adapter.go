package grpc_adapter

import (
	"context"
	"fmt"
	grpcPort "grpc-sample-server/internal/port/grpc_adapter_port"
	helloPort "grpc-sample-server/internal/port/hello_service_port"
	"log"
	"net"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcAdapter struct {
	hellov1.UnimplementedHelloServiceServer
	helloService helloPort.HelloServicePort
	grpcPort     int
	server       *grpc.Server
}

func NewGrpcAdapter(grpcPort int, helloService helloPort.HelloServicePort) grpcPort.GrpcAdapterPort {
	return &grpcAdapter{
		grpcPort:     grpcPort,
		helloService: helloService,
	}
}

func (adapter *grpcAdapter) Start(ctx context.Context) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.grpcPort))

	if err != nil {
		return err
	}

	adapter.server = grpc.NewServer()
	hellov1.RegisterHelloServiceServer(adapter.server, adapter)
	reflection.Register(adapter.server)

	serveErr := make(chan error, 1)

	go func() {
		err := adapter.server.Serve(listen)
		serveErr <- err
	}()

	select {
	case err := <-serveErr:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (adapter *grpcAdapter) Stop(ctx context.Context) {
	if adapter.server == nil {
		return
	}

	doneChan := make(chan struct{})

	go func() {
		log.Println("Gracefully shutting down gRPC server...")
		adapter.server.GracefulStop()
		close(doneChan)
	}()

	select {
	case <-doneChan:
		log.Println("gRPC server stopped gracefully.")
	case <-ctx.Done():
		// Context timeout, paksa berhenti
		log.Println("Shutdown deadline exceeded, forcing gRPC server stop...")
		adapter.server.Stop()
	}
}
