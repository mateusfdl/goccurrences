package events

type LikeCreatedEvent struct {
	ResourceLikedID   string
	ResourceLikedType string
	CreatedAt         string
}
