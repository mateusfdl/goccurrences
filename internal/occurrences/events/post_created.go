package events

import "github.com/mateusfdl/go-poc/internal/occurrences/entity"

type PostCreatedEvent struct {
	Code int
}

func NewPostCreatedEvent() *PostCreatedEvent {
	return &PostCreatedEvent{
		Code: entity.PostCreated,
	}
}
