package db

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"time"

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

type MigrateOptions struct {
	DryRun       bool
	OutputWriter io.Writer
}

type MigrateLogger struct {
	gormLogger.Interface
	writer io.Writer
}

func (l MigrateLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	if sql[0:6] != "SELECT" {
		l.writer.Write([]byte(fmt.Sprintln(sql)))
	}
}

func Migrate(opts MigrateOptions) error {
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

	gLog := gormLogger.New(
		log.New(io.Discard, "", 0),
		gormLogger.Config{
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(
		mysql.New(mysql.Config{Conn: conn, DisableDatetimePrecision: true}),
		&gorm.Config{Logger: MigrateLogger{Interface: gLog, writer: opts.OutputWriter}, DryRun: opts.DryRun, NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   false,
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
