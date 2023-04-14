package handler

import (
	"io"
	"os/exec"
	"syscall"

	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
)

type Cmd struct {
}

func (cmd *Cmd) Start(script string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	c := exec.Command("/bin/sh", "-c", script)
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
