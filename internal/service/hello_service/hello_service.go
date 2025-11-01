package hello_service

import (
	"context"
	helloPort "grpc-sample-server/internal/port/hello_service_port"
	"math/rand"
)

type helloService struct{}

func NewHelloService() helloPort.HelloServicePort {
	return &helloService{}
}

func (s *helloService) SayHello(name string) string {

	hello := []string{
		"Hello",
		"Bonjour",
		"Halo",
		"Hola",
		"Ciao",
	}

	helloRand := hello[rand.Intn(len(hello))]
	helloMessage := helloRand + ", " + name + "!"
	return helloMessage
}

func (s *helloService) SayHelloWithContext(ctx context.Context, name string) (string, error) {

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	if name == "error" {
		return "", context.Canceled
	}
	return s.SayHello(name), nil
}
