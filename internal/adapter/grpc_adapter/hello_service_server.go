package grpc_adapter

import (
	"context"
	"fmt"
	"time"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (adapter *grpcAdapter) SayHello(ctx context.Context, req *hellov1.SayHelloRequest) (*hellov1.SayHelloResponse, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	result := adapter.helloService.SayHello(req.GetName())
	now := timestamppb.Now()

	response := &hellov1.SayHelloResponse{
		Message:   result,
		CreatedAt: now,
	}

	return response, nil
}

func (adapter *grpcAdapter) SayManyHellos(req *hellov1.SayManyHellosRequest, stream grpc.ServerStreamingServer[hellov1.SayManyHellosResponse]) error {

	ctx := stream.Context()

	for i := range 100 {
		great, err := adapter.helloService.SayHelloWithContext(ctx, req.GetName())

		if err != nil {
			return err
		}

		now := timestamppb.Now()
		response := &hellov1.SayManyHellosResponse{
			Message:   fmt.Sprintf("%s - %d", great, i+1),
			CreatedAt: now,
		}

		if err := stream.Send(response); err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
		}
	}

	// Selesai dengan sukses
	return nil
}

func (adapter *grpcAdapter) SayHelloToEveryone(grpc.ClientStreamingServer[hellov1.SayHelloToEveryoneRequest, hellov1.SayHelloToEveryoneResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloToEveryone not implemented")
}

func (adapter *grpcAdapter) SayHelloContinuous(grpc.BidiStreamingServer[hellov1.SayHelloContinuousRequest, hellov1.SayHelloContinuousResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloContinuous not implemented")
}
