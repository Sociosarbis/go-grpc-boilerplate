package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"sync"
	"syscall"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/sociosarbis/grpc/boilerplate/internal/ctxkey"
	"github.com/sociosarbis/grpc/boilerplate/internal/dal/dao"
	"github.com/sociosarbis/grpc/boilerplate/internal/jwtgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/slice"
	"github.com/sociosarbis/grpc/boilerplate/proto"
)

type Cmd struct {
	db *gorm.DB
}

func NewCmd(db *gorm.DB) Cmd {
	return Cmd{
		db,
	}
}

var errInvalidUser = errors.New("invalid user")

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

func (cmd *Cmd) ListFolder(ctx context.Context, req *proto.CmdListFolderReq) (*proto.CmdListFolderRes, error) {
	var path = req.Folder
	if len(path) == 0 || path[0] != '/' {
		path = fmt.Sprintf("/%s", path)
	}
	i := len(path) - 1
	for i > 0 && path[i] != '/' {
		i--
	}
	entries, err := os.ReadDir(path[0 : i+1])
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "os.ReadDir: %v", err)
	}
	return &proto.CmdListFolderRes{
		Items: slice.Map(entries, func(item fs.DirEntry) *proto.FolderItem {
			return &proto.FolderItem{
				Name:     item.Name(),
				IsFolder: item.IsDir(),
			}
		}),
	}, nil
}

func (cmd *Cmd) Add(ctx context.Context, req *proto.CmdAddReq) (*proto.CmdAddRes, error) {
	claims, ok := ctx.Value(ctxkey.UseClaims).(*jwtgo.UserClaims)
	if !ok {
		return nil, errInvalidUser
	}
	newCmd := dao.Command{
		Data: datatypes.NewJSONType(dao.CommandData{
			Items: slice.Map(req.Items, func(item *proto.CmdItem) dao.CommandDataItem {
				return dao.CommandDataItem{
					Type:  item.Type,
					Value: item.Value,
				}
			}),
		}),
		Creator: dao.User{
			Model: gorm.Model{
				ID: uint(claims.User.ID),
			},
		},
	}
	err := cmd.db.Create(&newCmd).Error
	if err != nil {
		return nil, errgo.Wrap(err, "AddCmd")
	}
	return &proto.CmdAddRes{
		ID: uint32(newCmd.ID),
	}, nil
}
