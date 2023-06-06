package messageInterceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

func FinishMessageInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	grpclog.Info("Invoked rpc method: %s, duration: %v, err: %s", method, time.Since(start), err.Error())
	return err
}
