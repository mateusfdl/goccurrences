package dto

import (
	occurrences "buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	"time"

	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
)

type CreateOccurrenceDTO struct {
	SourceType entity.OccurrenceType `json:"sourceType"`
	UserID     string                `json:"userId"`
	Timestamp  time.Time             `json:"timestamp"`
}

func FromProto(oc *occurrences.CreateOccurrenceRequest) CreateOccurrenceDTO {
	return CreateOccurrenceDTO{
		SourceType: entity.OccurrenceType(oc.OccurrenceCode),
		UserID:     oc.UserId,
		Timestamp:  oc.OccurrenceTime.AsTime(),
	}
}
