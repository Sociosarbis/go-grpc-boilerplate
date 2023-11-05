package goredis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
)

type scripts struct {
	genID *redis.Script
}

type RedisClients struct {
	Main *redis.Client
	scripts
}

// go:embed scripts/gen_id.lua
var genIDFile string

var errorGenID = errors.New("gen id error")

func NewRedisClients(c config.AppConfig) RedisClients {

	return RedisClients{
		Main: redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort),
				Password: c.RedisPassword,
				DB:       0,
			},
		),
		scripts: scripts{
			genID: redis.NewScript(genIDFile),
		},
	}
}

func (c *RedisClients) GenID(ctx context.Context, client *redis.Client, key string, dataCenterID int, workerID int) (int64, error) {
	res, err := c.scripts.genID.Run(ctx, client, []string{key}, dataCenterID, workerID).Int64()
	if err != nil {
		logger.Err(err, "redis.GenID")
		return 0, errorGenID
	}
	return res, nil
}
