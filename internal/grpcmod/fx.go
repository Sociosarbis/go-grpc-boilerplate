package grpcmod

import (
	"go.uber.org/fx"

	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/handler"
)

var Module = fx.Module("grpc", fx.Provide(New, NewClient), handler.Module)
