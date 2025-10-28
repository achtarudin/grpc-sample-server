package hello_service

import (
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
