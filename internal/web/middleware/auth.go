package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
)

func attachToken(ctx *fiber.Ctx, token string) {
	ctx.Context().SetUserValue("authorization", token)
	ctx.SetUserContext(metadata.AppendToOutgoingContext(ctx.UserContext(), "authorization", token))
}

func AttachToken(ctx *fiber.Ctx) error {
	value := ctx.Get("Authorization")
	token := ctx.Query("token")
	if len(value) != 0 {
		values := strings.Split(value, " ")
		if len(values) == 2 {
			if values[0] == "Bearer" {
				attachToken(ctx, values[1])
			}
		}
	} else if len(token) != 0 {
		attachToken(ctx, token)
	}
	return ctx.Next()
}
