package grpcmod

import (
	"context"

	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type server struct {
	proto.UnimplementedUserServiceServer
	handler handler.Handler
}

func (s *server) UserDetail(ctx context.Context, req *proto.UserDetailReq) (*proto.UserDetailRes, error) {
	res, err := s.handler.User.Detail(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.Detail")
	}
	return res, nil
}
