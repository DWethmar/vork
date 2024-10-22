package persistence_test

import (
	"os"
	"testing"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/persistence"
	"github.com/dwethmar/vork/point"
	bolt "go.etcd.io/bbolt"
)

func openTestDB(t *testing.T, path string) *bolt.DB {
	t.Helper()
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		t.Fatalf("Failed to open test DB: %v", err)
	}
	return db
}

func closeTestDB(t *testing.T, db *bolt.DB, path string) {
	t.Helper()
	if err := db.Close(); err != nil {
		t.Errorf("Failed to close DB: %v", err)
	}
	if err := os.Remove(path); err != nil {
		t.Errorf("Failed to remove DB file: %v", err)
	}
}

func TestNew(t *testing.T) {
	t.Run("New should create a new system", func(t *testing.T) {
		eventBus := event.NewBus()
		stores := ecsys.NewStores()
		ecs := ecsys.New(eventBus, stores)
		s := persistence.New(eventBus, stores, ecs)
		if s == nil {
			t.Error("System should not be nil")
		}
	})

	t.Run("New should subscribe to component change events", func(t *testing.T) {
		eventBus := event.NewBus()
		stores := ecsys.NewStores()
		ecs := ecsys.New(eventBus, stores)
		s := persistence.New(eventBus, stores, ecs)
		if s == nil {
			t.Error("System should not be nil")
		}
		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 1 {
			t.Errorf("Expected 1 subscriptions, got %d", len(subscriptions))
		}
	})
}

func TestSystem_Save(t *testing.T) { //nolint: gocognit
	t.Run("Save and load components", func(t *testing.T) {
		path := t.TempDir() + "/test.db"
		db := openTestDB(t, path)
		t.Cleanup(func() {
			closeTestDB(t, db, path)
		})

		eventBus := event.NewBus()
		stores := ecsys.NewStores()
		ecs := ecsys.New(eventBus, stores)
		s := persistence.New(eventBus, stores, ecs)

		// Add components
		for i := 0; i < 100; i++ {
			x := i * 10
			y := i * 10
			if i == 50 {
				x = -100
				y = -100
			}

			// Add position component
			e, err := ecs.CreateEntity(entity.Entity(0), point.New(x, y))
			if err != nil {
				t.Fatalf("Failed to create entity: %v", err)
			}

			// Add controllable component
			ctrl := controllable.New(e)
			if _, err = ecs.AddControllable(*ctrl); err != nil {
				t.Fatalf("Failed to add controllable component: %v", err)
			}

			// Add skeleton component
			skel := skeleton.New(e)
			if _, err = ecs.AddSkeleton(*skel); err != nil {
				t.Fatalf("Failed to add skeleton component: %v", err)
			}
		}

		// Save components
		if err := s.Save(db); err != nil {
			t.Fatalf("Failed to save components: %v", err)
		}

		// Create new ECS and load components
		eventBus = event.NewBus()
		stores = ecsys.NewStores()
		ecs = ecsys.New(eventBus, stores)
		s = persistence.New(eventBus, stores, ecs)
		if err := s.Load(db); err != nil {
			t.Fatalf("Failed to load components: %v", err)
		}

		// Verify components
		for i := 0; i < 100; i++ {
			e := entity.Entity(i + 1)
			x := i * 10
			y := i * 10
			if i == 50 {
				x = -100
				y = -100
			}

			// Verify position component
			pos, err := ecs.GetPosition(e)
			if err != nil {
				t.Fatalf("Failed to get position component: %v", err)
			}
			if nx, ny := pos.Cords(); x != nx || y != ny {
				t.Errorf("Position mismatch for entity %d: expected (%d, %d), got (%d, %d)", e, x, y, nx, ny)
			}

			// Verify controllable component
			if _, err = ecs.GetControllable(e); err != nil {
				t.Fatalf("Failed to get controllable component: %v", err)
			}

			// Verify skeleton component
			if _, err = ecs.GetSkeleton(e); err != nil {
				t.Fatalf("Failed to get skeleton component: %v", err)
			}
		}
	})

	t.Run("Save and delete components", func(t *testing.T) {
		path := t.TempDir() + "/test.db"
		db := openTestDB(t, path)
		t.Cleanup(func() {
			closeTestDB(t, db, path)
		})

		eventBus := event.NewBus()
		stores := ecsys.NewStores()
		ecs := ecsys.New(eventBus, stores)
		s := persistence.New(eventBus, stores, ecs)

		// Add position components
		for i := 1; i <= 100; i++ {
			x := i * 10
			y := i * 10

			_, err := ecs.CreateEntity(entity.Entity(0), point.New(x, y))
			if err != nil {
				t.Fatalf("Failed to create entity: %v", err)
			}
		}

		// Delete position component of entity 50
		if con, err := ecs.GetPosition(entity.Entity(50)); err == nil {
			if err = ecs.DeletePosition(con); err != nil {
				t.Fatalf("Failed to delete position component: %v", err)
			}
		} else {
			t.Fatalf("Failed to get position component: %v", err)
		}

		// Save components
		if err := s.Save(db); err != nil {
			t.Fatalf("Failed to save components: %v", err)
		}

		// Create new ECS and load components
		eventBus = event.NewBus()
		stores = ecsys.NewStores()
		ecs = ecsys.New(eventBus, stores)
		s = persistence.New(eventBus, stores, ecs)
		if err := s.Load(db); err != nil {
			t.Fatalf("Failed to load components: %v", err)
		}

		// Verify components
		for i := 1; i <= 100; i++ {
			e := entity.Entity(i)
			pos, err := ecs.GetPosition(e)
			if i == 50 {
				if err == nil {
					t.Errorf("expected position component for entity %d to be deleted", e)
				}
			} else {
				if err != nil {
					t.Fatalf("failed to get position component: %v", err)
				}
				x := i * 10
				y := i * 10
				if nx, ny := pos.Cords(); x != nx || y != ny {
					t.Errorf("position mismatch for entity %d: expected (%d, %d), got (%d, %d)", e, x, y, nx, ny)
				}
			}
		}
	})
}

