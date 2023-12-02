package signalgo

import (
	"os"
	"os/signal"
	"syscall"
)

func OnShutdown() chan os.Signal {
	endChan := make(chan os.Signal, 1)
	signal.Notify(endChan, os.Interrupt, syscall.SIGTERM)
	return endChan
}
