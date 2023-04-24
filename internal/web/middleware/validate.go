package middleware

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
)

type ParamsType uint8

const (
	ParamsTypeBody ParamsType = iota + 1
	ParamsTypeQuery
	ParamsTypeParams
)

var ParamsCtxKey = struct{}{} //nolint:gochecknoglobals

type ValidateReqBuilder[T any] struct {
	validate   *validator.Validate
	paramsType ParamsType
	params     T
}

func NewValidateReqBuilder[T any](validate *validator.Validate, paramsType ParamsType, params T) ValidateReqBuilder[T] {
	return ValidateReqBuilder[T]{
		validate,
		paramsType,
		params,
	}
}

func (builder ValidateReqBuilder[T]) Build() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		params := builder.params
		var err error
		switch builder.paramsType {
		case ParamsTypeBody:
			err = ctx.BodyParser(&params)
		case ParamsTypeParams:
			err = ctx.ParamsParser(&params)
		case ParamsTypeQuery:
			err = ctx.QueryParser(&params)
		}
		if err != nil {
			return res.BadRequest(ctx, errcode.Unknown, "parse req")
		}

		err = builder.validate.Struct(&params)
		if err != nil {
			msg := "validate.Struct"
			errs := make(validator.ValidationErrors, 0)
			if errors.As(err, &errs) {
				msg = errs.Error()
			}
			return res.BadRequest(ctx, errcode.Unknown, msg)
		}
		ctx.SetUserContext(context.WithValue(ctx.UserContext(), ParamsCtxKey, params))
		return ctx.Next()
	}
}
