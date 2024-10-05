package component

import "github.com/dwethmar/vork/event"

// Event is a change in a component.
type Event interface {
	event.Event
	ComponentID() uint32
	ComponentType() ComponentType
	Deleted() bool
}
