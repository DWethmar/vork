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
	lifecycles map[component.Type]ComponentLifeCycle
	stores     *ecsys.Stores
}

// New creates a new persistence system.
func New(eventBus *event.Bus, stores *ecsys.Stores, ecs *ecsys.ECS) *Persistance {
	controllableRepo := boltrepo.NewRepository(controllable.Empty)
	positionRepo := boltrepo.NewRepository(position.Empty)
	skeletonRepo := boltrepo.NewRepository(skeleton.Empty)

	s := &Persistance{
		eventBus: eventBus,
		ecs:      ecs,
		stores:   stores,
		lifecycles: map[component.Type]ComponentLifeCycle{
			controllable.Type: NewGenericComponentLifeCycle(
				controllableRepo,
				stores.Controllable,
				func(c component.Component, m map[uint]*controllable.Controllable) error {
					cc, ok := c.(*controllable.Controllable)
					if !ok {
						return fmt.Errorf("expected *controllable.Controllable, got %T", c)
					}
					m[c.ID()] = cc
					return nil
				},
			),
			position.Type: NewGenericComponentLifeCycle(
				positionRepo,
				stores.Position,
				func(c component.Component, m map[uint]*position.Position) error {
					cc, ok := c.(*position.Position)
					if !ok {
						return fmt.Errorf("expected *position.Position, got %T", c)
					}
					m[c.ID()] = cc
					return nil
				},
			),
			skeleton.Type: NewGenericComponentLifeCycle(
				skeletonRepo,
				stores.Skeleton,
				func(e component.Component, m map[uint]*skeleton.Skeleton) error {
					cc, ok := e.(*skeleton.Skeleton)
					if !ok {
						return fmt.Errorf("expected *skeleton.Skeleton, got %T", e)
					}
					m[e.ID()] = cc
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

	switch c := ce.(type) {
	case position.Event:
		if err := l.Changed(c.Position(), c.Deleted()); err != nil {
			return fmt.Errorf("failed to mark position component as changed: %w", err)
		}
	case controllable.Event:
		if err := l.Changed(c.Controllable(), c.Deleted()); err != nil {
			return fmt.Errorf("failed to mark controllable component as changed: %w", err)
		}
	case skeleton.Event:
		if err := l.Changed(c.Skeleton(), c.Deleted()); err != nil {
			return fmt.Errorf("failed to mark skeleton component as changed: %w", err)
		}
	default:
		return fmt.Errorf("unsupported component type: %s", ce.ComponentType())
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
			if err := l.Load(tx); err != nil {
				return fmt.Errorf("failed to load controllable components: %w", err)
			}
		}

		return nil
	})
}
