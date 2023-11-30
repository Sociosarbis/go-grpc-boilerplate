package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/dal"
	"github.com/sociosarbis/grpc/boilerplate/internal/driver"
	"github.com/sociosarbis/grpc/boilerplate/internal/gocanal"
	"github.com/sociosarbis/grpc/boilerplate/internal/goredis"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod"
	"github.com/sociosarbis/grpc/boilerplate/internal/jwtgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/web"
	"github.com/sociosarbis/grpc/boilerplate/internal/zookeeper"
)

var Command = &cobra.Command{ //nolint:gochecknoglobals
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
		fx.Provide(
			logger.Copy,
			driver.NewMysqlConnectionPool,
			gocanal.NewCanal,
			zookeeper.NewZookeeper,
			jwtgo.NewJWTManager,
		),
		web.Module,
		dal.Module,
		goredis.Module,
		grpcmod.Module,
		fx.Invoke(func(c *gocanal.Canal) {
			go func() {
				c.Run()
			}()
		}),
		fx.Invoke(func(zk *zookeeper.ZookeeperService) {
			logger.Info("znode", zap.Bool("isMaster", zk.IsMaster.Load()))
		}),
		fx.Populate(&cfg, &app, &grpcSrv, &db),
	).Err()

	if err != nil {
		return errgo.Wrap(dig.RootCause(err), "fx.New")
	}
	return errgo.Wrap(web.Start(cfg, app), "failed to start app")
}
