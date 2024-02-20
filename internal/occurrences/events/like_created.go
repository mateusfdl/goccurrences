package events

import "github.com/mateusfdl/go-poc/internal/occurrences/entity"

type LikeCreatedEvent struct {
	Code int
}

func NewLikeCreatedEvent() *LikeCreatedEvent {
	return &LikeCreatedEvent{
		Code: entity.LikeCreated,
	}
}
