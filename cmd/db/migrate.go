package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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
		NoMigrateDialector{Dialector: mysql.New(mysql.Config{Conn: conn, DisableDatetimePrecision: true})},
		&gorm.Config{Logger: gLog, NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true,
			SingularTable: true,
		}, QueryFields: true, PrepareStmt: true, SkipDefaultTransaction: true},
	)

	if err != nil {
		return errgo.Wrap(err, "gorm.Open")
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	err = db.AutoMigrate(&dao.User{})
	fmt.Println(strings.Join(Log, "\n"))
	return errgo.Wrap(err, "AutoMigrate")
}

// copy from https://gist.github.com/molind/a67100448b886b7257e30799e06a0718
type NoMigrateDialector struct {
	gorm.Dialector
}

func (d NoMigrateDialector) Migrator(db *gorm.DB) gorm.Migrator {
	return NoMigrateMigrator{Migrator: d.Dialector.Migrator(db)}
}

type NoMigrateMigrator struct {
	gorm.Migrator
}

var Log = []string{}

// Tables
func (m NoMigrateMigrator) CreateTable(dst ...interface{}) error {
	Log = append(Log, "CreateTable: "+fmt.Sprintf("%T", dst))
	return nil
}
func (m NoMigrateMigrator) DropTable(dst ...interface{}) error {
	Log = append(Log, "DropTable: "+fmt.Sprintf("%v", dst))
	return nil
}
func (m NoMigrateMigrator) RenameTable(oldName, newName interface{}) error {
	Log = append(Log, "RenameTable: "+fmt.Sprintf("%v %v", oldName, newName))
	return nil
}

// Columns
func (m NoMigrateMigrator) AddColumn(dst interface{}, field string) error {
	Log = append(Log, "AddColumn: "+fmt.Sprintf("%v %v", dst, field))
	return nil
}
func (m NoMigrateMigrator) DropColumn(dst interface{}, field string) error {
	Log = append(Log, "DropColumn: "+fmt.Sprintf("%v %v", dst, field))
	return nil
}
func (m NoMigrateMigrator) AlterColumn(dst interface{}, field string) error {
	Log = append(Log, "AlterColumn: "+fmt.Sprintf("%v %v", dst, field))
	return nil
}
func (m NoMigrateMigrator) MigrateColumn(dst interface{}, field *schema.Field, columnType gorm.ColumnType) error {
	Log = append(Log, "MigrateColumn: "+fmt.Sprintf("%v %v", dst, field))
	return nil
}
func (m NoMigrateMigrator) RenameColumn(dst interface{}, oldName, field string) error {
	Log = append(Log, "RenameColumn: "+fmt.Sprintf("%v %v %v", dst, oldName, field))
	return nil
}

// Views
func (m NoMigrateMigrator) CreateView(name string, option gorm.ViewOption) error {
	Log = append(Log, "CreateView: "+fmt.Sprintf("%v %v", name, option))
	return nil
}
func (m NoMigrateMigrator) DropView(name string) error {
	Log = append(Log, "DropView: "+fmt.Sprintf("%v", name))
	return nil
}

// Constraints
func (m NoMigrateMigrator) CreateConstraint(dst interface{}, name string) error {
	Log = append(Log, "CreateConstraint: "+fmt.Sprintf("%v %v", dst, name))
	return nil
}
func (m NoMigrateMigrator) DropConstraint(dst interface{}, name string) error {
	Log = append(Log, "DropConstraint: "+fmt.Sprintf("%v %v", dst, name))
	return nil
}

// Indexes
func (m NoMigrateMigrator) CreateIndex(dst interface{}, name string) error {
	Log = append(Log, "CreateIndex: "+fmt.Sprintf("%v %v", dst, name))
	return nil
}
func (m NoMigrateMigrator) DropIndex(dst interface{}, name string) error {
	Log = append(Log, "DropIndex: "+fmt.Sprintf("%v %v", dst, name))
	return nil
}
func (m NoMigrateMigrator) RenameIndex(dst interface{}, oldName, newName string) error {
	Log = append(Log, "RenameIndex: "+fmt.Sprintf("%v %v %v", dst, oldName, newName))
	return nil
}
