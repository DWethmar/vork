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
)

type CreatedEvent struct {
	Position Position
}

func (e *CreatedEvent) Event() string { return CreatedEventType }
func (e *CreatedEvent) Component() any {
	c := e.Position
	return &c
}
func (e *CreatedEvent) ComponentID() uint32                    { return e.Position.ID() }
func (e *CreatedEvent) ComponentType() component.ComponentType { return e.Position.Type() }
func (e *CreatedEvent) Deleted() bool                          { return false }

type UpdatedEvent struct {
	Position Position
}

func (e *UpdatedEvent) Event() string { return UpdatedEventType }
func (e *UpdatedEvent) Component() any {
	c := e.Position
	return &c
}
func (e *UpdatedEvent) ComponentID() uint32                    { return e.Position.ID() }
func (e *UpdatedEvent) ComponentType() component.ComponentType { return e.Position.Type() }
func (e *UpdatedEvent) Deleted() bool                          { return false }

type DeletedEvent struct {
	Position Position
}

func (e *DeletedEvent) Event() string { return DeletedEventType }
func (e *DeletedEvent) Component() any {
	c := e.Position
	return &c
}
func (e *DeletedEvent) ComponentID() uint32                    { return e.Position.ID() }
func (e *DeletedEvent) ComponentType() component.ComponentType { return e.Position.Type() }
func (e *DeletedEvent) Deleted() bool                          { return true }
