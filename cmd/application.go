package main

import (
	"github.com/mateusfdl/go-poc/internal"
	"go.uber.org/fx"
)

type Application struct {
	BaseModule fx.Option
}

func NewApplication() *Application {
	return &Application{
		BaseModule: fx.Options(
			internal.CoreModules,
			internal.AdapterModules,
		),
	}
}

func (a *Application) Start() {
	a.Start()
}
