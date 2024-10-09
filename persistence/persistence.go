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

	boltrepo "github.com/dwethmar/vork/persistence/bbolt"
	bolt "go.etcd.io/bbolt"
)

// Persistance saves and loads components from the database.
type Persistance struct {
	eventBus   *event.Bus
	ecs        *ecsys.ECS
	lifecycles map[component.ComponentType]ComponentLifeCycle
}

// New creates a new persistence system.
func New(eventBus *event.Bus, ecs *ecsys.ECS) *Persistance {
	controllableRepo := boltrepo.NewRepository(func() *controllable.Controllable {
		return controllable.New(0)
	})

	positionRepo := boltrepo.NewRepository(func() *position.Position {
		return position.New(0, 0, 0)
	})

	skeletonRepo := boltrepo.NewRepository(func() *skeleton.Skeleton {
		return skeleton.New(0)
	})

	s := &Persistance{
		eventBus: eventBus,
		ecs:      ecs,
		lifecycles: map[component.ComponentType]ComponentLifeCycle{
			controllable.Type: NewGenericComponentLifeCycle(
				controllableRepo,
				func(c *controllable.Controllable) (uint, error) { // TODO refactor
					return ecs.AddControllableComponent(*c)
				},
				func(e component.Event, m map[uint]*controllable.Controllable) error {
					c, ok := e.(controllable.Event)
					if !ok {
						return fmt.Errorf("expected %T, got %T", c, e)
					}
					m[c.ComponentID()] = c.Controllable()
					return nil
				},
			),
			position.Type: NewGenericComponentLifeCycle(
				positionRepo,
				func(c *position.Position) (uint, error) {
					return ecs.AddPositionComponent(*c)
				},
				func(e component.Event, m map[uint]*position.Position) error {
					c, ok := e.(position.Event)
					if !ok {
						return fmt.Errorf("expected %T, got %T", c, e)
					}
					m[c.ComponentID()] = c.Position()
					return nil
				},
			),
			skeleton.Type: NewGenericComponentLifeCycle(
				skeletonRepo,
				func(c *skeleton.Skeleton) (uint, error) {
					return ecs.AddSkeletonComponent(*c)
				},
				func(e component.Event, m map[uint]*skeleton.Skeleton) error {
					c, ok := e.(skeleton.Event)
					if !ok {
						return fmt.Errorf("expected %T, got %T", c, e)
					}
					m[c.ComponentID()] = c.Skeleton()
					return nil
				},
			),
		},
	}

	persistentComponentTypes := PersistentComponentTypes()

	// subscribe to component change events for all persistent components.
	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		c, ok := e.(component.Event)
		return ok && slices.Contains(persistentComponentTypes, c.ComponentType())
	}), s.componentChangeHandler)

	return s
}

// componentChangeHandler is called when a component has changed or has been deleted.
func (s *Persistance) componentChangeHandler(e event.Event) error {
	ce, ok := e.(component.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", ce, e)
	}
	l, ok := s.lifecycles[ce.ComponentType()]
	if !ok {
		return fmt.Errorf("no lifecycle for component type: %s", ce.ComponentType())
	}
	if ce.Deleted() {
		if err := l.Deleted(ce); err != nil {
			return fmt.Errorf("failed to handle deleted event for component type %s (ID: %d): %w", ce.ComponentType(), ce.ComponentID(), err)
		}
		return nil
	} else if err := l.Changed(ce); err != nil {
		return fmt.Errorf("failed to handle changed event for component type %s (ID: %d): %w", ce.ComponentType(), ce.ComponentID(), err)
	}
	return nil
}

// Save saves all changed or deleted components to the database.
func (s *Persistance) Save(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		for n, l := range s.lifecycles {
			if err := l.Commit(tx); err != nil {
				return fmt.Errorf("failed to commit changes for component type %s: %w", n, err)
			}
		}
		return nil
	})
}

// Load loads all components from the database and adds them to the ECS.
func (s *Persistance) Load(db *bolt.DB) error {
	return db.View(func(tx *bolt.Tx) error {
		for _, r := range PersistentComponentTypes() {
			l, ok := s.lifecycles[r]
			if !ok {
				return fmt.Errorf("no lifecycle for component type: %s", r)
			}
			if err := l.Load(tx, s.ecs); err != nil {
				return fmt.Errorf("failed to load controllable components: %w", err)
			}
		}

		return nil
	})
}
