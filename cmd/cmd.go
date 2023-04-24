package cmd

import (
	"github.com/spf13/cobra"

	"github.com/sociosarbis/grpc/boilerplate/cmd/ui"
	"github.com/sociosarbis/grpc/boilerplate/cmd/web"
)

var Root = cobra.Command{ //nolint:gochecknoglobals
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   false,
		DisableNoDescFlag:   false,
		DisableDescriptions: false,
		HiddenDefaultCmd:    true,
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() { //nolint:gochecknoinits
	Root.AddCommand(ui.Command, web.Command)
}
