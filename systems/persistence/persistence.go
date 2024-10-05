package persistence

import (
	"fmt"
	"slices"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
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

type componentKey struct {
	componentType component.ComponentType
	componentID   uint32
}

type saveQueuerHandler func(event.Event) error

type saveHandlerFunc func(Repositories, component.Component) error
type deleteHandlerFunc func(Repositories, component.Component) error

// Repository is a interface that defines the methods that a persistence repository should implement.
type System struct {
	eventBus          *event.Bus
	ecs               *ecsys.ECS
	repos             Repositories
	changedComponents map[componentKey]component.Component // map of components that have changed by type
	deleteComponents  map[componentKey]component.Component // map of components that have been deleted by type

	// lifecycle handlers
	saveHandlers   map[component.ComponentType]saveHandlerFunc
	deleteHandlers map[component.ComponentType]deleteHandlerFunc
}

type ComponentLifeCycle interface {
	Changed(component.Event) error
	Deleted(component.Event) error
	Save(component.Component) error
	Delete(component.Component) error
}

func New(eventBus *event.Bus, ecs *ecsys.ECS, r Repositories) *System {
	s := &System{
		eventBus:          eventBus,
		ecs:               ecs,
		repos:             r,
		changedComponents: make(map[componentKey]component.Component),
		deleteComponents:  make(map[componentKey]component.Component),
		saveHandlers: map[component.ComponentType]saveHandlerFunc{
			controllable.Type: func(r Repositories, c component.Component) error {
				comp, ok := c.(*controllable.Controllable)
				if !ok {
					return fmt.Errorf("expected %T, got %T", comp, c)
				}
				return r.ControllableRepo.Save(comp)
			},
			position.Type: func(r Repositories, c component.Component) error {
				comp, ok := c.(*position.Position)
				if !ok {
					return fmt.Errorf("expected %T, got %T", comp, c)
				}
				return r.PositionRepo.Save(comp)
			},
			skeleton.Type: func(r Repositories, c component.Component) error {
				comp, ok := c.(*skeleton.Skeleton)
				if !ok {
					return fmt.Errorf("expected %T, got %T", comp, c)
				}
				return r.SkeletonRepo.Save(comp)
			},
		},
		deleteHandlers: map[component.ComponentType]deleteHandlerFunc{
			controllable.Type: func(r Repositories, c component.Component) error {
				comp, ok := c.(*controllable.Controllable)
				if !ok {
					return fmt.Errorf("expected %T, got %T", comp, c)
				}
				return r.ControllableRepo.Delete(comp.ID())
			},
			position.Type: func(r Repositories, c component.Component) error {
				comp, ok := c.(*position.Position)
				if !ok {
					return fmt.Errorf("expected %T, got %T", comp, c)
				}
				return r.PositionRepo.Delete(comp.ID())
			},
			skeleton.Type: func(r Repositories, c component.Component) error {
				comp, ok := c.(*skeleton.Skeleton)
				if !ok {
					return fmt.Errorf("expected %T, got %T", comp, c)
				}
				return r.SkeletonRepo.Delete(comp.ID())
			},
		},
	}

	persistentComponentTypes := PersistentComponentTypes()

	// subscribe to component change events for all persistent components.
	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		c, ok := e.(component.Event)
		return ok && slices.Contains(persistentComponentTypes, c.ComponentType())
	}), s.componentChangeHandler)

	// subscribe to component delete events for all persistent components.
	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		c, ok := e.(component.Event)
		return ok && slices.Contains(persistentComponentTypes, c.ComponentType()) && c.Deleted()
	}), s.componentDeleteHandler)

	return s
}

