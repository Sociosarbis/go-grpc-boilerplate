package grpcmod

import (
	"go.uber.org/fx"

	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/grpcmod/interceptor"
)

var Module = fx.Module("grpc", fx.Provide(New, NewClient, interceptor.NewAuth), handler.Module)
