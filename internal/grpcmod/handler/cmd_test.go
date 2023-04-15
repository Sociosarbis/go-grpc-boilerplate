package handler_test

import (
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/handler"
)

func getCmd(t *testing.T) *handler.Cmd {
	t.Helper()
	return &handler.Cmd{}
}

func TestCmdStart(t *testing.T) {
	t.Parallel()
	cmd := getCmd(t)
	_, stdout, stderr, err := cmd.Start("ls")
	require.NoError(t, err)
	buf := make([]byte, 16)
	errBuf := make([]byte, 16)
	outChan := make(chan []byte)
	defer close(outChan)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var n int
		var err error
		for {
			n, err = stdout.Read(buf)
			if err == nil {
				data := make([]byte, n)
				copy(data, buf[:n])
				outChan <- data
			} else {
				break
			}
		}
		wg.Done()
		require.Equal(t, err, io.EOF)
	}()

	go func() {
		var n int
		var err error
		for {
			n, err = stderr.Read(errBuf)
			if err == nil {
				outChan <- buf[:n]
			} else {
				break
			}
		}
		wg.Done()
		require.Equal(t, err, io.EOF)
	}()

	go func() {
		for buf := range outChan {
			fmt.Printf("%s", buf)
		}
	}()
	wg.Wait()
}
