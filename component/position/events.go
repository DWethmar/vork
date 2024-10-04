package position

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "position.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "position.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "position.deleted"
)

type CreatedEvent struct {
	Position Position
}

func (e *CreatedEvent) Event() string { return CreatedEventType }

type UpdatedEvent struct {
	Position Position
}

func (e *UpdatedEvent) Event() string { return UpdatedEventType }

type DeletedEvent struct {
	Position Position
}

func (e *DeletedEvent) Event() string { return DeletedEventType }
