package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/dal/dao"
	"github.com/sociosarbis/grpc/boilerplate/internal/driver"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Migrate() error {
	var cfg config.AppConfig
	var conn *sql.DB

	err := fx.New(
		fx.NopLogger,
		config.Module,
		fx.Provide(logger.Copy, driver.NewMysqlConnectionPool),
		fx.Populate(&cfg, &conn),
	).Err()

	if err != nil {
		return errgo.Wrap(err, "fx.New")
	}

	var gLog gormLogger.Interface
	logger.Info("enable gorm debug mode, will log all sql")
	gLog = gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormLogger.Config{
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(
		mysql.New(mysql.Config{Conn: conn, DisableDatetimePrecision: true}),
		&gorm.Config{Logger: gLog, DryRun: true, NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true,
			SingularTable: true,
		}, QueryFields: true, SkipDefaultTransaction: true},
	)

	if err != nil {
		return errgo.Wrap(err, "gorm.Open")
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	err = db.AutoMigrate(&dao.User{})
	return errgo.Wrap(err, "AutoMigrate")
}
