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

	_ Event = &CreatedEvent{}
	_ Event = &UpdatedEvent{}
	_ Event = &DeletedEvent{}
)

type Event interface {
	component.Event
	Skeleton() *Skeleton
}

// CreatedEvent is an event that is sent when a component is created.
type CreatedEvent struct {
	skeleton Skeleton
}

func NewCreatedEvent(skeleton Skeleton) *CreatedEvent {
	return &CreatedEvent{skeleton: skeleton}
}

func (e *CreatedEvent) Event() string                 { return CreatedEventType }
func (e *CreatedEvent) Skeleton() *Skeleton           { return &e.skeleton }
func (e *CreatedEvent) ComponentID() uint             { return e.skeleton.ID() }
func (e *CreatedEvent) ComponentType() component.Type { return e.skeleton.Type() }
func (e *CreatedEvent) Deleted() bool                 { return false }

// UpdatedEvent is an event that is sent when a component is updated.
type UpdatedEvent struct {
	skeleton Skeleton
}

func NewUpdatedEvent(skeleton Skeleton) *UpdatedEvent {
	return &UpdatedEvent{skeleton: skeleton}
}

func (e *UpdatedEvent) Event() string                 { return UpdatedEventType }
func (e *UpdatedEvent) Skeleton() *Skeleton           { return &e.skeleton }
func (e *UpdatedEvent) ComponentID() uint             { return e.skeleton.ID() }
func (e *UpdatedEvent) ComponentType() component.Type { return e.skeleton.Type() }
func (e *UpdatedEvent) Deleted() bool                 { return false }

// DeletedEvent is an event that is sent when a component is deleted.
type DeletedEvent struct {
	skeleton Skeleton
}

func NewDeletedEvent(skeleton Skeleton) *DeletedEvent {
	return &DeletedEvent{skeleton: skeleton}
}

func (e *DeletedEvent) Event() string                 { return DeletedEventType }
func (e *DeletedEvent) Skeleton() *Skeleton           { return &e.skeleton }
func (e *DeletedEvent) ComponentID() uint             { return e.skeleton.ID() }
func (e *DeletedEvent) ComponentType() component.Type { return e.skeleton.Type() }
func (e *DeletedEvent) Deleted() bool                 { return true }