func (s *System) componentChangeHandler(e event.Event) error {
	ce, ok := e.(component.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", ce, e)
	}
	key := componentKey{
		componentType: ce.ComponentType(),
		componentID:   ce.ComponentID(),
	}
	// check if the component is not already deleted
	if _, ok := s.deleteComponents[key]; ok {
		return fmt.Errorf("component %d is already deleted", ce.ComponentID())
	}
	var c component.Component
	switch ce.ComponentType() {
	case controllable.Type:
		e, ok := ce.(controllable.Event)
		if !ok {
			return fmt.Errorf("expected %T, got %T", c, ce)
		}
		c = e.Controllable()
	case position.Type:
		e, ok := ce.(position.Event)
		if !ok {
			return fmt.Errorf("expected %T, got %T", c, ce)
		}
		c = e.Position()
	case skeleton.Type:
		e, ok := ce.(skeleton.Event)
		if !ok {
			return fmt.Errorf("expected %T, got %T", c, ce)
		}
		c = e.Skeleton()
	default:
		return fmt.Errorf("unknown component type: %s", ce.ComponentType())
	}

	s.changedComponents[key] = c
	return nil
}

func (s *System) componentDeleteHandler(e event.Event) error {
	ce, ok := e.(component.Event)
	if !ok {
		return fmt.Errorf("unknown event type: %T", e)
	}
	// delete from changed components
	var c component.Component
	switch ce.ComponentType() {
	case controllable.Type:
		e, ok := ce.(controllable.Event)
		if !ok {
			return fmt.Errorf("expected %T, got %T", c, ce)
		}
		c = e.Controllable()
	case position.Type:
		e, ok := ce.(position.Event)
		if !ok {
			return fmt.Errorf("expected %T, got %T", c, ce)
		}
		c = e.Position()
	case skeleton.Type:
		e, ok := ce.(skeleton.Event)
		if !ok {
			return fmt.Errorf("expected %T, got %T", c, ce)
		}
		c = e.Skeleton()
	default:
		return fmt.Errorf("unknown component type: %s", ce.ComponentType())
	}
	key := componentKey{
		componentType: ce.ComponentType(),
		componentID:   ce.ComponentID(),
	}
	delete(s.changedComponents, key)
	s.deleteComponents[key] = c
	return nil
}

// Save saves all changed components to the database.
func (s *System) Save() error {
	fmt.Printf("number of changed components: %d\n", len(s.changedComponents))
	for key, c := range s.changedComponents {
		handler, exists := s.saveHandlers[key.componentType]
		if !exists {
			return fmt.Errorf("no save handler for component type: %s", key.componentType)
		}
		if err := handler(s.repos, c); err != nil {
			return err
		}
	}
	s.changedComponents = make(map[componentKey]component.Component)

	for key, c := range s.deleteComponents {
		handler, exists := s.deleteHandlers[key.componentType]
		if !exists {
			return fmt.Errorf("no delete handler for component type: %s", key.componentType)
		}
		if err := handler(s.repos, c); err != nil {
			return err
		}
	}
	s.deleteComponents = make(map[componentKey]component.Component)

	return nil
}

func (s *System) Load() error {
	for _, r := range PersistentComponentTypes() {
		switch r {
		case controllable.Type:
			l, err := s.repos.ControllableRepo.List()
			if err != nil {
				return fmt.Errorf("failed to list controllable components: %w", err)
			}
			for _, c := range l {
				if _, err := s.ecs.AddControllable(*c); err != nil {
					return fmt.Errorf("failed to add controllable component: %w", err)
				}
			}
		case position.Type:
			l, err := s.repos.PositionRepo.List()
			if err != nil {
				return fmt.Errorf("failed to list position components: %w", err)
			}
			for _, c := range l {
				if _, err := s.ecs.AddPosition(*c); err != nil {
					return fmt.Errorf("failed to add position component: %w", err)
				}
			}
		case skeleton.Type:
			l, err := s.repos.SkeletonRepo.List()
			if err != nil {
				return fmt.Errorf("failed to list skeleton components: %w", err)
			}
			for _, c := range l {
				if _, err := s.ecs.AddSkeleton(*c); err != nil {
					return fmt.Errorf("failed to add skeleton component: %w", err)
				}
			}
		default:
			return fmt.Errorf("unknown component type: %s", r)
		}
	}
	return nil
}

func (s *System) Update() error {
	return nil
}

func (s *System) Draw(_ *ebiten.Image) error {
	return nil
}
