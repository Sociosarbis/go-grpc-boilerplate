package web

import (
	"go.uber.org/fx"

	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler"
)

var Module = fx.Module("web", fx.Provide(New), handler.Module)
