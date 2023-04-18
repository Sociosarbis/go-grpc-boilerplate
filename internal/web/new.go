package web

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/middleware"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
)

const throttleTime = time.Second

func Start(c config.AppConfig, app *fiber.App) error {
	addr := c.ListenAddr()
	logger.Info(fmt.Sprintf("http server listening at %s", addr))

	return errgo.Wrap(app.Listen(c.ListenAddr()), "fiber.App.Listen")
}

func New(userHandler *handler.User, cmdHandler *handler.Cmd) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		StrictRouting:         true,
		CaseSensitive:         true,
	})

	app.Use(recover.New())

	AddRouters(app, userHandler, cmdHandler)

	return app
}

func NewTestApp(userHandler *handler.User, cmdHandler *handler.Cmd) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		StrictRouting:         true,
		CaseSensitive:         true,
	})

	AddRouters(app, userHandler, cmdHandler)

	return app
}

func AddRouters(app *fiber.App, userHandler *handler.User, cmdHandler *handler.Cmd) {
	perRequestLimiterMiddleware := limiter.New(limiter.Config{
		Max:        1,
		Expiration: throttleTime,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s?ip=%s", c.Route().Path, c.IP())
		},
	})

	app.Use(middleware.AttachToken)

	router := app.Group("/api")

	router.Get("/user/:id", perRequestLimiterMiddleware, userHandler.Detail)

	router.Post("/cmd", perRequestLimiterMiddleware, cmdHandler.Call)

	app.Use(func(ctx *fiber.Ctx) error {
		return res.NotFound(ctx, errcode.NotFound, "Not Found")
	})
}
