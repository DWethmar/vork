package persistence

import (
	"fmt"
	"slices"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/event"
)

// Repository is a interface that defines the methods that a persistence repository should implement.
type System struct {
	eventBus          *event.Bus
	r                 Repository
	changedComponents map[component.ComponentType]component.Component // map of components that have changed by type
	deleteComponents  map[component.ComponentType]component.Component // map of components that have been deleted by type
}

func New(eventBus *event.Bus, r Repository) *System {
	s := &System{
		eventBus:          eventBus,
		r:                 r,
		changedComponents: make(map[component.ComponentType]component.Component),
		deleteComponents:  make(map[component.ComponentType]component.Component),
	}

	persistentComponentTypes := PersistentComponentTypes()

	// subscribe to component change events for all persistent components.
	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		c, ok := e.(component.Event)
		return ok && slices.Contains(persistentComponentTypes, c.Component().Type())
	}), s.componentChangeHandler)

	// subscribe to component delete events for all persistent components.
	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		c, ok := e.(component.Event)
		return ok && slices.Contains(persistentComponentTypes, c.Component().Type()) && c.Deleted()
	}), s.componentDeleteHandler)

	return s
}

func (s *System) componentChangeHandler(e event.Event) error {
	ce, ok := e.(component.Event)
	if !ok {
		return fmt.Errorf("unknown event type: %T", e)
	}
	// check if the component is not already deleted
	if _, ok := s.deleteComponents[ce.Component().Type()]; ok {
		return fmt.Errorf("component %d is already deleted", ce.Component().ID())
	}
	s.changedComponents[ce.Component().Type()] = ce.Component()
	return nil
}

func (s *System) componentDeleteHandler(e event.Event) error {
	ce, ok := e.(component.Event)
	if !ok {
		return fmt.Errorf("unknown event type: %T", e)
	}
	// delete from changed components
	delete(s.changedComponents, ce.Component().Type())
	s.deleteComponents[ce.Component().Type()] = ce.Component()
	return nil
}

func (s *System) Save() error {
	for _, c := range s.changedComponents {
		if err := s.r.Save(c); err != nil {
			return fmt.Errorf("failed to save component: %w", err)
		}
	}
	s.changedComponents = make(map[component.ComponentType]component.Component)

	for _, c := range s.deleteComponents {
		if err := s.r.Delete(c.Type(), c.ID()); err != nil {
			return fmt.Errorf("failed to delete component: %w", err)
		}
	}
	s.deleteComponents = make(map[component.ComponentType]component.Component)
	return nil
}

func (s *System) Load() error {
	return nil
}
