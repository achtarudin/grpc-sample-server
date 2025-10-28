package grpc_adapter

import (
	"context"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (adapter *grpcAdapter) SayHello(ctx context.Context, req *hellov1.SayHelloRequest) (*hellov1.SayHelloResponse, error) {

	result := adapter.helloService.SayHello(req.Name)
	now := timestamppb.Now()

	response := &hellov1.SayHelloResponse{
		Message:   result,
		CreatedAt: now,
	}

	return response, nil
}

func (adapter *grpcAdapter) SayManyHellos(*hellov1.SayManyHellosRequest, grpc.ServerStreamingServer[hellov1.SayManyHellosResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayManyHellos not implemented")
}

func (adapter *grpcAdapter) SayHelloToEveryone(grpc.ClientStreamingServer[hellov1.SayHelloToEveryoneRequest, hellov1.SayHelloToEveryoneResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloToEveryone not implemented")
}

func (adapter *grpcAdapter) SayHelloContinuous(grpc.BidiStreamingServer[hellov1.SayHelloContinuousRequest, hellov1.SayHelloContinuousResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloContinuous not implemented")
}
