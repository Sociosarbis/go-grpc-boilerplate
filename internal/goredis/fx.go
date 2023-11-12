package goredis

import "go.uber.org/fx"

var Module = fx.Module("goredis", fx.Provide(NewRedisClients))
