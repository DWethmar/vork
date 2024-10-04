package event

import (
	"sync"
)

// Event is an interface that requires implementing the Event method.
type Event interface {
	Event() string
}

// EventHandler is a custom type for handler functions that process events.
type EventHandler func(Event) error

// Subscription is a struct that represents a handler subscribed to a specific matching.
type Subscription struct {
	id      int
	matcher Matcher
	handler EventHandler
}

// Bus is a struct that manages event handlers in a thread-safe manner.
type Bus struct {
	handlers []Subscription
	nextID   int // Used to assign a unique ID to each handler
	mu       sync.RWMutex
}

// NewBus creates and returns a new Bus instance.
func NewBus() *Bus {
	return &Bus{
		handlers: []Subscription{},
		nextID:   1, // Start IDs from 1
	}
}

// Subscribe adds a new handler function to the Bus for a specific event type.
// It returns an identifier for the handler, which can be used to unsubscribe it later.
func (b *Bus) Subscribe(m Matcher, handler EventHandler) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	id := b.nextID
	b.nextID++
	b.handlers = append(b.handlers, Subscription{
		id:      id,
		matcher: m,
		handler: handler,
	})
	return id
}

// Unsubscribe removes a handler function from the Bus for a specific event type using its identifier.
func (b *Bus) Unsubscribe(event string, id int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i, entry := range b.handlers {
		if entry.id == id {
			b.handlers = append(b.handlers[:i], b.handlers[i+1:]...)
			break
		}
	}
}

func (b *Bus) Subscriptions() []Subscription {
	b.mu.RLock()
	defer b.mu.RUnlock()
	subscriptions := make([]Subscription, len(b.handlers))
	copy(subscriptions, b.handlers)
	return subscriptions
}

// Publish sends an event to all the handlers subscribed to the event's type.
func (b *Bus) Publish(event Event) error {
	b.mu.RLock()
	handlers := make([]Subscription, len(b.handlers))
	copy(handlers, b.handlers)
	b.mu.RUnlock()

	for _, entry := range handlers {
		if !entry.matcher.Match(event) {
			continue
		}
		if err := entry.handler(event); err != nil {
			return err
		}
	}

	return nil
}
