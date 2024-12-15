package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jsbridge/jsbridge"
)

func main() {
	if fdStr, ok := os.LookupEnv("NODE_CHANNEL_FD"); ok {
		if fd, err := strconv.ParseInt(fdStr, 10, 32); err == nil {
			rx := make(chan []byte, 10)
			ctx, _ := context.WithTimeout(context.Background(), 5*1e9)
			tx, _ := jsbridge.Run(ctx, uintptr(fd), rx)
			var waitGroup sync.WaitGroup
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				defer close(rx)
				for {
					select {
					case msg, ok := <-rx:
						if !ok {
							return
						}
						fmt.Println(string(msg))
					case <-ctx.Done():
						return
					}
				}
			}()
			go func() {
				for {
					if bytes, err := jsbridge.Stringify("msg", "pong"); err == nil {
						tx <- bytes
					}
					time.Sleep(1e9)
				}
			}()
			waitGroup.Wait()
		}
	}
}
