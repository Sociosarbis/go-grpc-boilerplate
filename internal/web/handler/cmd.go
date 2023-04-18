package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/req"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type Cmd struct {
	common *common.Common
	client *grpcmod.Client
}

func NewCmd(common *common.Common, client *grpcmod.Client) (*Cmd, error) {
	return &Cmd{
		common,
		client,
	}, nil
}

func (u *Cmd) Call(ctx *fiber.Ctx) error {
	var params req.CmdCallDto
	err := ctx.BodyParser(&params)
	if err != nil {
		return res.BadRequest(ctx, errcode.Unknown, "fiber.Ctx.BodyParser")
	}

	err = u.common.Validate.Struct(&params)
	if err != nil {
		msg := "validate.Struct"
		errs := make(validator.ValidationErrors, 0)
		if errors.As(err, &errs) {
			msg = errs.Error()
		}
		return res.BadRequest(ctx, errcode.Unknown, msg)
	}
	r, err := u.client.Cmd.CmdCall(ctx.UserContext(), &proto.Cmd{
		Script: params.Script,
		Wd:     params.Wd,
	})

	if err != nil {
		u.common.Logger.Error("client.Cmd.CmdCall", zap.Error(err))
		return res.InternalError(ctx, errcode.Unknown, "client.Cmd.CmdCall")
	}

	for {
		msg, revErr := r.Recv()
		if revErr == nil {
			writeErr := res.WriteJSON(ctx, msg)
			ctx.Write([]byte("\n"))
			if writeErr != nil {
				u.common.Logger.Error("res.WriteJSON", zap.Error(writeErr))
				err = writeErr
				break
			}
		} else {
			err = revErr
			break
		}
	}
	if !errors.Is(err, io.EOF) {
		u.common.Logger.Error("client.Recv", zap.Error(err))
		return res.InternalError(ctx, http.StatusInternalServerError, "Cmd.CmdCall")
	}
	return nil
}
