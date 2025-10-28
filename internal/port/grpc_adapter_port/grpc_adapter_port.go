package grpc_adapter_port

import "context"

type GrpcAdapterPort interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context)
}
