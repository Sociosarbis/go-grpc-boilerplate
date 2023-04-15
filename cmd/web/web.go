package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/dal"
	"github.com/sociosarbis/grpc/boilerplate/internal/driver"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/jwtgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/web"
)

var Command = &cobra.Command{
	Use:   "web",
	Short: "start web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return start()
	},
}

func start() error {
	var cfg config.AppConfig
	var app *fiber.App
	var grpcSrv *grpc.Server
	var db *gorm.DB

	err := fx.New(
		fx.NopLogger,
		config.Module,
		fx.Provide(logger.Copy, driver.NewMysqlConnectionPool, jwtgo.NewJWTManager),
		web.Module,
		dal.Module,
		grpcmod.Module,
		fx.Populate(&cfg, &app, &grpcSrv, &db),
	).Err()

	if err != nil {
		return errgo.Wrap(dig.RootCause(err), "fx.New")
	}
	return errgo.Wrap(web.Start(cfg, app), "failed to start app")
}
