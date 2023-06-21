package interceptor

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const defaultRightToken = "123456789"

func ServerAuthTokenInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	fmt.Printf("Method: %s \n", info.FullMethod)

	if err := authorize(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func authorize(ctx context.Context) error {
	clientMetata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("auth header error")
	}

	authHeader, ok := clientMetata["token"]
	if !ok {
		return errors.New("auth token getting error")
	}

	token := authHeader[0]
	if token != defaultRightToken {
		return errors.New("wrong auth token")
	}

	return nil
}
