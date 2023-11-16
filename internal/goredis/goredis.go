package goredis

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
)

type scripts struct {
	genID          *redis.Script
	allocWorkerID  *redis.Script
	expireWorkerID *redis.Script
	seedWorkerID   *redis.Script
}

type RedisClients struct {
	Main *redis.Client
	scripts
}

//go:embed scripts/gen_id.lua
var genIDFile string

//go:embed scripts/alloc_worker_id.lua
var allocWorkerIDFile string

//go:embed scripts/expire_worker_id.lua
var expireWorkerIDFile string

//go:embed scripts/seed_worker_id.lua
var seedWorkerIDFile string

var errorGenID = errors.New("gen id error")
var errorAllocWorkerID = errors.New("alloc worker id error")
var errorExpireWorkerID = errors.New("expire worker id error")
var errorSeedWorkerID = errors.New("seed worker id error")

func NewRedisClients(c config.AppConfig) *RedisClients {

	return &RedisClients{
		Main: redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort),
				Password: c.RedisPassword,
				DB:       0,
			},
		),
		scripts: scripts{
			genID:          redis.NewScript(genIDFile),
			allocWorkerID:  redis.NewScript(allocWorkerIDFile),
			expireWorkerID: redis.NewScript(expireWorkerIDFile),
			seedWorkerID:   redis.NewScript(seedWorkerIDFile),
		},
	}
}

func (c *RedisClients) GenID(ctx context.Context, client *redis.Client, key string, dataCenterID int, workerID int) (uint64, error) {
	res, err := c.scripts.genID.Run(ctx, client, []string{key}, dataCenterID, workerID).Uint64Slice()
	if err != nil {
		logger.Err(err, "redis.GenID")
		return 0, errorGenID
	}
	return res[0]<<22 | res[1], nil
}

func (c *RedisClients) AllocWorkerID(ctx context.Context, client *redis.Client, key string) (int, error) {
	res, err := c.scripts.allocWorkerID.Run(ctx, client, []string{key}).Int()
	if err != nil {
		logger.Err(err, "redis.AllocWorkerID")
		return 0, errorAllocWorkerID
	}
	return res, nil
}

func (c *RedisClients) ExpireWorkerID(ctx context.Context, client *redis.Client, key string) error {
	err := c.scripts.expireWorkerID.Run(ctx, client, []string{key}, 30*time.Minute.Milliseconds()).Err()
	if err != nil {
		logger.Err(err, "redis.ExpireWorkerID")
		return errorExpireWorkerID
	}
	return nil
}

func (c *RedisClients) SeedWorkerID(ctx context.Context, client *redis.Client, key string) error {
	err := c.scripts.seedWorkerID.Run(ctx, client, []string{key}, 32).Err()
	if err != nil {
		logger.Err(err, "redis.SeedWorkerID")
		return errorSeedWorkerID
	}
	return nil
}
