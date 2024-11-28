package grpc

import (
	"context"

	"buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	"github.com/mateusfdl/go-poc/internal/occurrences/dto"
	"github.com/mateusfdl/go-poc/internal/occurrences/services"
	"go.uber.org/zap"
)

type OccurrenceHandler struct {
	service *services.OccurrenceService
	l       *zap.SugaredLogger
}

func NewOccurrenceHandler(s *services.OccurrenceService, l *zap.SugaredLogger) *OccurrenceHandler {
	return &OccurrenceHandler{service: s, l: l}
}

func (h OccurrenceHandler) CreateOccurrence(ctx context.Context, b *occurrencesv1.CreateOccurrenceRequest,
) (*occurrencesv1.CreateOccurrenceResponse, error) {
	dto := dto.FromCreateOccurrenceProto(b)

	id, err := h.service.Create(ctx, dto)
	if err != nil {
		h.l.Fatal(err)
	}

	return &occurrencesv1.CreateOccurrenceResponse{OccurrenceId: id}, nil
}

func (h OccurrenceHandler) ListUserOccurrences(ctx context.Context, b *occurrencesv1.ListUserOccurrencesRequest,
) (*occurrencesv1.ListUserOccurrencesResponse, error) {
	dto := dto.FromListUserOccurrenceProto(b)

	occurrences, err := h.service.UserOccurrences(ctx, dto)
	if err != nil {
		h.l.Fatal(err)
	}

	p := dto.ToProto(occurrences)
	return p, nil
}
