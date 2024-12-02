package achievements

import (
	"github.com/mateusfdl/go-poc/internal/achievements/badges"
	"github.com/mateusfdl/go-poc/internal/achievements/listeners"
	"go.uber.org/fx"
)

var Module = fx.Options(
	listeners.Module,
	badges.Module,
)
