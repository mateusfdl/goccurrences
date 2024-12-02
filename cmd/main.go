package main

import (
	"github.com/mateusfdl/go-poc/internal"
	"go.uber.org/fx"
)

var BaseModule = fx.Options(
	internal.CoreModules,
	internal.AdapterModules,
)

func main() {
	fx.New(BaseModule).Run()
}
