package entity

import "time"

const (
	DrillPerformed = iota
	PostInCommunityFeed
)

type Occurrence struct {
	ID         string    `bson:"_id"`
	SourceType string    `bson:"sourceType"`
	UserID     string    `bson:"userId"`
	CreatedAt  time.Time `bson:"createdAt"`
}
