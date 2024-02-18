package grpc

import (
	"context"

	occurrences "buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/services"
	"go.uber.org/zap"
)

type OccurrenceHandler struct {
	service *services.OccurrenceService
	l       *zap.SugaredLogger
}

func NewOccurrenceHandler(
	s *services.OccurrenceService,
	l *zap.SugaredLogger,
) *OccurrenceHandler {
	return &OccurrenceHandler{
		service: s,
		l:       l,
	}
}

func (h OccurrenceHandler) CreateOccurrence(
	ctx context.Context,
	b *occurrences.CreateOccurrenceRequest,
) (*occurrences.CreateOccurrenceResponse, error) {
	dto := dto.FromCreateOccurrenceProto(b)

	id, err := h.service.Create(ctx, dto)
	if err != nil {
		h.l.Fatal(err)
	}

	return &occurrences.CreateOccurrenceResponse{OccurrenceId: id}, nil
}

func (h OccurrenceHandler) ListUserOccurrences(ctx context.Context, b *occurrences.ListUserOccurrencesRequest) (*occurrences.ListUserOccurrencesResponse, error) {
	dto.FromListUserOccurrenceProto(b)

	return &occurrences.ListUserOccurrencesResponse{}, nil
}
