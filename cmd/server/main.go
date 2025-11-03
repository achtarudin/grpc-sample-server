package main

import (
	"context"
	"grpc-sample-server/internal/adapter/grpc_adapter"
	"grpc-sample-server/internal/service/hello_service"
	"grpc-sample-server/internal/utils/console"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"buf.build/go/protovalidate"
)

const (
	DEFAULT_GRPC_PORT = 9000
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	grpcPort := getGrpcPort("GRPC_PORT", DEFAULT_GRPC_PORT)

	// Validator Initialization
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("Failed to create validator: %v", err)
	}
	// Service Initialization
	helloService := hello_service.NewHelloService()

	// Grpc Adapter Initialization
	grpcAdapter := grpc_adapter.NewGrpcAdapter(grpcPort, validator, helloService)

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err != nil {
		stop()
		log.Fatalf("Failed to create validator: %v", err)
	}

	go func() {
		console.Log("Starting gRPC server on port %d...", grpcPort)
		err := grpcAdapter.Start(shutdown)
		if err != nil {
			stop()
		}
	}()

	<-shutdown.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	grpcAdapter.Stop(ctx)
}

func getGrpcPort(key string, defaultPort int) int {
	portStr := os.Getenv(key)
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		return defaultPort
	}
	return portInt
}
