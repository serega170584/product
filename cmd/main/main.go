package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"product/internal/config"
	"product/internal/proto"
	"product/internal/server"
)

func main() {

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", net.JoinHostPort(conf.App.Host, conf.App.Port))
	if err != nil {
		fmt.Printf("Listen error: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	productServer := server.ProductHandlerServer{}
	proto.RegisterProductHandlerServer(grpcServer, productServer)

	if err = grpcServer.Serve(lis); err != nil {
		fmt.Printf("Server error: %s", err.Error())
	}
}
