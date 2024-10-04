package component

import "github.com/dwethmar/vork/event"

// Event is a change in a component.
type Event interface {
	event.Event
	Component() any
	ComponentID() uint32
	ComponentType() ComponentType
	Deleted() bool
}
