package component

import "github.com/dwethmar/vork/event"

// Event is a change in a component.
type Event interface {
	event.Event
	Component() Component
	Deleted() bool
}
