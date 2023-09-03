package web

import (
	"fmt"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/errcode"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/middleware"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/req"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/res"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wshandler"
)

const throttleTime = time.Second

func Start(c config.AppConfig, app *fiber.App) error {
	addr := c.ListenAddr()
	logger.Info(fmt.Sprintf("http server listening at %s", addr))

	return errgo.Wrap(app.Listen(c.ListenAddr()), "fiber.App.Listen")
}

func New(comm *common.Common, userHandler *handler.User, cmdHandler *handler.Cmd, userWsHandler *wshandler.User) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		StrictRouting:         true,
		CaseSensitive:         true,
	})

	app.Use(recover.New(), cors.New())

	AddRouters(app, comm, userHandler, cmdHandler, userWsHandler)

	return app
}

func NewTestApp(userHandler *handler.User, comm *common.Common, cmdHandler *handler.Cmd, userWsHandler *wshandler.User) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		StrictRouting:         true,
		CaseSensitive:         true,
	})

	AddRouters(app, comm, userHandler, cmdHandler, userWsHandler)

	return app
}

func AddRouters(app *fiber.App, comm *common.Common, userHandler *handler.User, cmdHandler *handler.Cmd, userWsHandler *wshandler.User) {
	perRequestLimiterMiddleware := limiter.New(limiter.Config{
		Max:        1,
		Expiration: throttleTime,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s?ip=%s", c.Route().Path, c.IP())
		},
	})

	app.Use(middleware.AttachToken)

	router := app.Group("/api")

	wsRouter := app.Group("/ws")

	router.Get("/user/:id",
		perRequestLimiterMiddleware,
		middleware.NewValidateReqBuilder(
			comm.Validate,
			middleware.ParamsTypeParams,
			req.UserDetailDto{}).Build(),
		userHandler.Detail)

	router.Post("/user/ms-login",
		perRequestLimiterMiddleware,
		middleware.NewValidateReqBuilder(
			comm.Validate,
			middleware.ParamsTypeBody,
			req.UserMsLoginDto{}).Build(),
		userHandler.LoginMs)

	router.Post("/cmd/call",
		perRequestLimiterMiddleware,
		middleware.NewValidateReqBuilder(
			comm.Validate,
			middleware.ParamsTypeBody,
			req.CmdCallDto{}).Build(),
		cmdHandler.Call)

	router.Get("/cmd/folder",
		perRequestLimiterMiddleware,
		middleware.NewValidateReqBuilder(
			comm.Validate,
			middleware.ParamsTypeQuery,
			req.CmdListFolderDto{}).Build(),
		cmdHandler.ListFolder)

	router.Post("/cmd",
		perRequestLimiterMiddleware,
		middleware.NewValidateReqBuilder(
			comm.Validate,
			middleware.ParamsTypeBody,
			req.CmdAddDto{}).Build(),
		cmdHandler.Add)

	router.Patch("/cmd",
		perRequestLimiterMiddleware,
		middleware.NewValidateReqBuilder(
			comm.Validate,
			middleware.ParamsTypeBody,
			req.CmdUpdateDto{}).Build(),
		cmdHandler.Update)

	router.Delete("/cmd/:id",
		perRequestLimiterMiddleware,
		middleware.NewValidateReqBuilder(
			comm.Validate,
			middleware.ParamsTypeParams,
			req.CmdDeleteDto{}).Build(),
		cmdHandler.Delete)

	router.Get("/cmd/list", perRequestLimiterMiddleware, middleware.NewValidateReqBuilder(
		comm.Validate,
		middleware.ParamsTypeQuery,
		req.CmdListDto{}).Build(),
		cmdHandler.List)

	wsRouter.Get("/user", websocket.New(userWsHandler.Handle))
	app.Use(func(ctx *fiber.Ctx) error {
		return res.NotFound(ctx, errcode.NotFound, "Not Found")
	})
}
