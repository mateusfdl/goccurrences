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

type ListUserOccurrenceDTO struct {
	UserID string `json:"userId"`
	Skip   uint32 `json:"skip"`
	Limit  uint32 `json:"limit"`
}

func FromCreateOccurrenceProto(oc *occurrences.CreateOccurrenceRequest) *CreateOccurrenceDTO {
	return &CreateOccurrenceDTO{
		SourceType: entity.OccurrenceType(oc.OccurrenceCode),
		UserID:     oc.UserId,
		Timestamp:  oc.OccurrenceTime.AsTime(),
	}
}

func FromListUserOccurrenceProto(oc *occurrences.ListUserOccurrencesRequest) *ListUserOccurrenceDTO {
	return &ListUserOccurrenceDTO{
		UserID: oc.UserId,
		Limit:  oc.Limit,
		Skip:   oc.Skip,
	}
}
