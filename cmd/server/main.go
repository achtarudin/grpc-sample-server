package main

import (
	"context"
	"grpc-sample-server/internal/adapter/grpc_adapter"
	"log"
	"os/signal"
	"syscall"
)

const (
	PORT = 9000
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	// Grpc Adapter Initialization
	grpcAdapter := grpc_adapter.NewGrpcAdapter(PORT)

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Starting gRPC server on port %d...\n", PORT)
		err := grpcAdapter.Start()
		if err != nil {
			log.Printf("Failed to start gRPC server: %v\n", err)
			stop()
		}
	}()

	<-shutdown.Done()
	log.Printf("Shutting down gRPC server on port %d...\n", PORT)
	grpcAdapter.Stop(shutdown)
}
