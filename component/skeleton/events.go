package skeleton

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "skeleton.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "skeleton.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "skeleton.deleted"
)

type CreatedEvent struct {
	Skeleton Skeleton
}

func (e *CreatedEvent) Event() string { return CreatedEventType }

type UpdatedEvent struct {
	Skeleton Skeleton
}

func (e *UpdatedEvent) Event() string { return UpdatedEventType }

type DeletedEvent struct {
	Skeleton Skeleton
}

func (e *DeletedEvent) Event() string { return DeletedEventType }
