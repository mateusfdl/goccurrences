package services

import (
	"context"

	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/emitter"
	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
	"github.com/mateusfdl/go-poc/internal/occurrences/events"
	"github.com/mateusfdl/go-poc/internal/occurrences/repository"
	"go.uber.org/zap"
)

type OccurrenceService struct {
	occurrenceRepository repository.OccurrenceRepository
	logger               *zap.Logger
}

func NewOccurrenceService(occurrenceRepository repository.OccurrenceRepository, logger *zap.Logger) *OccurrenceService {
	return &OccurrenceService{
		occurrenceRepository: occurrenceRepository,
		logger:               logger,
	}
}

func (s *OccurrenceService) Create(
	ctx context.Context,
	dto *dto.CreateOccurrenceDTO,
) (string, error) {
	id, err := s.occurrenceRepository.Create(ctx, dto)
	if err != nil {
		s.logger.Error("error creating occurrence:", zap.Error(err))
		return "", err
	}

	s.dispatchOccurrenceEvent(dto, dto.SourceType)
	return id, nil
}

func (s *OccurrenceService) UserOccurrences(
	ctx context.Context,
	dto *dto.ListUserOccurrenceDTO,
) (*[]entity.Occurrence, error) {
	o, err := s.occurrenceRepository.List(
		ctx,
		dto.ActorID,
		dto.Limit,
		dto.Skip,
	)

	if err != nil {
		s.logger.Error("error listing user occurrences:", zap.Error(err))
		return &[]entity.Occurrence{}, err
	}

	return o, nil
}

func (s *OccurrenceService) dispatchOccurrenceEvent(
	dto *dto.CreateOccurrenceDTO,
	occurrenceType entity.OccurrenceType,
) {
	switch occurrenceType {
	case entity.AccountCreated:
		emitter.Emit(events.AccountCreatedEvent{UserID: dto.ActorID})
	default:
		s.logger.Error("unknown occurrence type")
	}
}
