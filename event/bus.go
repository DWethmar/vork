package event

import (
	"fmt"
	"sync"
)

// Event is an interface that requires implementing the Event method.
type Event interface {
	Event() string
}

// EventHandler is a custom type for handler functions that process events.
type EventHandler func(Event) error

// handlerEntry holds an EventHandler with a unique identifier.
type handlerEntry struct {
	id      int
	handler EventHandler
}

// Bus is a struct that manages event handlers in a thread-safe manner.
type Bus struct {
	handlers map[string][]handlerEntry
	nextID   int // Used to assign a unique ID to each handler
	mu       sync.RWMutex
}

// NewBus creates and returns a new Bus instance.
func NewBus() *Bus {
	return &Bus{
		handlers: make(map[string][]handlerEntry),
		nextID:   1, // Start IDs from 1
	}
}

// Subscribe adds a new handler function to the Bus for a specific event type.
// It returns an identifier for the handler, which can be used to unsubscribe it later.
func (b *Bus) Subscribe(event string, handler EventHandler) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	id := b.nextID
	b.nextID++
	b.handlers[event] = append(b.handlers[event], handlerEntry{id: id, handler: handler})
	return id
}

// Unsubscribe removes a handler function from the Bus for a specific event type using its identifier.
func (b *Bus) Unsubscribe(event string, id int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	handlers := b.handlers[event]
	for i, entry := range handlers {
		if entry.id == id {
			b.handlers[event] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
}

// Publish sends an event to all the handlers subscribed to the event's type.
func (b *Bus) Publish(event Event) error {
	b.mu.RLock()
	handlers, ok := b.handlers[event.Event()]
	b.mu.RUnlock()

	if !ok {
		fmt.Printf("Warning: No handlers for event %s\n", event.Event())
		return nil
	}

	for _, entry := range handlers {
		if err := entry.handler(event); err != nil {
			return err
		}
	}

	return nil
}
