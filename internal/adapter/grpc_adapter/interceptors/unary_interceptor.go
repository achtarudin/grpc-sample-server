package interceptors

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func LogUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		// switch t := req.(type) {
		// case *hellov1.SayHelloRequest:
		// 	console.Log("Request Name: %v", t.GetName())
		// default:
		// 	console.Log("Request type: %T", t)
		// }

		// console.Log("Log request %v", req)
		// console.Log("Method: %v", info.FullMethod)

		return handler(ctx, req)
	}
}

func ReadMetadataUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			grpcKey := md.Get("grpc-key")

			if len(grpcKey) > 0 {
				return handler(ctx, req)

			}
			return nil, status.Errorf(codes.Aborted, "grpc-key: %v", errors.New("GRPC_KEY not exists"))
		}

		return nil, status.Errorf(codes.Aborted, "grpc-key: %v", errors.New("GRPC_KEY not exists"))

	}
}
