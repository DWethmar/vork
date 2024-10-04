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

func (e *CreatedEvent) Event() string                  { return CreatedEventType }
func (e *CreatedEvent) Component() component.Component { return e.Position }
func (e *CreatedEvent) Deleted() bool                  { return false }

type UpdatedEvent struct {
	Position Position
}

func (e *UpdatedEvent) Event() string                  { return UpdatedEventType }
func (e *UpdatedEvent) Component() component.Component { return e.Position }
func (e *UpdatedEvent) Deleted() bool                  { return false }

type DeletedEvent struct {
	Position Position
}

func (e *DeletedEvent) Event() string                  { return DeletedEventType }
func (e *DeletedEvent) Component() component.Component { return e.Position }
func (e *DeletedEvent) Deleted() bool                  { return true }
