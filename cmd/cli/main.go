package main

import (
	"grpc-sample-server/internal/adapter/logging"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logging.Format{})
}
