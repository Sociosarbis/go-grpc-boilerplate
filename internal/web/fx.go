package web

import (
	"go.uber.org/fx"

	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/wshandler"
)

var Module = fx.Module("web", fx.Provide(New), handler.Module, wshandler.Module)
