package services

import (
	"context"

	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/repository"
	"go.uber.org/zap"
)

type OccurrenceService struct {
	occurrenceRepository repository.OccurrenceRepository
	logger               *zap.Logger
}

func New(occurrenceRepository repository.OccurrenceRepository, logger *zap.Logger) *OccurrenceService {
	return &OccurrenceService{
		occurrenceRepository: occurrenceRepository,
		logger:               logger,
	}
}

func (s *OccurrenceService) Create(
	ctx context.Context,
	dto dto.CreateOccurrenceDTO,
) (string, error) {
	id, err := s.occurrenceRepository.Create(ctx, dto)
	if err != nil {
		s.logger.Error("error creating occurrence:", zap.Error(err))
		return "", err
	}

	return id, nil
}
