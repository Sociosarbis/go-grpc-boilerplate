package grpcmod

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type Client struct {
	User proto.UserServiceClient
	Cmd  proto.CmdServiceClient
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("127.0.0.1:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errgo.Wrap(err, "grpc.Dial")
	}

	var client Client

	client.User = proto.NewUserServiceClient(conn)

	client.Cmd = proto.NewCmdServiceClient(conn)

	return &client, nil
}
