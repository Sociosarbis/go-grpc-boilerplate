package dal

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
)

func NewDB(conn *sql.DB, c config.AppConfig) (*gorm.DB, error) {
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
		&gorm.Config{Logger: gLog, NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   false,
			SingularTable: true,
		}, QueryFields: true, PrepareStmt: true, SkipDefaultTransaction: true},
	)
	if err != nil {
		return nil, errgo.Wrap(err, "gorm.Open")
	}

	return db, nil
}
