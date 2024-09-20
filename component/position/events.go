package position

import "github.com/dwethmar/vork/component"

const (
	// CreatedEvent is the event type for when a component is created.
	CreatedEvent = "position.created"
	// UpdatedEvent is the event type for when a component is updated.
	UpdatedEvent = "position.updated"
	// DeletedEvent is the event type for when a component is deleted.
	DeletedEvent = "position.deleted"
)

// NewCreatedEvent creates a new event with the type CreatedEvent.
func NewCreatedEvent(c Position) *component.Event {
	return component.NewEvent(CreatedEvent, c)
}

// NewUpdatedEvent creates a new event with the type UpdatedEvent.
func NewUpdatedEvent(c Position) *component.Event {
	return component.NewEvent(UpdatedEvent, c)
}

// NewDeletedEvent creates a new event with the type DeletedEvent.
func NewDeletedEvent(c Position) *component.Event {
	return component.NewEvent(DeletedEvent, c)
}
