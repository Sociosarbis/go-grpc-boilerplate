package wshandler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"go.uber.org/zap"
)

type handler func([]byte) (any, error)

type User struct {
	common   *common.Common
	handlers map[string]handler
}

func NewUser(common *common.Common) (*User, error) {
	u := User{
		common:   common,
		handlers: map[string]handler{},
	}
	return &u, nil
}

func (u *User) Handler(c *websocket.Conn) {
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
					if fn, ok := u.handlers[s]; ok {
						res, err := fn(rawData)
						if err != nil {
							u.common.Logger.Error(fmt.Sprintf("%s handler", t), zap.Error(err))
						}
						err = c.WriteJSON(res)
						if err != nil {
							u.common.Logger.Error("writeJSON", zap.Error(err))
						}
					}
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
