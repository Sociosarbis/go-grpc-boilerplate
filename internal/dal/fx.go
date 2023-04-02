package dal

import "go.uber.org/fx"

var Module = fx.Module("dal", fx.Provide(NewDB))
