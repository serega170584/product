package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"os"
	"product/internal/config"
	"product/internal/interceptor"
	"product/internal/proto"
	"time"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Printf("Usage: %s <subject> <body> <bodyType> <to> ...\n", os.Args[0])
		os.Exit(1)
	}

	inputBodyFormat := os.Args[3]
	bodyFormatId, ok := proto.EmailRequest_BodyFormat_value[inputBodyFormat]
	if !ok {
		fmt.Printf("Body format values: %v", proto.EmailRequest_BodyFormat_value)
		os.Exit(1)
	}

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(
		net.JoinHostPort(conf.App.Host, conf.App.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.FinishMessageInterceptor),
		grpc.WithUnaryInterceptor(interceptor.ClientAuthTokenInterceptor),
	)
	if err != nil {
		panic(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Close connection error: %s", err.Error())
		}
	}(conn)

	grpcClient := proto.NewProductHandlerClient(conn)

	clientContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	to := os.Args[4:]

	_, err = grpcClient.Email(clientContext, &proto.EmailRequest{
		Subject:  os.Args[1],
		Body:     os.Args[2],
		BodyType: proto.EmailRequest_BodyFormat(bodyFormatId),
		To:       to,
	})
	if err != nil {
		fmt.Printf("Send email error: %s", err.Error())
	}
}
