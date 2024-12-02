package events

import (
	"time"
)

type AccountCreatedEvent struct {
	UserID    string
	CreatedAt time.Time
}
