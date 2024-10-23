package hitbox

import "github.com/dwethmar/vork/component"

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "hitbox.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "hitbox.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "hitbox.deleted"
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
	Hitbox() *Hitbox
}

type CreatedEvent struct {
	hitbox Hitbox
}

func NewCreatedEvent(hitbox Hitbox) *CreatedEvent {
	return &CreatedEvent{hitbox: hitbox}
}

func (e *CreatedEvent) Event() string                 { return CreatedEventType }
func (e *CreatedEvent) Hitbox() *Hitbox               { return &e.hitbox }
func (e *CreatedEvent) ComponentID() uint             { return e.hitbox.ID() }
func (e *CreatedEvent) ComponentType() component.Type { return e.hitbox.Type() }
func (e *CreatedEvent) Deleted() bool                 { return false }

type UpdatedEvent struct {
	hitbox Hitbox
}

func NewUpdatedEvent(hitbox Hitbox) *UpdatedEvent {
	return &UpdatedEvent{hitbox: hitbox}
}

func (e *UpdatedEvent) Event() string                 { return UpdatedEventType }
func (e *UpdatedEvent) Hitbox() *Hitbox               { return &e.hitbox }
func (e *UpdatedEvent) ComponentID() uint             { return e.hitbox.ID() }
func (e *UpdatedEvent) ComponentType() component.Type { return e.hitbox.Type() }
func (e *UpdatedEvent) Deleted() bool                 { return false }

type DeletedEvent struct {
	hitbox Hitbox
}

func NewDeletedEvent(hitbox Hitbox) *DeletedEvent {
	return &DeletedEvent{hitbox: hitbox}
}

func (e *DeletedEvent) Event() string                 { return DeletedEventType }
func (e *DeletedEvent) Hitbox() *Hitbox               { return &e.hitbox }
func (e *DeletedEvent) ComponentID() uint             { return e.hitbox.ID() }
func (e *DeletedEvent) ComponentType() component.Type { return e.hitbox.Type() }
func (e *DeletedEvent) Deleted() bool                 { return true }
