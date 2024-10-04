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
)

type CreatedEvent struct {
	Controllable Controllable
}

func (e *CreatedEvent) Event() string                          { return CreatedEventType }
func (e *CreatedEvent) Component() any                         { return e.Controllable }
func (e *CreatedEvent) ComponentID() uint32                    { return e.Controllable.ID() }
func (e *CreatedEvent) ComponentType() component.ComponentType { return e.Controllable.Type() }
func (e *CreatedEvent) Deleted() bool                          { return false }

type UpdatedEvent struct {
	Controllable Controllable
}

func (e *UpdatedEvent) Event() string                          { return UpdatedEventType }
func (e *UpdatedEvent) Component() any                         { return e.Controllable }
func (e *UpdatedEvent) ComponentID() uint32                    { return e.Controllable.ID() }
func (e *UpdatedEvent) ComponentType() component.ComponentType { return e.Controllable.Type() }
func (e *UpdatedEvent) Deleted() bool                          { return false }

type DeletedEvent struct {
	Controllable Controllable
}

func (e *DeletedEvent) Event() string                          { return DeletedEventType }
func (e *DeletedEvent) Component() any                         { return e.Controllable }
func (e *DeletedEvent) ComponentID() uint32                    { return e.Controllable.ID() }
func (e *DeletedEvent) ComponentType() component.ComponentType { return e.Controllable.Type() }
func (e *DeletedEvent) Deleted() bool                          { return true }
