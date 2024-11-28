package entity

import "time"

const (
	AccountCreated = iota
	PostCreated
	LikeCreated
)

type OccurrenceType int

type Occurrence struct {
	ID             string         `bson:"_id"`
	SourceId       string         `bson:"sourceId"`
	SourceType     OccurrenceType `bson:"sourceType"`
	ActorId        string         `bson:"actorId"`
	ActorType      string         `bson:"actorType"`
	OccurrenceTime time.Time      `bson:"occurrenceTime"`
	CreatedAt      time.Time      `bson:"createdAt"`
}
