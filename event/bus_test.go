package event

import (
	"testing"
)

// MockEvent is a simple implementation of the Event interface for testing.
type MockEvent struct {
	name string
}

func (e *MockEvent) Event() string {
	return e.name
}

func TestNewBus(t *testing.T) {
	bus := NewBus()
	if bus == nil {
		t.Fatalf("NewBus() returned nil")
	}

	if len(bus.handlers) != 0 {
		t.Errorf("NewBus() handlers map should be empty initially, got %d", len(bus.handlers))
	}
}

func TestBus_Subscribe(t *testing.T) {
	bus := NewBus()
	handlerCalled := false

	bus.Subscribe("testEvent", func(e Event) error {
		handlerCalled = true
		return nil
	})

	if len(bus.handlers) != 1 {
		t.Errorf("Bus.Subscribe() should add a handler, got %d", len(bus.handlers))
	}

	// Trigger the event to check if the handler is called.
	bus.Publish(&MockEvent{name: "testEvent"})

	if !handlerCalled {
		t.Errorf("Handler was not called after Bus.Publish()")
	}
}

func TestBus_Publish(t *testing.T) {
	t.Run("publish an event", func(t *testing.T) {
		bus := NewBus()
		handlerCalled := false

		bus.Subscribe("testEvent", func(e Event) error {
			handlerCalled = true
			return nil
		})

		bus.Publish(&MockEvent{name: "testEvent"})

		if !handlerCalled {
			t.Errorf("Bus.Publish() did not trigger the handler")
		}

		// Test with an event that has no handlers
		handlerCalled = false
		bus.Publish(&MockEvent{name: "nonExistentEvent"})

		if handlerCalled {
			t.Errorf("Bus.Publish() should not trigger a handler for an event with no subscribers")
		}
	})
}

func TestBus_Unsubscribe(t *testing.T) {
	bus := NewBus()
	handlerCalled := false

	handler := func(e Event) error {
		handlerCalled = true
		return nil
	}

	id := bus.Subscribe("testEvent", handler)

	bus.Publish(&MockEvent{name: "testEvent"})

	if !handlerCalled {
		t.Errorf("Handler was not called after Bus.Publish()")
	}

	handlerCalled = false
	bus.Unsubscribe("testEvent", id)

	bus.Publish(&MockEvent{name: "testEvent"})

	if handlerCalled {
		t.Errorf("Handler was called after Bus.Unsubscribe()")
	}
}
