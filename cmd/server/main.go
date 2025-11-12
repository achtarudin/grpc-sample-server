package main

import (
	"context"
	"grpc-sample-server/internal/adapter/grpc_adapter"
	"grpc-sample-server/internal/adapter/logging"
	"grpc-sample-server/internal/service/hello_service"
	"grpc-sample-server/internal/utils/console"
	"grpc-sample-server/internal/utils/helper"
	"log"
	"os/signal"
	"syscall"
	"time"

	"buf.build/go/protovalidate"
	"github.com/joho/godotenv"
)

const (
	DEFAULT_GRPC_PORT = 9000
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logging.Format{})

	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file, proceeding with environment variables.")
		return
	}

	// Get environment variables
	grpcPort := helper.GetEnvOrDefault("GRPC_PORT", DEFAULT_GRPC_PORT)

	// Setup graceful shutdown
	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Validator Initialization
	validator, err := protovalidate.New()
	if err != nil {
		log.Printf("Failed to create validator: %v", err)
		stop()
		return
	}

	// Service Initialization
	helloService := hello_service.NewHelloService()

	// Grpc Adapter Initialization
	grpcAdapter := grpc_adapter.NewGrpcAdapter(grpcPort, validator, helloService)

	go func() {
		console.Log("Starting gRPC server on port %d...", grpcPort)
		err := grpcAdapter.Start(shutdown)
		if err != nil {
			log.Printf("gRPC server stopped with error: %v", err)
			stop()
		}
	}()

	<-shutdown.Done()

	// Create a context with timeout for the shutdown process
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	grpcAdapter.Stop(shutdownCtx)
}
