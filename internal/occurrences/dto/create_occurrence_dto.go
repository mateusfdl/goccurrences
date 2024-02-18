package dto

import (
	"time"

	occurrences "buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	occurrencesv1 "buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mateusfdl/go-poc/internal/occurrences/entity"
)

type CreateOccurrenceDTO struct {
	SourceType     entity.OccurrenceType `json:"sourceType"`
	UserID         string                `json:"userId"`
	OccurrenceTime time.Time             `json:"timestamp"`
}

type ListUserOccurrenceDTO struct {
	UserID string `json:"userId"`
	Skip   uint32 `json:"skip"`
	Limit  uint32 `json:"limit"`
}

func FromCreateOccurrenceProto(oc *occurrences.CreateOccurrenceRequest) *CreateOccurrenceDTO {
	return &CreateOccurrenceDTO{
		SourceType:     entity.OccurrenceType(oc.OccurrenceCode),
		UserID:         oc.UserId,
		OccurrenceTime: oc.OccurrenceTime.AsTime(),
	}
}

func FromListUserOccurrenceProto(oc *occurrences.ListUserOccurrencesRequest) *ListUserOccurrenceDTO {
	return &ListUserOccurrenceDTO{
		UserID: oc.UserId,
		Limit:  oc.Limit,
		Skip:   oc.Skip,
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
