package persistence_test

import (
	"log"
	"os"
	"testing"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems/persistence"

	boltrepo "github.com/dwethmar/vork/systems/persistence/bbolt"
	bolt "go.etcd.io/bbolt"
)

func TestNew(t *testing.T) {
	t.Run("New should create a new system", func(t *testing.T) {
		eventBus := event.NewBus()
		s := persistence.New(eventBus, persistence.Repositories{})
		if s == nil {
			t.Error("System should not be nil")
		}
	})

	t.Run("New should subscribe to component change events", func(t *testing.T) {
		eventBus := event.NewBus()
		s := persistence.New(eventBus, persistence.Repositories{})
		if s == nil {
			t.Error("System should not be nil")
		}

		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 2 {
			t.Errorf("Expected 2 subscriptions, got %d", len(subscriptions))
		}
	})
}

func TestSystem_Save(t *testing.T) {
	t.Run("Save should save all changed components", func(t *testing.T) {
		path := t.TempDir() + "/test.db"
		db, err := bolt.Open(path, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := db.Close(); err != nil {
				log.Fatal(err)
			}
			if err := os.Remove(path); err != nil {
				log.Fatal(err)
			}
		}()

		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus)
		// create dummy entity to offset the id
		ecs.CreateEntity()

		// create system
		repos := persistence.Repositories{
			ControllableRepo: boltrepo.NewRepository(db, func() *controllable.Controllable {
				return controllable.New(0)
			}),
			PositionRepo: boltrepo.NewRepository(db, func() *position.Position {
				return position.New(0, 0, 0)
			}),
			SkeletonRepo: boltrepo.NewRepository(db, func() *skeleton.Skeleton {
				return skeleton.New(0)
			}),
		}

		// create system
		s := persistence.New(eventBus, repos)

		// Create a new component
		entity := ecs.CreateEntity()
		position := position.New(entity, 0, 0)
		id, err := ecs.AddPosition(*position)
		if err != nil {
			t.Errorf("AddPosition failed: %v", err)
		}

		// save all changed components
		if err := s.Save(); err != nil {
			t.Errorf("Save failed: %v", err)
		}

		c, err := repos.PositionRepo.Get(id)
		if err != nil {
			t.Errorf("Get failed: %v", err)
		}

		if c.ID() != id {
			t.Errorf("Get failed: expected %d, got %d", id, c.ID())
		}
	})
}

func TestSystem_Load(t *testing.T) {

}
