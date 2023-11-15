package redis

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "redis [command]",
	Short: "redis management",
	RunE: func(cmd *cobra.Command, args []string) error {
		print("use --help for more infos")
		return nil
	},
}

var seedCommand = &cobra.Command{
	Use:   "seed [item]",
	Short: "initialize seed data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Seed(args[0])
	},
}

func init() {
	Command.AddCommand(seedCommand)
}
