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

func (e *CreatedEvent) Event() string                  { return CreatedEventType }
func (e *CreatedEvent) Component() component.Component { return e.Skeleton }
func (e *CreatedEvent) Deleted() bool                  { return false }

// UpdatedEvent is an event that is sent when a component is updated.
type UpdatedEvent struct {
	Skeleton Skeleton
}

func (e *UpdatedEvent) Event() string                  { return UpdatedEventType }
func (e *UpdatedEvent) Component() component.Component { return e.Skeleton }
func (e *UpdatedEvent) Deleted() bool                  { return false }

// DeletedEvent is an event that is sent when a component is deleted.
type DeletedEvent struct {
	Skeleton Skeleton
}

func (e *DeletedEvent) Event() string                  { return DeletedEventType }
func (e *DeletedEvent) Component() component.Component { return e.Skeleton }
func (e *DeletedEvent) Deleted() bool                  { return true }
