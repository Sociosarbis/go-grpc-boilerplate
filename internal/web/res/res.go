package res

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *T     `json:"data,omitempty"`
}

type empty struct{}

func Ok[T any](ctx *fiber.Ctx, data *T) error {
	return ctx.JSON(Response[T]{
		0,
		"",
		data,
	})
}

func Err(ctx *fiber.Ctx, status int, code int, msg string) error {
	return ctx.Status(status).JSON(Response[empty]{
		code,
		msg,
		nil,
	})
}

func BadRequest(ctx *fiber.Ctx, code int, msg string) error {
	return Err(ctx, http.StatusBadRequest, code, msg)
}

func NotFound(ctx *fiber.Ctx, code int, msg string) error {
	return Err(ctx, http.StatusNotFound, code, msg)
}

func InternalError(ctx *fiber.Ctx, code int, msg string) error {
	return Err(ctx, http.StatusInternalServerError, code, msg)
}

func WriteJSON[T any](ctx *fiber.Ctx, msg T) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return errgo.Wrap(err, "json.Marshal")
	}
	_, err = ctx.Write(data)
	if err != nil {
		return errgo.Wrap(err, "ctx.Write")
	}
	return nil
}
