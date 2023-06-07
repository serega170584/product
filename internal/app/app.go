package app

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"product/internal/config"
	"product/internal/interceptor"
	"product/internal/proto"
	"product/internal/server"
)

type App struct {
	conf *config.AppConfig
}

func New(config *config.AppConfig) *App {
	return &App{conf: config}
}

func (appInstance *App) Run() {
	conf := appInstance.conf

	lis, err := net.Listen("tcp", net.JoinHostPort(conf.Host, conf.Port))
	if err != nil {
		fmt.Printf("Listen error: %s", err.Error())
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.ServerAuthTokenInterceptor))
	productServer := server.ProductHandlerServer{}
	proto.RegisterProductHandlerServer(grpcServer, productServer)

	if err = grpcServer.Serve(lis); err != nil {
		fmt.Printf("Server error: %s", err.Error())
	}
}
