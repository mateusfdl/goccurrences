package events

import "github.com/mateusfdl/go-poc/internal/occurrences/entity"

type AccountCreatedEvent struct {
	Code int
}

func NewAccountCreatedEvent() *AccountCreatedEvent {
	return &AccountCreatedEvent{
		Code: entity.AccountCreated,
	}
}
