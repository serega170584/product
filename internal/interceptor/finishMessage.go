package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	fmt.Printf("Invoked rpc method: %s, duration: %s, err: %v", method, time.Since(start), err)
	return err
}
