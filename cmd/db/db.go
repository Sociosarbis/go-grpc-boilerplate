package db

import (
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
		return Migrate()
	},
}

func init() {
	Command.AddCommand(dbMigrate)
}