func TestSystem_Load(t *testing.T) {
	t.Run("Load should load all components", func(t *testing.T) {
		path := t.TempDir() + "/test.db"
		db := openTestDB(t, path)
		t.Cleanup(func() {
			closeTestDB(t, db, path)
		})

		var e entity.Entity
		{
			eventBus := event.NewBus()
			stores := ecsys.NewStores()
			ecs := ecsys.New(eventBus, stores)
			// create system
			s := persistence.New(eventBus, stores, ecs)
			var err error
			// load some data
			e, err = ecs.CreateEntity(entity.Entity(0), point.New(11, 22))
			if err != nil {
				t.Errorf("CreateEntity failed: %v", err)
			}

			pos, err := ecs.GetPosition(e)
			if err != nil {
				t.Errorf("Position failed: %v", err)
			}

			// update position
			pos.SetCords(33, 44)

			if err = ecs.UpdatePositionComponent(pos); err != nil {
				t.Errorf("UpdatePosition failed: %v", err)
			}

			if err = s.Save(db); err != nil {
				t.Errorf("Load failed: %v", err)
			}
		}

		{
			eventBus := event.NewBus()
			stores := ecsys.NewStores()
			ecs := ecsys.New(eventBus, stores)
			// create system
			s := persistence.New(eventBus, stores, ecs)
			if err := s.Load(db); err != nil {
				t.Errorf("Load failed: %v", err)
			}

			// check ecs for loaded components
			position, err := ecs.GetPosition(e)
			if err != nil {
				t.Errorf("Position failed: %v", err)
			}
			if x, y := position.Cords(); x != 33 || y != 44 {
				t.Errorf("Position failed: expected (33, 44), got (%d, %d)", x, y)
			}
		}
	})
}
