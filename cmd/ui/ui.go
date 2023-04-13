package ui

import (
	"fmt"
	"os"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "ui",
	Short: "open cli ui page",
	RunE: func(cmd *cobra.Command, args []string) error {
		return start()
	},
}

func init() {
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
