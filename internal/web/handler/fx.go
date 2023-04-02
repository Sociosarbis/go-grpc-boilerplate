package handler

import (
	"go.uber.org/fx"

	"github.com/sociosarbis/grpc/boilerplate/internal/web/handler/common"
)

var Module = fx.Module("webHandler", fx.Provide(common.New, NewUser))
