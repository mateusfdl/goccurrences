package dto

import (
	"time"

	occurrences "buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
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

func FromCreateOccurrenceProto(oc *occurrences.CreateOccurrenceRequest) *CreateOccurrenceDTO {
	return &CreateOccurrenceDTO{
		SourceType:     entity.OccurrenceType(oc.OccurrenceCode),
		SourceID:       oc.UserId,
		ActorID:        oc.UserId,
		ActorType:      "User",
		OccurrenceTime: oc.OccurrenceTime.AsTime(),
	}
}

func FromListUserOccurrenceProto(oc *occurrences.ListUserOccurrencesRequest) *ListUserOccurrenceDTO {
	return &ListUserOccurrenceDTO{
		ActorID: oc.UserId,
		Limit:   oc.Limit,
		Skip:    oc.Skip,
	}
}

func (l *ListUserOccurrenceDTO) ToProto(oc *[]entity.Occurrence) *occurrences.ListUserOccurrencesResponse {
	var res = make([]*occurrences.Occurrence, len(*oc))

	for i, occurrence := range *oc {
		res[i] = &occurrences.Occurrence{
			OccurrenceId:   occurrence.ID,
			OccurrenceCode: occurrences.OccurrenceType(occurrence.SourceType),
			OccurrenceTime: timestamppb.New(occurrence.OccurrenceTime),
		}
	}

	return &occurrences.ListUserOccurrencesResponse{
		Occurrences: res,
	}
}
