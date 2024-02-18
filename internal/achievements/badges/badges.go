package badges

import "github.com/mateusfdl/go-poc/internal/occurrences/entity"

type Badge interface {
	EvaluateRule([]*entity.Occurrence) error
}
