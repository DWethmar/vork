package skeleton

import "github.com/dwethmar/vork/component"

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "skeleton.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "skeleton.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "skeleton.deleted"
)

var (
	_ component.Event = &CreatedEvent{}
	_ component.Event = &UpdatedEvent{}
	_ component.Event = &DeletedEvent{}
)

// CreatedEvent is an event that is sent when a component is created.
type CreatedEvent struct {
	Skeleton Skeleton
}

func (e *CreatedEvent) Event() string { return CreatedEventType }
func (e *CreatedEvent) Component() any {
	c := e.Skeleton
	return &c
}
func (e *CreatedEvent) ComponentID() uint32                    { return e.Skeleton.ID() }
func (e *CreatedEvent) ComponentType() component.ComponentType { return e.Skeleton.Type() }
func (e *CreatedEvent) Deleted() bool                          { return false }

// UpdatedEvent is an event that is sent when a component is updated.
type UpdatedEvent struct {
	Skeleton Skeleton
}

func (e *UpdatedEvent) Event() string { return UpdatedEventType }
func (e *UpdatedEvent) Component() any {
	c := e.Skeleton
	return &c
}
func (e *UpdatedEvent) ComponentID() uint32                    { return e.Skeleton.ID() }
func (e *UpdatedEvent) ComponentType() component.ComponentType { return e.Skeleton.Type() }
func (e *UpdatedEvent) Deleted() bool                          { return false }

// DeletedEvent is an event that is sent when a component is deleted.
type DeletedEvent struct {
	Skeleton Skeleton
}

func (e *DeletedEvent) Event() string { return DeletedEventType }
func (e *DeletedEvent) Component() any {
	c := e.Skeleton
	return &c
}
func (e *DeletedEvent) ComponentID() uint32                    { return e.Skeleton.ID() }
func (e *DeletedEvent) ComponentType() component.ComponentType { return e.Skeleton.Type() }
func (e *DeletedEvent) Deleted() bool                          { return true }
