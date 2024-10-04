package controllable

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "controllable.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "controllable.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "controllable.deleted"
)

type CreatedEvent struct {
	Controllable Controllable
}

func (e *CreatedEvent) Event() string { return CreatedEventType }

type UpdatedEvent struct {
	Controllable Controllable
}

func (e *UpdatedEvent) Event() string { return UpdatedEventType }

type DeletedEvent struct {
	Controllable Controllable
}

func (e *DeletedEvent) Event() string { return DeletedEventType }
