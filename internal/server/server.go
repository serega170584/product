package server

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/internal/proto"
)

type ProductHandlerServer struct {
	proto.UnimplementedProductHandlerServer
}

func (handler ProductHandlerServer) Email(ctx context.Context, request *proto.EmailRequest) (*emptypb.Empty, error) {
	fmt.Println("Listen good")
	return &emptypb.Empty{}, nil
}
