package persistence

import (
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	boltrepo "github.com/dwethmar/vork/systems/persistence/bbolt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	bolt "go.etcd.io/bbolt"
)

// Persistance saves and loads components from the database.
type Persistance struct {
	logger     *slog.Logger
	eventBus   *event.Bus
	ecs        *ecsys.ECS
	lifecycles map[component.Type]ComponentLifeCycle
	stores     *ecsys.Stores
	db         *bolt.DB
}

// Options is the configuration for the persistence system.
type Options struct {
	// Logger is the logger used by the persistence system.
	Logger *slog.Logger
	// EventBus is the event bus used by the persistence system.
	EventBus *event.Bus
	// Stores is the component stores used by the persistence system.
	Stores *ecsys.Stores
	// ECS is the ECS system used by the persistence system.
	ECS *ecsys.ECS

	DB *bolt.DB
}

// New creates a new persistence system.
func New(opts Options) *Persistance {
	stores := opts.Stores
	s := &Persistance{
		logger:   opts.Logger.With("system", "persistence"),
		eventBus: opts.EventBus,
		ecs:      opts.ECS,
		stores:   opts.Stores,
		lifecycles: map[component.Type]ComponentLifeCycle{
			controllable.Type: NewGenericComponentLifeCycle(
				boltrepo.NewRepository(controllable.Empty),
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
				boltrepo.NewRepository(position.Empty),
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
			velocity.Type: NewGenericComponentLifeCycle(
				boltrepo.NewRepository(velocity.Empty),
				stores.Velocity,
				func(c component.Component, m map[uint]*velocity.Velocity) error {
					cc, ok := c.(*velocity.Velocity)
					if !ok {
						return fmt.Errorf("expected *velocity.Velocity, got %T", c)
					}
					m[c.ID()] = cc
					return nil
				},
			),
			skeleton.Type: NewGenericComponentLifeCycle(
				boltrepo.NewRepository(skeleton.Empty),
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
		db: opts.DB,
	}

	persistentComponentTypes := PersistentComponentTypes()

	// subscribe to component change events for all persistent components.
	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		c, ok := e.(component.Event)
		return ok && slices.Contains(persistentComponentTypes, c.ComponentType())
	}), s.componentChangeHandler)

	s.logger.Info("persistence system created", "persistent_components", persistentComponentTypes)

	return s
}

func (s *Persistance) Init() error {
	if err := s.Load(s.db); err != nil {
		return fmt.Errorf("failed to load game: %w", err)
	}
	if err := s.ecs.BuildHierarchy(); err != nil {
		return fmt.Errorf("failed to rebuild hierarchy: %w", err)
	}
	return nil
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
	case velocity.Event:
		if err := l.Changed(c.Velocity(), c.Deleted()); err != nil {
			return fmt.Errorf("failed to mark velocity component as changed: %w", err)
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
		return fmt.Errorf("no handler for component type: %T", c)
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

func (s *Persistance) Close() error {
	return nil
}

func (s *Persistance) Draw(*ebiten.Image) error {
	return nil
}

func (s *Persistance) Update() error {
	// check if F5 is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		started := time.Now()
		if err := s.Save(s.db); err != nil {
			return fmt.Errorf("failed to save game: %w", err)
		}
		s.logger.Info("game saved", slog.Duration("duration", time.Since(started)))
		return nil
	}
	return nil
}
