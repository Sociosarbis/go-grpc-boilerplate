package test

import (
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/dal"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/jwtgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/web"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
	"go.uber.org/fx"
)

type Mock struct {
	GrpcClient *grpcmod.Client
}

const throttleTime = time.Second

func GetWebApp(tb testing.TB, m Mock) *fiber.App {
	var app *fiber.App

	err := fx.New(
		fx.NopLogger,
		config.Module,
		fx.Provide(logger.Copy, jwtgo.NewJWTManager),
		fx.Provide(web.NewTestApp, common.New, handler.NewUser),
		dal.Module,
		fx.Supply(m.GrpcClient),
		fx.Populate(&app),
	).Err()

	if err != nil {
		tb.Fatal("can't create web app", err)
	}

	return app
}
