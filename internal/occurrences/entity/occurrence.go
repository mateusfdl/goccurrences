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
	SourceType string    `bson:"sourceType"`
	UserID     string    `bson:"userId"`
	CreatedAt  time.Time `bson:"createdAt"`
}
