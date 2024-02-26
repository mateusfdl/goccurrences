package events

type Event interface {
	AccountCreatedEvent | PostCreatedEvent | LikeCreatedEvent
}
