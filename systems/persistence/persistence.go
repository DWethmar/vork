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

type ComponentLifeCycle interface {
	Changed(component.Event) error // Changed is called when a component has changed.
	Deleted(component.Event) error // Deleted is called when a component has been deleted.
	Commit() error                 // Commit saves all changes to the database.
	// Load(*ecsys.ECS) error         // Load function added here
}

// Repository is a interface that defines the methods that a persistence repository should implement.
type System struct {
	eventBus   *event.Bus
	ecs        *ecsys.ECS
	repos      Repositories
	lifecycles map[component.ComponentType]ComponentLifeCycle
}

func New(eventBus *event.Bus, ecs *ecsys.ECS, r Repositories) *System {
	s := &System{
		eventBus: eventBus,
		ecs:      ecs,
		repos:    r,
		lifecycles: map[component.ComponentType]ComponentLifeCycle{
			controllable.Type: NewControllableLifeCycle(r.ControllableRepo),
			position.Type:     NewPositionLifeCycle(r.PositionRepo),
			skeleton.Type:     NewSkeletonLifeCycle(r.SkeletonRepo),
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

	l, ok := s.lifecycles[ce.ComponentType()]
	if !ok {
		return fmt.Errorf("no lifecycle for component type: %s", ce.ComponentType())
	}

	if err := l.Changed(ce); err != nil {
		return fmt.Errorf("failed to handle changed event for component type %s (ID: %d): %w", ce.ComponentType(), ce.ComponentID(), err)
	}

	return nil
}

func (s *System) componentDeleteHandler(e event.Event) error {
	ce, ok := e.(component.Event)
	if !ok {
		return fmt.Errorf("unknown event type: %T", e)
	}

	l, ok := s.lifecycles[ce.ComponentType()]
	if !ok {
		return fmt.Errorf("no lifecycle for component type: %s", ce.ComponentType())
	}

	if err := l.Deleted(ce); err != nil {
		return fmt.Errorf("failed to handle deleted event for component type %s (ID: %d): %w", ce.ComponentType(), ce.ComponentID(), err)
	}

	return nil
}

// Save saves all changed components to the database.
func (s *System) Save() error {
	for _, l := range s.lifecycles {
		if err := l.Commit(); err != nil {
			return fmt.Errorf("failed to commit lifecycle: %w", err)
		}
	}
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
