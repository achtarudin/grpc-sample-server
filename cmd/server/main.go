package main

import (
	"context"
	"grpc-sample-server/internal/adapter/grpc_adapter"
	"grpc-sample-server/internal/service/hello_service"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	DEFAULT_GRPC_PORT = 9000
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	grpcPort := getGrpcPort("GRPC_PORT", DEFAULT_GRPC_PORT)

	// Service Initialization
	helloService := hello_service.NewHelloService()

	// Grpc Adapter Initialization
	grpcAdapter := grpc_adapter.NewGrpcAdapter(grpcPort, helloService)

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Starting gRPC server on port %d...\n", grpcPort)
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
