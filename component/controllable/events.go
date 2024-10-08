package controllable

import "github.com/dwethmar/vork/component"

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "controllable.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "controllable.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "controllable.deleted"
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
	Controllable() *Controllable
}

type CreatedEvent struct {
	controllable Controllable
}

func NewCreatedEvent(controllable Controllable) *CreatedEvent {
	return &CreatedEvent{controllable: controllable}
}

func (e *CreatedEvent) Event() string                          { return CreatedEventType }
func (e *CreatedEvent) Controllable() *Controllable            { return &e.controllable }
func (e *CreatedEvent) ComponentID() uint                      { return e.controllable.ID() }
func (e *CreatedEvent) ComponentType() component.ComponentType { return e.controllable.Type() }
func (e *CreatedEvent) Deleted() bool                          { return false }

type UpdatedEvent struct {
	controllable Controllable
}

func NewUpdatedEvent(controllable Controllable) *UpdatedEvent {
	return &UpdatedEvent{controllable: controllable}
}

func (e *UpdatedEvent) Event() string                          { return UpdatedEventType }
func (e *UpdatedEvent) Controllable() *Controllable            { return &e.controllable }
func (e *UpdatedEvent) ComponentID() uint                      { return e.controllable.ID() }
func (e *UpdatedEvent) ComponentType() component.ComponentType { return e.controllable.Type() }
func (e *UpdatedEvent) Deleted() bool                          { return false }

type DeletedEvent struct {
	controllable Controllable
}

func NewDeletedEvent(controllable Controllable) *DeletedEvent {
	return &DeletedEvent{controllable: controllable}
}

func (e *DeletedEvent) Event() string                          { return DeletedEventType }
func (e *DeletedEvent) Controllable() *Controllable            { return &e.controllable }
func (e *DeletedEvent) ComponentID() uint                      { return e.controllable.ID() }
func (e *DeletedEvent) ComponentType() component.ComponentType { return e.controllable.Type() }
func (e *DeletedEvent) Deleted() bool                          { return true }
