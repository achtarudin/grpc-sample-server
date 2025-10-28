package main

import (
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	// req := &hellov1.SayHelloRequest{
	// 	Name: "Test",
	// }

	// now := timestamppb.Now()

	// res := &hellov1.SayHelloResponse{
	// 	Message:   "Hello, Test!",
	// 	CreatedAt: now,
	// }

	// log.Printf("Response: %+v\n", res.GetCreatedAt().AsTime().UTC().Format("02/01/2006 15:04:05"))

	// log.Printf("Request: %+v\n", req.GetName())
}
