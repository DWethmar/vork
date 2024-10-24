package chunk

import "github.com/dwethmar/vork/component"

const (
	// CreatedEventType is the event type for when a component is created.
	CreatedEventType = "chunk.created"
	// UpdatedEventType is the event type for when a component is updated.
	UpdatedEventType = "chunk.updated"
	// DeletedEventType is the event type for when a component is deleted.
	DeletedEventType = "chunk.deleted"
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
	Chunk() *Chunk
}

type CreatedEvent struct {
	chunk *Chunk
}

func NewCreatedEvent(chunk *Chunk) *CreatedEvent {
	return &CreatedEvent{chunk: chunk}
}

func (e *CreatedEvent) Event() string                 { return CreatedEventType }
func (e *CreatedEvent) Chunk() *Chunk                 { return e.chunk }
func (e *CreatedEvent) ComponentID() uint             { return e.chunk.ID() }
func (e *CreatedEvent) ComponentType() component.Type { return e.chunk.Type() }
func (e *CreatedEvent) Deleted() bool                 { return false }

type UpdatedEvent struct {
	chunk *Chunk
}

func NewUpdatedEvent(chunk *Chunk) *UpdatedEvent {
	return &UpdatedEvent{chunk: chunk}
}

func (e *UpdatedEvent) Event() string                 { return UpdatedEventType }
func (e *UpdatedEvent) Chunk() *Chunk                 { return e.chunk }
func (e *UpdatedEvent) ComponentID() uint             { return e.chunk.ID() }
func (e *UpdatedEvent) ComponentType() component.Type { return e.chunk.Type() }
func (e *UpdatedEvent) Deleted() bool                 { return false }

type DeletedEvent struct {
	chunk *Chunk
}

func NewDeletedEvent(chunk *Chunk) *DeletedEvent {
	return &DeletedEvent{chunk: chunk}
}

func (e *DeletedEvent) Event() string                 { return DeletedEventType }
func (e *DeletedEvent) Chunk() *Chunk                 { return e.chunk }
func (e *DeletedEvent) ComponentID() uint             { return e.chunk.ID() }
func (e *DeletedEvent) ComponentType() component.Type { return e.chunk.Type() }
func (e *DeletedEvent) Deleted() bool                 { return true }
