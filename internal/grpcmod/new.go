package grpcmod

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/interceptor"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

func New(handler handler.Handler, authInterceptor *interceptor.AuthInterceptor) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", "127.0.0.1:4040")
	if err != nil {
		return nil, errgo.Wrap(err, "net.Listen")
	}

	authInterceptor.
		SkipReq(&proto.UserLoginReq{}).
		SkipReq(&proto.UserMsLoginReq{})

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Auth),
		grpc.StreamInterceptor(authInterceptor.AuthStream),
	)

	srvImpl := server{
		proto.UnimplementedUserServiceServer{},
		proto.UnimplementedCmdServiceServer{},
		handler,
	}

	proto.RegisterUserServiceServer(srv, &srvImpl)

	proto.RegisterCmdServiceServer(srv, &srvImpl)

	go func() {
		err = srv.Serve(lis)
		if err != nil {
			logger.Fatal("grpc.Server.Serve", zap.Error(err))
		}
	}()

	return srv, nil
}
