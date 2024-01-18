package dto

import (
	"time"

	"buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences"
)

type CreateOccurrenceDTO struct {
	SourceType string    `json:"sourceType"`
	UserID     string    `json:"userId"`
	Timestamp  time.Time `json:"timestamp"`
}

func fromProto(oc *occurrences.NewOccurrence) CreateOccurrenceDTO {
	return CreateOccurrenceDTO{
		SourceType: oc.OccurrenceCode.String(),
		UserID:     oc.UserId,
		Timestamp:  oc.OccurrenceTime.AsTime(),
	}
}
