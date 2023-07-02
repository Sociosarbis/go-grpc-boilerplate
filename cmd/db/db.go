package db

import (
	"io"
	"os"
	"path"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "db [command]",
	Short: "db management",
	RunE: func(cmd *cobra.Command, args []string) error {
		print("use --help for more infos")
		return nil
	},
}

var dbMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "migrate db",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Migrate(MigrateOptions{DryRun: false, OutputWriter: os.Stdout})
	},
}

var dbMigrateGenerate = &cobra.Command{
	Use:   "generate [output path]",
	Short: "generate migration sqls",
	RunE: func(cmd *cobra.Command, args []string) error {
		var writer io.Writer
		if len(args) != 0 {
			cwd, err := os.Getwd()
			if err != nil {
				return errgo.Wrap(err, "Getwd")
			}
			var filePath string
			if path.IsAbs(args[0]) {
				filePath = args[0]
			} else {
				filePath = path.Join(cwd, args[0])
			}
			var fileDir = path.Dir(filePath)
			err = os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				return errgo.Wrap(err, "MkdirAll")
			}
			writer, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
			if err != nil {
				return errgo.Wrap(err, "OpenFile")
			}
		} else {
			writer = os.Stdout
		}
		return Migrate(MigrateOptions{
			DryRun:       true,
			OutputWriter: writer,
		})
	},
}

func init() {
	dbMigrate.AddCommand(dbMigrateGenerate)
	Command.AddCommand(dbMigrate)
}
