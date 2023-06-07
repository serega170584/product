package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const defaultToken = "12345678"

func ClientAuthTokenInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	options ...grpc.CallOption,
) error {
	clientMetadata := metadata.Pairs("token", defaultToken)
	ctx = metadata.NewOutgoingContext(ctx, clientMetadata)
	err := invoker(ctx, method, req, reply, cc, options...)
	return err
}
