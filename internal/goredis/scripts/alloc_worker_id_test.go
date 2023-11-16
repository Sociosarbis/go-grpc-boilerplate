package scripts_test

import (
	"context"
	"testing"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/goredis"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestAllocWorkerID(t *testing.T) {
	var clients *goredis.RedisClients
	fx.New(fx.NopLogger, config.Module, goredis.Module, fx.Populate(&clients))
	id, err := clients.AllocWorkerID(context.Background(), clients.Main, "service_worker")
	require.NoError(t, err)
	require.GreaterOrEqual(t, id, 0)
}
