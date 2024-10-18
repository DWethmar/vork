package mouse

import "github.com/dwethmar/vork/event"

// LeftMouseClickedEventType is the event type for when a component is clicked.
const LeftMouseClickedEventType = "input.clicked"

var _ event.Event = &LeftClickedEvent{}

// LeftClickedEvent is an event that is sent when a component is clicked.
type LeftClickedEvent struct {
	X, Y int
}

func NewLeftClickedEvent(x, y int) *LeftClickedEvent {
	return &LeftClickedEvent{
		X: x,
		Y: y,
	}
}

func (e *LeftClickedEvent) Event() string { return LeftMouseClickedEventType }
