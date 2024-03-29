package grpcmod

import (
	"context"

	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	proto.UnimplementedUserServiceServer
	proto.UnimplementedCmdServiceServer
	handler handler.Handler
}

func (s *server) UserDetail(ctx context.Context, req *proto.UserDetailReq) (*proto.UserDetailRes, error) {
	res, err := s.handler.User.Detail(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.Detail")
	}
	return res, nil
}

func (s *server) UserLogin(ctx context.Context, req *proto.UserLoginReq) (*proto.UserLoginRes, error) {
	res, err := s.handler.User.Login(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.Login")
	}
	return res, nil
}

func (s *server) UserMsLogin(ctx context.Context, req *proto.UserMsLoginReq) (*proto.UserLoginRes, error) {
	res, err := s.handler.User.LoginMs(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.LoginMs")
	}
	return res, nil
}

func (s *server) CmdCall(cmdReq *proto.Cmd, srv proto.CmdService_CmdCallServer) error {
	err := s.handler.Cmd.Call(cmdReq, srv)
	if err != nil {
		return errgo.Wrap(err, "server.CmdCall")
	}
	return nil
}

func (s *server) CmdListFolder(ctx context.Context, req *proto.CmdListFolderReq) (*proto.CmdListFolderRes, error) {
	res, err := s.handler.Cmd.ListFolder(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.CmdListFolder")
	}
	return res, nil
}

func (s *server) CmdAdd(ctx context.Context, req *proto.CmdAddReq) (*proto.CmdAddRes, error) {
	res, err := s.handler.Cmd.Add(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.CmdAdd")
	}
	return res, nil
}

func (s *server) CmdUpdate(ctx context.Context, req *proto.CmdUpdateReq) (*emptypb.Empty, error) {
	res, err := s.handler.Cmd.Update(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.CmdUpdate")
	}
	return res, nil
}

func (s *server) CmdDelete(ctx context.Context, req *proto.CmdDeleteReq) (*emptypb.Empty, error) {
	res, err := s.handler.Cmd.Delete(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.CmdDelete")
	}
	return res, nil
}

func (s *server) CmdList(ctx context.Context, req *proto.CmdListReq) (*proto.CmdListRes, error) {
	res, err := s.handler.Cmd.List(ctx, req)
	if err != nil {
		return nil, errgo.Wrap(err, "server.CmdList")
	}
	return res, nil
}
