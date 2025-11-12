package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *wrappedStream) Context() context.Context {
	return s.ctx
}

func (s *wrappedStream) RecvMsg(m any) error {
	err := s.ServerStream.RecvMsg(m)

	if err == nil {
		return nil
	}

	return err
}

func LogStreamingInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		md, ok := metadata.FromIncomingContext(ss.Context())

		if !ok {
			md = metadata.New(nil)
		}

		md.Append("LogStreamingInterceptor", "LogStreamingInterceptor Value")

		ctx := metadata.NewIncomingContext(ss.Context(), md)

		wrapped := &wrappedStream{
			ServerStream: ss,
			ctx:          ctx,
		}

		wrapped.SetHeader(metadata.Pairs("streaming", "LogStreamingInterceptor"))

		err := handler(srv, wrapped)
		return err

	}
}
