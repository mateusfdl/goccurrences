package grpc

import (
	"go.uber.org/fx"
)

var Module = fx.Module("occurrences-grpc", fx.Provide())

var Invokables = fx.Invoke()
