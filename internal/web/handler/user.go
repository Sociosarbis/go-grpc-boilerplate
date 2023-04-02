package handler

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
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
	var params req.UserDetailDto
	err := ctx.ParamsParser(&params)
	if err != nil {
		return res.BadRequest(ctx, errcode.Unknown, "fiber.Ctx.ParamsParser")
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
	r, err := u.client.User.UserDetail(ctx.UserContext(), &proto.UserDetailReq{
		Id: params.ID,
	})

	if err != nil {
		return res.InternalError(ctx, errcode.Unknown, "client.User.UserDetail")
	}
	return res.Ok(ctx, r)
}
