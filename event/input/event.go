package input

import "github.com/dwethmar/vork/event"

// LeftMouseClickedEventType is the event type for when a component is clicked.
const LeftMouseClickedEventType = "input.clicked"

var _ event.Event = &LeftMouseClickedEvent{}

// LeftMouseClickedEvent is an event that is sent when a component is clicked.
type LeftMouseClickedEvent struct {
	X, Y int
}

func NewLeftMouseClickedEvent(x, y int) *LeftMouseClickedEvent {
	return &LeftMouseClickedEvent{
		X: x,
		Y: y,
	}
}

func (e *LeftMouseClickedEvent) Event() string { return LeftMouseClickedEventType }
