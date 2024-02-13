package repository

import (
	"context"

	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
)

type OccurrenceRepository interface {
	Create(ctx context.Context, occurrence dto.CreateOccurrenceDTO) (string, error)
	List(ctx context.Context, ID string) ([]entity.Occurrence, error)
}