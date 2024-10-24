package position

import "github.com/dwethmar/vork/component"

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "position.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "position.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "position.deleted"
)

var (
	_ component.Event = &CreatedEvent{}
	_ component.Event = &UpdatedEvent{}
	_ component.Event = &DeletedEvent{}

	_ Event = &CreatedEvent{}
	_ Event = &UpdatedEvent{}
	_ Event = &DeletedEvent{}
)

// Event is a change in a component.
type Event interface {
	component.Event
	Position() *Position
}

type CreatedEvent struct {
	position *Position
}

func NewCreatedEvent(position *Position) *CreatedEvent {
	return &CreatedEvent{position: position}
}

func (e *CreatedEvent) Event() string                 { return CreatedEventType }
func (e *CreatedEvent) Position() *Position           { return e.position }
func (e *CreatedEvent) ComponentID() uint             { return e.position.ID() }
func (e *CreatedEvent) ComponentType() component.Type { return e.position.Type() }
func (e *CreatedEvent) Deleted() bool                 { return false }

type UpdatedEvent struct {
	position *Position
}

func NewUpdatedEvent(position *Position) *UpdatedEvent {
	return &UpdatedEvent{position: position}
}

func (e *UpdatedEvent) Event() string                 { return UpdatedEventType }
func (e *UpdatedEvent) Position() *Position           { return e.position }
func (e *UpdatedEvent) ComponentID() uint             { return e.position.ID() }
func (e *UpdatedEvent) ComponentType() component.Type { return e.position.Type() }
func (e *UpdatedEvent) Deleted() bool                 { return false }

type DeletedEvent struct {
	position *Position
}

func NewDeletedEvent(position *Position) *DeletedEvent {
	return &DeletedEvent{position: position}
}

func (e *DeletedEvent) Event() string                 { return DeletedEventType }
func (e *DeletedEvent) Position() *Position           { return e.position }
func (e *DeletedEvent) ComponentID() uint             { return e.position.ID() }
func (e *DeletedEvent) ComponentType() component.Type { return e.position.Type() }
func (e *DeletedEvent) Deleted() bool                 { return true }
