package server

import (
	"context"
	"product/internal/proto"
	servererr "product/internal/server/error"
)

type ProductHandlerServer struct {
	proto.UnimplementedProductHandlerServer
}

func (handler ProductHandlerServer) Email(ctx context.Context, request *proto.EmailRequest) (*proto.EmailReply, error) {
	err := validate(request)
	if err != nil {
		return &proto.EmailReply{Success: false}, err
	}
	return &proto.EmailReply{Success: true}, nil
}

func validate(request *proto.EmailRequest) error {
	var err error

	if request.GetTo() == nil {
		err = servererr.New("To", nil, err)
	}

	if request.GetBodyType() == proto.EmailRequest__UNSPECIFIED {
		err = servererr.New("To", request.GetBodyType(), err)
	}

	return err
}
