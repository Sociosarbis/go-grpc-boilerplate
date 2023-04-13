package cmd

import (
	"github.com/sociosarbis/grpc/boilerplate/cmd/ui"
	"github.com/sociosarbis/grpc/boilerplate/cmd/web"
	"github.com/spf13/cobra"
)

var Root = cobra.Command{
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   false,
		DisableNoDescFlag:   false,
		DisableDescriptions: false,
		HiddenDefaultCmd:    true,
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	Root.AddCommand(ui.Command, web.Command)
}
