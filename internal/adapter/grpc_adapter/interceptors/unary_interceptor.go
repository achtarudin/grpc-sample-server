package interceptors

import (
	"context"
	"grpc-sample-server/internal/utils/console"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"google.golang.org/grpc"
)

func LogUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		switch t := req.(type) {
		case *hellov1.SayHelloRequest:
			console.Log("Request Name: %v", t.GetName())
		default:
			console.Log("Request type: %T", t)
		}

		console.Log("Log request %v", req)
		console.Log("Method: %v", info.FullMethod)

		return handler(ctx, req)
	}
}
