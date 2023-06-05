package client

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"product/internal/proto"
)

func NewClient(host string, port string) proto.ProductHandlerClient {
	conn, err := grpc.Dial(net.JoinHostPort(host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Close connection error: %s", err.Error())
		}
	}(conn)

	client := proto.NewProductHandlerClient(conn)

	return client
}
