package main

import (
	"fmt"

	"github.com/sociosarbis/grpc/boilerplate/cmd"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		logger.Fatal("failed to start app:\n" + fmt.Sprintf("\n%+v", err))
	}
}
