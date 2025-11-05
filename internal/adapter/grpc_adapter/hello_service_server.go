package grpc_adapter

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func readHeaderFromContext(ctx context.Context) {
	// if md, ok := metadata.FromIncomingContext(ctx); ok {
	// 	for key := range md {
	// 		if len(md[key]) > 0 {
	// 			log.Printf("Metadata Headers from client - %s: %v\n", key, md[key][0])
	// 		}
	// 	}
	// }
}

func (adapter *grpcAdapter) SayHello(ctx context.Context, req *hellov1.SayHelloRequest) (*hellov1.SayHelloResponse, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	readHeaderFromContext(ctx)
	result := adapter.helloService.SayHello(req.GetName())

	response := &hellov1.SayHelloResponse{
		Message:   result,
		CreatedAt: timestamppb.Now(),
	}

	return response, nil
}

func (adapter *grpcAdapter) SayManyHellos(req *hellov1.SayManyHellosRequest, stream grpc.ServerStreamingServer[hellov1.SayManyHellosResponse]) error {

	ctx := stream.Context()
	total := 10
	for i := range total {
		great, err := adapter.helloService.SayHelloWithContext(ctx, req.GetName())

		if err != nil {
			return err
		}

		response := &hellov1.SayManyHellosResponse{
			Message:   fmt.Sprintf("%s - %d", great, i+1),
			CreatedAt: timestamppb.Now(),
		}

		if err := stream.Send(response); err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(500 * time.Millisecond):
		}
	}

	return nil
}

func (adapter *grpcAdapter) SayHelloToEveryone(stream grpc.ClientStreamingServer[hellov1.SayHelloToEveryoneRequest, hellov1.SayHelloToEveryoneResponse]) error {
	var resBuilder strings.Builder
	ctx := stream.Context()
	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&hellov1.SayHelloToEveryoneResponse{
				Message:   resBuilder.String(),
				CreatedAt: timestamppb.Now(),
			})
		}

		if err != nil {
			return err
		}

		great, err := adapter.helloService.SayHelloWithContext(ctx, req.GetName())

		if err != nil {
			return err
		}

		resBuilder.WriteString(great)
		resBuilder.WriteString(" ")

	}

}

func (adapter *grpcAdapter) SayHelloContinuous(stream grpc.BidiStreamingServer[hellov1.SayHelloContinuousRequest, hellov1.SayHelloContinuousResponse]) error {
	ctx := stream.Context()
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		great, err := adapter.helloService.SayHelloWithContext(ctx, req.GetName())

		if err != nil {
			return err
		}

		err = stream.Send(&hellov1.SayHelloContinuousResponse{
			Message:   great,
			CreatedAt: timestamppb.Now(),
		})

		if err != nil {
			return err
		}
	}
}
