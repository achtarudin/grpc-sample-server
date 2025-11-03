package grpc_adapter

import (
	"context"
	"fmt"
	grpcPort "grpc-sample-server/internal/port/grpc_adapter_port"
	helloPort "grpc-sample-server/internal/port/hello_service_port"
	"grpc-sample-server/internal/utils/console"
	"net"

	"buf.build/go/protovalidate"
	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcAdapter struct {
	hellov1.UnimplementedHelloServiceServer
	server *grpc.Server

	grpcPort     int
	validator    protovalidate.Validator
	helloService helloPort.HelloServicePort
}

func NewGrpcAdapter(grpcPort int, validator protovalidate.Validator, helloService helloPort.HelloServicePort) grpcPort.GrpcAdapterPort {
	return &grpcAdapter{
		grpcPort:     grpcPort,
		helloService: helloService,
		validator:    validator,
	}
}

func (adapter *grpcAdapter) Start(ctx context.Context) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.grpcPort))

	if err != nil {
		return err
	}

	adapter.server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(protovalidate_middleware.UnaryServerInterceptor(adapter.validator)),
		grpc.ChainStreamInterceptor(protovalidate_middleware.StreamServerInterceptor(adapter.validator)),
	)

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
		console.Log("Gracefully shutting down gRPC server...")
		adapter.server.GracefulStop()
		close(doneChan)
	}()

	select {
	case <-doneChan:
		console.Log("gRPC server stopped gracefully.")
	case <-ctx.Done():
		// Context timeout, paksa berhenti
		console.Log("Shutdown deadline exceeded, forcing gRPC server stop...")
		adapter.server.Stop()
	}
}
