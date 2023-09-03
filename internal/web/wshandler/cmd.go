package wshandler

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gofiber/contrib/websocket"
	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wsreq"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wsres"
	"github.com/sociosarbis/grpc/boilerplate/proto"
	"go.uber.org/zap"
)

type Cmd struct {
	HandlerMod
	client *grpcmod.Client
}

func NewCmd(common *common.Common, client *grpcmod.Client) (*Cmd, error) {
	cmd := Cmd{
		HandlerMod: NewHandlerMod(common),
		client:     client,
	}
	cmd.Register("execute", cmd.execute)
	return &cmd, nil
}

func (cmd *Cmd) execute(c *websocket.Conn, data []byte, out chan any) error {
	req := &wsreq.CmdExecuteReq{}
	err := json.Unmarshal(data, req)
	if err != nil {
		out <- wsres.Err(errcode.BadRequest, "Unmarshal CmdExecuteReq")
		return err
	}
	err = cmd.common.Validate.Struct(req)
	if err != nil {
		out <- wsres.Err(errcode.BadRequest, "Validate CmdExecuteReq")
		return err
	}
	r, err := cmd.client.Cmd.CmdCall(cmd.UserContext(c), &proto.Cmd{
		Script: req.Script,
		Wd:     req.Wd,
	})

	if err != nil {
		out <- wsres.Err(errcode.BadRequest, "CmdCall")
		return err
	}

	index := 0
	for {
		msg, revErr := r.Recv()
		if revErr == nil {
			out <- wsres.Seq(req.Id, index, wsres.Ok(msg))
			index++
		} else {
			err = revErr
			break
		}
	}
	if !errors.Is(err, io.EOF) {
		cmd.common.Logger.Error("client.Recv", zap.Error(err))
		out <- wsres.Seq(req.Id, -1, wsres.Err(errcode.Unknown, "client.Recv"))
		return err
	}
	out <- wsres.Seq(req.Id, -1, wsres.Ok[proto.CmdCallRes](nil))
	return nil
}
