package driver

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
)

var setLoggerOnce = sync.Once{}

const pingTimeout = time.Second * 5

func NewMysqlConnectionPool(c config.AppConfig) (*sql.DB, error) {
	const maxIdleTime = time.Hour * 6

	setLoggerOnce.Do(func() {
		_ = mysql.SetLogger(logger.StdAt(zap.ErrorLevel))
	})

	logger.Info(fmt.Sprintf("creating sql connection pool with size %d", c.MySQLMaxConn))

	u := mysql.NewConfig()
	u.User = c.MySQLUserName
	u.Passwd = c.MySQLPassword
	u.Net = "tcp"
	u.Addr = net.JoinHostPort(c.MySQLHost, c.MySQLPort)
	u.DBName = c.MySQLDatabase
	u.Loc = time.UTC
	u.ParseTime = true

	db, err := sql.Open("mysql", u.FormatDSN())
	if err != nil {
		return nil, errgo.Wrap(err, "sql.open")
	}

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, errgo.Wrap(err, "sql.DB.PingContext")
	}

	db.SetMaxOpenConns(c.MySQLMaxConn)
	// default mysql has 7 hour timeout
	db.SetConnMaxIdleTime(maxIdleTime)

	return db, nil
}
