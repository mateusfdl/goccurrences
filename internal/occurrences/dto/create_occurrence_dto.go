package dto

import (
	"time"

	"buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
)

type CreateOccurrenceDTO struct {
	SourceID       string                       `json:"sourceId" bson:"sourceId"`
	SourceType     string                       `json:"sourceType" bson:"sourceType"`
	ActorID        string                       `json:"actorId" bson:"actorId"`
	ActorType      string                       `json:"actorType" bson:"actorType"`
	OccurrenceTime time.Time                    `json:"timestamp" bson:"occurrenceTime"`
	OccurrenceCode occurrencesv1.OccurrenceType `json:"occurrenceCode" bson:"occurrenceCode"`
}

type ListUserOccurrenceDTO struct {
	ActorID string `json:"actorId"`
	Skip    uint32 `json:"skip"`
	Limit   uint32 `json:"limit"`
}

func FromCreateOccurrenceProto(oc *occurrencesv1.CreateOccurrenceRequest) *CreateOccurrenceDTO {
	return &CreateOccurrenceDTO{
		SourceType:     oc.SourceType,
		OccurrenceCode: oc.OccurrenceCode,
		SourceID:       oc.SourceId,
		ActorID:        oc.ActorId,
		ActorType:      oc.ActorType,
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
			OccurrenceCode: occurrencesv1.OccurrenceType(occurrence.OccurenceCode),
			OccurrenceTime: timestamppb.New(occurrence.OccurrenceTime),
			SourceId:       occurrence.SourceId,
			SourceType:     occurrence.SourceType,
			ActorId:        occurrence.ActorId,
			ActorType:      occurrence.ActorType,
		}
	}

	return &occurrencesv1.ListUserOccurrencesResponse{Occurrences: res}
}
