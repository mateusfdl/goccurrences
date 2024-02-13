package grpc

import (
	"context"

	"buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences"
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
	b *occurrences.NewOccurrence,
) (*occurrences.OccurrenceResponse, error) {
	dto := dto.FromProto(b)

	id, err := h.service.Create(ctx, dto)
	if err != nil {
          h.l.Fatal(err)
	}

	return &occurrences.OccurrenceResponse{OccurrenceId: id}, nil
}
