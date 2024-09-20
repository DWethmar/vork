package controllable

import "github.com/dwethmar/vork/component"

const (
	// CreatedEvent is the event type for when a component is created.
	CreatedEvent = "controllable.created"
	// UpdatedEvent is the event type for when a component is updated.
	UpdatedEvent = "controllable.updated"
	// DeletedEvent is the event type for when a component is deleted.
	DeletedEvent = "controllable.deleted"
)

// NewCreatedEvent creates a new event with the type CreatedEvent.
func NewCreatedEvent(c Controllable) *component.Event {
	return component.NewEvent(CreatedEvent, c)
}

// NewUpdatedEvent creates a new event with the type UpdatedEvent.
func NewUpdatedEvent(c Controllable) *component.Event {
	return component.NewEvent(UpdatedEvent, c)
}

// NewDeletedEvent creates a new event with the type DeletedEvent.
func NewDeletedEvent(c Controllable) *component.Event {
	return component.NewEvent(DeletedEvent, c)
}
