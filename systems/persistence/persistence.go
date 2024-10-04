package persistence

import (
	"fmt"
	"slices"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ systems.System = &System{}

type Repositories struct {
	ControllableRepo Repository[*controllable.Controllable]
	PositionRepo     Repository[*position.Position]
	SkeletonRepo     Repository[*skeleton.Skeleton]
}

// Repository is a interface that defines the methods that a persistence repository should implement.
type System struct {
	eventBus          *event.Bus
	repos             Repositories
	changedComponents map[component.ComponentType]component.Component // map of components that have changed by type
	deleteComponents  map[component.ComponentType]component.Component // map of components that have been deleted by type
}

func New(eventBus *event.Bus, r Repositories) *System {
	s := &System{
		eventBus:          eventBus,
		repos:             r,
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
	fmt.Printf("number of changed components: %d\n", len(s.changedComponents))
	for _, c := range s.changedComponents {
		switch t := c.(type) {
		case controllable.Controllable:
			if err := s.repos.ControllableRepo.Save(&t); err != nil {
				return fmt.Errorf("failed to save controllable component: %w", err)
			}
		case position.Position:
			if err := s.repos.PositionRepo.Save(&t); err != nil {
				return fmt.Errorf("failed to save position component: %w", err)
			}
		case skeleton.Skeleton:
			if err := s.repos.SkeletonRepo.Save(&t); err != nil {
				return fmt.Errorf("failed to save skeleton component: %w", err)
			}
		default:
			return fmt.Errorf("unknown component type: %T", c)
		}
	}
	s.changedComponents = make(map[component.ComponentType]component.Component)

	for _, c := range s.deleteComponents {
		switch t := c.(type) {
		case controllable.Controllable:
			if err := s.repos.ControllableRepo.Delete(t.ID()); err != nil {
				return fmt.Errorf("failed to delete controllable component: %w", err)
			}
		case position.Position:
			if err := s.repos.PositionRepo.Delete(t.ID()); err != nil {
				return fmt.Errorf("failed to delete position component: %w", err)
			}
		case skeleton.Skeleton:
			if err := s.repos.SkeletonRepo.Delete(t.ID()); err != nil {
				return fmt.Errorf("failed to delete skeleton component: %w", err)
			}
		default:
			return fmt.Errorf("unknown component type: %T", c)
		}
	}
	s.deleteComponents = make(map[component.ComponentType]component.Component)
	return nil
}

func (s *System) Load() error {
	return nil
}

func (s *System) Update() error {
	return nil
}

func (s *System) Draw(_ *ebiten.Image) error {
	return nil
}
