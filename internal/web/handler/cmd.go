package handler

import (
	"errors"
	"io"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/slice"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/middleware"
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

func (c *Cmd) Call(ctx *fiber.Ctx) error {
	params, ok := ctx.UserContext().Value(middleware.ParamsCtxKey).(req.CmdCallDto)
	if !ok {
		return res.BadRequest(ctx, errcode.Unknown, "assert req.CmdCallDto")
	}
	r, err := c.client.Cmd.CmdCall(ctx.UserContext(), &proto.Cmd{
		Script: params.Script,
		Wd:     params.Wd,
	})

	if err != nil {
		c.common.Logger.Error("client.Cmd.CmdCall", zap.Error(err))
		return res.GrpcError(ctx, err, "client.Cmd.CmdCall")
	}

	for {
		msg, revErr := r.Recv()
		if revErr == nil {
			writeErr := res.WriteJSON(ctx, msg)
			if writeErr != nil {
				c.common.Logger.Error("res.WriteJSON", zap.Error(writeErr))
				err = writeErr
				break
			}
			_, writeErr = ctx.Write([]byte("\n"))
			if writeErr != nil {
				c.common.Logger.Error("ctx.Write", zap.Error(writeErr))
				err = writeErr
				break
			}
		} else {
			err = revErr
			break
		}
	}
	if !errors.Is(err, io.EOF) {
		c.common.Logger.Error("client.Recv", zap.Error(err))
		return res.GrpcError(ctx, err, "Cmd.CmdCall")
	}
	return nil
}

func (c *Cmd) ListFolder(ctx *fiber.Ctx) error {
	params, ok := ctx.UserContext().Value(middleware.ParamsCtxKey).(req.CmdListFolderDto)
	if !ok {
		return res.BadRequest(ctx, errcode.Unknown, "assert req.CmdListFolderDto")
	}

	r, err := c.client.Cmd.CmdListFolder(ctx.UserContext(), &proto.CmdListFolderReq{
		Folder: params.Prefix,
	})

	if err != nil {
		c.common.Logger.Error("client.Cmd.CmdListFolder", zap.Error(err))
		return res.GrpcError(ctx, err, "client.Cmd.CmdListFolder")
	}

	return res.Ok(ctx, r)
}

func (c *Cmd) Add(ctx *fiber.Ctx) error {
	params, ok := ctx.UserContext().Value(middleware.ParamsCtxKey).(req.CmdAddDto)
	if !ok {
		return res.BadRequest(ctx, errcode.Unknown, "assert req.CmdAddDto")
	}

	r, err := c.client.Cmd.CmdAdd(ctx.UserContext(), &proto.CmdAddReq{
		Items: slice.Map(params.Items, func(item req.CmdItem) *proto.CmdItem {
			return &proto.CmdItem{
				Type:   item.Type,
				Value:  item.Value,
				Config: item.Config,
			}
		}),
	})

	if err != nil {
		c.common.Logger.Error("client.Cmd.CmdAdd", zap.Error(err))
		return res.GrpcError(ctx, err, "client.Cmd.CmdAdd")
	}

	return res.Ok(ctx, r)
}

func (c *Cmd) Update(ctx *fiber.Ctx) error {
	params, ok := ctx.UserContext().Value(middleware.ParamsCtxKey).(req.CmdUpdateDto)
	if !ok {
		return res.BadRequest(ctx, errcode.Unknown, "assert req.CmdUpdateDto")
	}

	r, err := c.client.Cmd.CmdUpdate(ctx.UserContext(), &proto.CmdUpdateReq{
		Id:   params.ID,
		Name: params.Name,
		Items: slice.Map(params.Items, func(item req.CmdItem) *proto.CmdItem {
			return &proto.CmdItem{
				Type:   item.Type,
				Value:  item.Value,
				Config: item.Config,
			}
		}),
	})

	if err != nil {
		c.common.Logger.Error("client.Cmd.CmdAdd", zap.Error(err))
		return res.GrpcError(ctx, err, "client.Cmd.CmdAdd")
	}

	return res.Ok(ctx, r)
}

func (c *Cmd) Delete(ctx *fiber.Ctx) error {
	params, ok := ctx.UserContext().Value(middleware.ParamsCtxKey).(req.CmdDeleteDto)
	if !ok {
		return res.BadRequest(ctx, errcode.Unknown, "assert req.CmdDeleteDto")
	}

	r, err := c.client.Cmd.CmdDelete(ctx.UserContext(), &proto.CmdDeleteReq{
		Id: params.ID,
	})

	if err != nil {
		c.common.Logger.Error("client.Cmd.CmdDelete", zap.Error(err))
		return res.GrpcError(ctx, err, "client.Cmd.CmdDelete")
	}

	return res.Ok(ctx, r)
}

func (c *Cmd) List(ctx *fiber.Ctx) error {
	params, ok := ctx.UserContext().Value(middleware.ParamsCtxKey).(req.CmdListDto)
	if !ok {
		return res.BadRequest(ctx, errcode.Unknown, "assert req.CmdListDto")
	}

	r, err := c.client.Cmd.CmdList(ctx.UserContext(), &proto.CmdListReq{
		Page: params.Page,
		Size: params.Size,
	})

	if err != nil {
		c.common.Logger.Error("client.Cmd.CmdList", zap.Error(err))
		return res.GrpcError(ctx, err, "client.Cmd.CmdList")
	}

	return res.Ok(ctx, r)
}
