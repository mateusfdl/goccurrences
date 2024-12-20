package entity

import "time"

const (
	AccountCreated = iota
	PostCreated
	LikeCreated
)

type OccurrenceType int

type Occurrence struct {
	ID             string         `bson:"id"`
	SourceId       string         `bson:"sourceId"`
	SourceType     string         `bson:"sourceType"`
	OccurenceCode  OccurrenceType `bson:"occurrenceCode"`
	ActorId        string         `bson:"actorId"`
	ActorType      string         `bson:"actorType"`
	OccurrenceTime time.Time      `bson:"occurrenceTime"`
	CreatedAt      time.Time      `bson:"createdAt"`
}
