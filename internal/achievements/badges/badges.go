package badges

import (
	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
	"go.uber.org/fx"
)

type Badge interface {
	EvaluateRule(*[]entity.Occurrence) error
}

var Module = fx.Module("badges", fx.Provide(NewAccountCreatedBadge))
