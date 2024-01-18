package occurrences_test

import (
	"github.com/mateusfdl/go-poc/internal/occurrences"
	"go.uber.org/fx"
)

var TestApp = fx.New(
	occurrences.Module,
)
