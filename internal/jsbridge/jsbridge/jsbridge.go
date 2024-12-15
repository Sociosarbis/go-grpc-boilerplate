package jsbridge

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
)

func Stringify(typ string, msg string) ([]byte, error) {
	out, err := json.Marshal(map[string]any{
		"type": typ,
		"msg":  msg,
	})
	return append(out, '\n'), err
}

func Run(ctx context.Context, fd uintptr, tx chan<- []byte) (chan<- []byte, context.CancelFunc) {
	file := os.NewFile(fd, "pipe")
	scanner := bufio.NewScanner(file)
	rx := make(chan []byte, 10)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		defer file.Close()
		defer close(rx)
		for {
			select {
			case msg, ok := <-rx:
				if !ok {
					return
				}
				if _, err := file.Write(msg); err != nil {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		for scanner.Scan() {
			tx <- scanner.Bytes()
		}
	}()
	return rx, cancel
}
