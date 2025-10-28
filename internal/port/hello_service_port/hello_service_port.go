package hello_service_port

type HelloServicePort interface {
	SayHello(name string) string
}
