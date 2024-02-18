package entity

import "time"

const (
  AccountCreated = iota
  PostCreated
  LikeCreated
)

type OccurrenceType int

type Occurrence struct {
	ID         string    `bson:"_id"`
	SourceType int32    `bson:"sourceType"`
	UserID     string    `bson:"userId"`
        OccurrenceTime time.Time `bson:"occurrenceTime"`
	CreatedAt  time.Time `bson:"createdAt"`
}
