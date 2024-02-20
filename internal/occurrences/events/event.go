package events

import "go.uber.org/fx"

type Event interface {
	AccountCreatedEvent | PostCreatedEvent | LikeCreatedEvent
}

var Module = fx.Module("occurrence-events",
	fx.Provide(NewAccountCreatedEvent),
	fx.Provide(NewLikeCreatedEvent),
	fx.Provide(NewPostCreatedEvent),
)
