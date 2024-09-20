package component

import "github.com/dwethmar/vork/event"

var _ event.Event = &Event{}

// Event is a struct that represents an event that is sent to a system.
type Event struct {
	Type      string
	Component Component
}

func (e *Event) Event() string {
	return e.Type
}

// NewEvent creates a new event with the given type and component.
func NewEvent(t string, c Component) *Event {
	return &Event{
		Type:      t,
		Component: c,
	}
}
