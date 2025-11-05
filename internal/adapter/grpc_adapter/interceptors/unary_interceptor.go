package interceptors

import (
	"context"
	"grpc-sample-server/internal/utils/console"
	"log"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		switch t := req.(type) {
		case *hellov1.SayHelloRequest:
			console.Log("Request type: %T, Name: %s", t, t.GetName())
		default:
			console.Log("Request type: %T", t)
		}

		console.Log("Log request %v", req)
		console.Log("Method: %v", info.FullMethod)
		return handler(ctx, req)
	}
}

func WriteMetadataUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			md = metadata.New(nil)
		}

		md.Append("WriteMetadataUnaryInterceptor", "WriteMetadataUnaryInterceptor Value")

		ctx = metadata.NewIncomingContext(ctx, md)

		return handler(ctx, req)
	}
}

func ReadMetadataUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for key := range md {
				if len(md[key]) > 0 {
					log.Printf("Read Metadata Headers - %s: %v\n", key, md[key][0])
				}
			}
		}

		// Call the handler to complete the normal execution of the RPC.
		resp, err = handler(ctx, req)

		return resp, err
	}
}
