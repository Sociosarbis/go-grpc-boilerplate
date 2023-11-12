package redis

import (
	"context"
	"errors"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/goredis"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"go.uber.org/fx"
)

func Seed(item string) error {
	if item != "worker_id" {
		return errors.New("not a supported seed item")
	}

	var cfg config.AppConfig
	var clients *goredis.RedisClients

	err := fx.New(config.Module, goredis.Module, fx.Populate(&cfg, &clients)).Err()

	if err != nil {
		return errgo.Wrap(err, "fx.New")
	}

	err = clients.SeedWorkerID(context.Background(), clients.Main, "service_worker")

	if err != nil {
		return errgo.Wrap(err, "seed worker id")
	}
	return nil
}
