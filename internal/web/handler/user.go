package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/middleware"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/req"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type User struct {
	common *common.Common
	client *grpcmod.Client
}

func NewUser(common *common.Common, client *grpcmod.Client) (*User, error) {
	return &User{
		common,
		client,
	}, nil
}

func (u *User) Detail(ctx *fiber.Ctx) error {
	params, ok := ctx.UserContext().Value(middleware.ParamsCtxKey).(req.UserDetailDto)
	if !ok {
		return res.BadRequest(ctx, errcode.Unknown, "assert req.UserDetailDto")
	}
	r, err := u.client.User.UserDetail(ctx.UserContext(), &proto.UserDetailReq{
		Id: params.ID,
	})

	if err != nil {
		u.common.Logger.Error("client.User.UserDetail", zap.Error(err))
		return res.InternalError(ctx, errcode.Unknown, "client.User.UserDetail")
	}
	return res.Ok(ctx, r)
}
