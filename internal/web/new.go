package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/middleware"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
)

func Start(c config.AppConfig, app *fiber.App) error {
	addr := c.ListenAddr()
	logger.Info(fmt.Sprintf("http server listening at %s", addr))

	return errgo.Wrap(app.Listen(c.ListenAddr()), "fiber.App.Listen")
}

func New(userHandler *handler.User) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		StrictRouting:         true,
		CaseSensitive:         true,
	})

	app.Use(recover.New())

	app.Use(middleware.AttachToken)

	router := app.Group("/api")

	router.Get("/user/:id", userHandler.Detail)

	app.Use(func(ctx *fiber.Ctx) error {
		return res.NotFound(ctx, errcode.NotFound, "Not Found")
	})

	return app
}
