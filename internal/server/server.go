package server

import (
	"context"
	"fmt"
	"product/internal/proto"
)

type ProductHandlerServer struct {
	proto.UnimplementedProductHandlerServer
}

func (handler ProductHandlerServer) Email(ctx context.Context, request *proto.EmailRequest) (*proto.EmailReply, error) {
	fmt.Printf("Listen good, request: %v \n", request)
	return &proto.EmailReply{Success: true}, nil
}
