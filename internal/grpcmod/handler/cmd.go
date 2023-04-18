package handler

import (
	"io"
	"os/exec"
	"sync"
	"syscall"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type Cmd struct {
}

func NewCmd() Cmd {
	return Cmd{}
}

type CmdOptions struct {
	Wd string
}

func (cmd *Cmd) Start(script string, options CmdOptions) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	c := exec.Command("/bin/sh", "-c", script)
	if len(options.Wd) != 0 {
		c.Dir = options.Wd
	}
	// 设置进程组id为子进程id
	c.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	stderr, err := c.StderrPipe()
	if err != nil {
		return nil, nil, nil, errgo.Wrap(err, "Cmd.StderrPipe")
	}

	stdout, err := c.StdoutPipe()
	if err != nil {
		return nil, nil, nil, errgo.Wrap(err, "Cmd.StdoutPipe")
	}

	err = c.Start()
	if err != nil {
		return nil, nil, nil, errgo.Wrap(err, "Cmd.Start")
	}

	return c, stdout, stderr, nil
}

func (cmd *Cmd) Call(cmdReq *proto.Cmd, srv proto.CmdService_CmdCallServer) error {
	_, stdout, stderr, err := cmd.Start(cmdReq.Script, CmdOptions{Wd: cmdReq.Wd})
	if err != nil {
		return status.Errorf(codes.Unknown, "Cmd.Start: %v", err)
	}
	outBuf := make([]byte, 256)
	errBuf := make([]byte, 256)

	wg := sync.WaitGroup{}
	wg.Add(2)

	var ret error

	go func() {
		var n int
		var err error
		for {
			n, err = stdout.Read(outBuf)
			if err == nil {
				sendErr := srv.Send(&proto.CmdCallRes{
					Type:   1,
					Output: string(outBuf[:n]),
				})
				if sendErr != nil {
					ret = sendErr
					break
				}
			} else {
				break
			}
		}
		wg.Done()
	}()

	go func() {
		var n int
		var err error
		for {
			n, err = stderr.Read(errBuf)
			if err == nil {
				sendErr := srv.Send(&proto.CmdCallRes{
					Type:   2,
					Output: string(errBuf[:n]),
				})
				if sendErr != nil {
					ret = sendErr
					break
				}
			} else {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()

	return ret
}
