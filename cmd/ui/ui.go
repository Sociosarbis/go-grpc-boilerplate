package ui

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
)

var Command = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "ui",
	Short: "open cli ui page",
	RunE: func(cmd *cobra.Command, args []string) error {
		return start()
	},
}

func init() { //nolint:gochecknoinits
	fmt.Println("ui cmd init")
}

func start() error {
	wd, err := os.Getwd()
	if err != nil {
		return errgo.Wrap(err, "os.Getwd")
	}
	fmt.Println(wd)
	return nil
}
