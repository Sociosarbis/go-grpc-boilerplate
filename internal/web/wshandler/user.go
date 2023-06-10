package wshandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wsreq"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wsres"
	"github.com/sociosarbis/grpc/boilerplate/proto"
	"go.uber.org/zap"
)

type handler func(*websocket.Conn, []byte, chan any) error

type User struct {
	common   *common.Common
	client   *grpcmod.Client
	handlers map[string]handler
}

func NewUser(common *common.Common, client *grpcmod.Client) (*User, error) {
	u := User{
		common:   common,
		client:   client,
		handlers: map[string]handler{},
	}

	u.register("login", u.login)

	return &u, nil
}

func (u *User) Handler(c *websocket.Conn) {
	out := make(chan any)
	defer close(out)
	go func() {
		for msg := range out {
			c.WriteJSON(msg)
		}
	}()
	for {
		if mt, msg, err := c.ReadMessage(); err == nil {
			u.common.Logger.Info("readMessage", zap.Int("messageType", mt))
			data := map[string]any{}
			err = json.Unmarshal(msg, &data)
			if err != nil {
				u.common.Logger.Error("unmarshalMessage", zap.Error(err))
			}
			if t, ok := data["type"]; ok {
				s, ok := t.(string)
				rawData := make([]byte, 0)
				if payload, ok := data["payload"]; ok {
					rawData, err = json.Marshal(payload)
					if err != nil {
						u.common.Logger.Error("marshalPayload", zap.Error(err))
					}
				}
				if ok {
					go func() {
						if fn, ok := u.handlers[s]; ok {
							err := fn(c, rawData, out)
							if err != nil {
								u.common.Logger.Error(fmt.Sprintf("%s handler", t), zap.Error(err))
							}
						}
					}()
				}
			}
		} else {
			break
		}
	}
}

func (u *User) register(event string, fn handler) error {
	u.handlers[event] = fn
	return nil
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
