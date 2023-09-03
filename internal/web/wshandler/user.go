package wshandler

import (
	"context"
	"encoding/json"

	"github.com/gofiber/contrib/websocket"
	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wsreq"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wsres"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type User struct {
	HandlerMod
	client *grpcmod.Client
}

func NewUser(common *common.Common, client *grpcmod.Client) (*User, error) {
	u := User{
		client:     client,
		HandlerMod: NewHandlerMod(common),
	}

	u.Register("login", u.login)

	return &u, nil
}

func (u *User) login(c *websocket.Conn, data []byte, out chan any) error {
	req := wsreq.UserLoginReq{}
	err := json.Unmarshal(data, &req)
	if err != nil {
		out <- wsres.Err(errcode.Unknown, err.Error())
		return err
	}
	res, err := u.client.User.UserLogin(context.Background(), &proto.UserLoginReq{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		out <- wsres.Err(errcode.Unknown, err.Error())
		return err
	}
	out <- wsres.Ok(res)
	return nil
}
