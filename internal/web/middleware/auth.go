package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
)

func AttachToken(ctx *fiber.Ctx) error {
	value := ctx.Get("Authorization")
	if len(value) != 0 {
		values := strings.Split(value, " ")
		if len(values) == 2 {
			if values[0] == "Bearer" {
				ctx.SetUserContext(metadata.AppendToOutgoingContext(ctx.UserContext(), "authorization", values[1]))
			}
		}
	}
	return ctx.Next()
}
