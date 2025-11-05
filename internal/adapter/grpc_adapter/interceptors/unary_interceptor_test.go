package interceptors

import (
	"context"
	"errors"
	"testing"

	hellov1 "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestLogUnaryInterceptor(t *testing.T) {

	md := metadata.Pairs("test", "value")

	ctx := metadata.NewIncomingContext(context.Background(), md)

	req := &hellov1.SayHelloRequest{
		Name: "Encang Cutbray",
	}

	serverInfo := &grpc.UnaryServerInfo{
		FullMethod: "/hellov1.HelloService/SayHello",
	}

	interceptor := LogUnaryInterceptor()

	handler := func(ctx context.Context, req any) (any, error) {
		_, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, errors.New("no metadata found")
		}

		return ctx, nil
	}

	result, err := interceptor(ctx, req, serverInfo, handler)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Implements(t, (*context.Context)(nil), result)

	resultCtx, ok := result.(context.Context)
	assert.True(t, ok)
	assert.NotNil(t, resultCtx)
}

func TestWriteMetadataUnaryInterceptor(t *testing.T) {

	md := metadata.Pairs("test", "value")

	ctx := metadata.NewIncomingContext(context.Background(), md)

	req := &hellov1.SayHelloRequest{
		Name: "Encang Cutbray",
	}

	serverInfo := &grpc.UnaryServerInfo{
		FullMethod: "/hellov1.HelloService/SayHello",
	}

	interceptor := WriteMetadataUnaryInterceptor()

	handler := func(ctx context.Context, req any) (any, error) {
		return ctx, nil
	}

	result, err := interceptor(ctx, req, serverInfo, handler)

	assert.NoError(t, err)
	assert.Implements(t, (*context.Context)(nil), result)

	resultCtx, ok := result.(context.Context)
	assert.True(t, ok)

	mdResult, ok := metadata.FromIncomingContext(resultCtx)
	assert.True(t, ok)

	for key := range mdResult {
		if len(mdResult[key]) > 0 {

			if key == "writemetadataunaryinterceptor" {
				assert.Equal(t, "WriteMetadataUnaryInterceptor Value", mdResult[key][0])
			}

		}
	}
}
