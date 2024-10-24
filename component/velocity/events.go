package velocity

import "github.com/dwethmar/vork/component"

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "velocity.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "velocity.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "velocity.deleted"
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
	Velocity() *Velocity
}

type CreatedEvent struct {
	velocity *Velocity
}

func NewCreatedEvent(velocity *Velocity) *CreatedEvent {
	return &CreatedEvent{velocity: velocity}
}

func (e *CreatedEvent) Event() string                 { return CreatedEventType }
func (e *CreatedEvent) Velocity() *Velocity           { return e.velocity }
func (e *CreatedEvent) ComponentID() uint             { return e.velocity.ID() }
func (e *CreatedEvent) ComponentType() component.Type { return e.velocity.Type() }
func (e *CreatedEvent) Deleted() bool                 { return false }

type UpdatedEvent struct {
	velocity *Velocity
}

func NewUpdatedEvent(velocity *Velocity) *UpdatedEvent {
	return &UpdatedEvent{velocity: velocity}
}

func (e *UpdatedEvent) Event() string                 { return UpdatedEventType }
func (e *UpdatedEvent) Velocity() *Velocity           { return e.velocity }
func (e *UpdatedEvent) ComponentID() uint             { return e.velocity.ID() }
func (e *UpdatedEvent) ComponentType() component.Type { return e.velocity.Type() }
func (e *UpdatedEvent) Deleted() bool                 { return false }

type DeletedEvent struct {
	velocity *Velocity
}

func NewDeletedEvent(velocity *Velocity) *DeletedEvent {
	return &DeletedEvent{velocity: velocity}
}

func (e *DeletedEvent) Event() string                 { return DeletedEventType }
func (e *DeletedEvent) Velocity() *Velocity           { return e.velocity }
func (e *DeletedEvent) ComponentID() uint             { return e.velocity.ID() }
func (e *DeletedEvent) ComponentType() component.Type { return e.velocity.Type() }
func (e *DeletedEvent) Deleted() bool                 { return true }
