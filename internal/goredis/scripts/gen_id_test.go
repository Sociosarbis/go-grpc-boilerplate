package scripts_test

import (
	"context"
	"testing"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/goredis"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestGenID(t *testing.T) {
	var clients *goredis.RedisClients
	fx.New(fx.NopLogger, config.Module, goredis.Module, fx.Populate(&clients))
	id, err := clients.GenID(context.Background(), clients.Main, "service_worker", 0, 0)
	require.NoError(t, err)
	require.GreaterOrEqual(t, id, uint64(0))
}
