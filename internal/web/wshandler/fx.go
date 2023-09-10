package wshandler

import (
	"go.uber.org/fx"
)

var Module = fx.Module("wsHandler", fx.Provide(NewUser, NewCmd))
