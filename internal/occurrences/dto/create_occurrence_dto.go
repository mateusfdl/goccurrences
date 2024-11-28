package dto

import (
	"time"

	"buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
)

type CreateOccurrenceDTO struct {
	SourceID       string                `json:"sourceId"`
	SourceType     entity.OccurrenceType `json:"sourceType"`
	ActorID        string                `json:"actorId"`
	ActorType      string                `json:"actorType"`
	OccurrenceTime time.Time             `json:"timestamp"`
}

type ListUserOccurrenceDTO struct {
	ActorID string `json:"actorId"`
	Skip    uint32 `json:"skip"`
	Limit   uint32 `json:"limit"`
}

func FromCreateOccurrenceProto(oc *occurrencesv1.CreateOccurrenceRequest) *CreateOccurrenceDTO {
	return &CreateOccurrenceDTO{
		SourceType:     entity.OccurrenceType(oc.OccurrenceCode),
		SourceID:       oc.SourceId,
		ActorID:        oc.ActorId,
		ActorType:      "User",
		OccurrenceTime: oc.OccurrenceTime.AsTime(),
	}
}

func FromListUserOccurrenceProto(oc *occurrencesv1.ListUserOccurrencesRequest) *ListUserOccurrenceDTO {
	return &ListUserOccurrenceDTO{
		ActorID: oc.ActorId,
		Limit:   oc.Limit,
		Skip:    oc.Skip,
	}
}

func (l *ListUserOccurrenceDTO) ToProto(oc *[]entity.Occurrence) *occurrencesv1.ListUserOccurrencesResponse {
	var res = make([]*occurrencesv1.Occurrence, len(*oc))

	for i, occurrence := range *oc {
		res[i] = &occurrencesv1.Occurrence{
			OccurrenceId:   occurrence.ID,
			OccurrenceCode: occurrencesv1.OccurrenceType(occurrence.SourceType),
			OccurrenceTime: timestamppb.New(occurrence.OccurrenceTime),
		}
	}

	return &occurrencesv1.ListUserOccurrencesResponse{
		Occurrences: res,
	}
}
