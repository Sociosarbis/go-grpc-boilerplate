package wshandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type Handler func(*websocket.Conn, []byte, chan any) error

type HandlerMod struct {
	common   *common.Common
	handlers map[string]Handler
}

func NewHandlerMod(common *common.Common) HandlerMod {
	return HandlerMod{
		common:   common,
		handlers: map[string]Handler{},
	}
}

func (m *HandlerMod) Register(event string, fn Handler) error {
	m.handlers[event] = fn
	return nil
}

func (m *HandlerMod) Handle(c *websocket.Conn) {
	out := make(chan any)
	defer close(out)
	go func() {
		for msg := range out {
			c.WriteJSON(msg)
		}
	}()
	for {
		if mt, msg, err := c.ReadMessage(); err == nil {
			m.common.Logger.Info("readMessage", zap.Int("messageType", mt))
			data := map[string]any{}
			err = json.Unmarshal(msg, &data)
			if err != nil {
				m.common.Logger.Error("unmarshalMessage", zap.Error(err))
			}
			if t, ok := data["type"]; ok {
				s, ok := t.(string)
				rawData := make([]byte, 0)
				if payload, ok := data["payload"]; ok {
					rawData, err = json.Marshal(payload)
					if err != nil {
						m.common.Logger.Error("marshalPayload", zap.Error(err))
					}
				}
				if ok {
					go func() {
						if fn, ok := m.handlers[s]; ok {
							err := fn(c, rawData, out)
							if err != nil {
								m.common.Logger.Error(fmt.Sprintf("%s handler", t), zap.Error(err))
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

func (m *HandlerMod) UserContext(c *websocket.Conn) context.Context {
	token, ok := c.Locals("authorization").(string)
	ctx := context.Background()
	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	}
	return ctx
}
