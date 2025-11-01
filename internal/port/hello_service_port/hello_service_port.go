package hello_service_port

import "context"

type HelloServicePort interface {
	SayHello(name string) string
	SayHelloWithContext(ctx context.Context, name string) (string, error)
}
