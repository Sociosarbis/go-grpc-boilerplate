package handler

import "go.uber.org/fx"

var Module = fx.Module("grpcHandler", fx.Provide(NewUser, NewCmd, New))
