package grpcmod

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type Client struct {
	User proto.UserServiceClient
	Cmd  proto.CmdServiceClient
}

func NewClient(config config.AppConfig) (*Client, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.GRPCHost, config.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errgo.Wrap(err, "grpc.Dial")
	}

	var client Client

	client.User = proto.NewUserServiceClient(conn)

	client.Cmd = proto.NewCmdServiceClient(conn)

	return &client, nil
}
