package persistence_test

import (
	"log"
	"os"
	"testing"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/persistence"

	bolt "go.etcd.io/bbolt"
)

func testDB(t *testing.T, path string) (*bolt.DB, error) {
	t.Helper()
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestNew(t *testing.T) {
	t.Run("New should create a new system", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus)
		s := persistence.New(eventBus, ecs)
		if s == nil {
			t.Error("System should not be nil")
		}
	})

	t.Run("New should subscribe to component change events", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus)
		s := persistence.New(eventBus, ecs)
		if s == nil {
			t.Error("System should not be nil")
		}
		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 1 {
			t.Errorf("Expected 1 subscriptions, got %d", len(subscriptions))
		}
	})
}

// saveTest saves and loads components to a database.
// it created entities with the given offset and limit.
func saveTest(t *testing.T, path string, offset, limit int) {
	// SAVE
	save := func() {
		db, err := testDB(t, path)
		if err != nil {
			t.Errorf("testDB failed: %v", err)
		}
		defer db.Close()

		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus)
		// create system
		s := persistence.New(eventBus, ecs)
		for _, pt := range persistence.PersistentComponentTypes() {
			for i := range limit {
				e := entity.Entity(offset + i)
				switch pt {
				case position.Type:
					x := i * 10
					y := i * 10
					if i == 50 { // sanity check
						x = -100
						y = -100
					}
					position := position.New(e, x, y)
					if _, err = ecs.AddPositionComponent(*position); err != nil {
						t.Errorf("AddPosition failed: %v", err)
					}
				case controllable.Type:
					controllable := controllable.New(e)
					if _, err = ecs.AddControllableComponent(*controllable); err != nil {
						t.Errorf("UpdateControllable failed: %v", err)
					}
				case skeleton.Type:
					skeleton := skeleton.New(e)
					if _, err = ecs.AddSkeletonComponent(*skeleton); err != nil {
						t.Errorf("UpdateSkeleton failed: %v", err)
					}
				default:
					t.Fatalf("unknown component type: %s", pt)
				}
			}
		}
		// save all changed components
		if err := s.Save(db); err != nil {
			t.Errorf("Save failed: %v", err)
		}
	}

	// LOAD
	load := func() {
		db, err := testDB(t, path)
		if err != nil {
			t.Errorf("testDB failed: %v", err)
		}
		defer db.Close()

		// create new system
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus)
		s := persistence.New(eventBus, ecs)
		s.Load(db)

		for _, pt := range persistence.PersistentComponentTypes() {
			for i := range limit {
				e := entity.Entity(offset + i)
				switch pt {
				case controllable.Type:
					controllable, err := ecs.Controllable(e)
					if err != nil {
						t.Errorf("Controllable failed: %v", err)
					}
					if controllable.Entity() != e {
						t.Errorf("Controllable entity failed: expected %d, got %d", i, controllable.Entity())
					}
				case position.Type:
					position, err := ecs.Position(e)
					if err != nil {
						t.Errorf("Position failed: %v", err)
					}
					x := i * 10
					y := i * 10
					if i == 50 { // sanity check
						x = -100
						y = -100
					}
					if position.Entity() != e {
						t.Errorf("Position entity failed: expected %d, got %d", i, position.Entity())
					}
					if position.X != x || position.Y != y {
						t.Errorf("Position failed: expected %d, %d, got %d, %d", x, y, position.X, position.Y)
					}
				case skeleton.Type:
					skeleton, err := ecs.Skeleton(e)
					if err != nil {
						t.Errorf("Skeleton failed: %v", err)
					}
					if skeleton.Entity() != e {
						t.Errorf("Skeleton entity failed: expected %d, got %d", i, skeleton.Entity())
					}

				default:
					t.Fatalf("unknown component type: %s", pt)
				}
			}
		}
	}

	save()
	load()
}

func TestSystem_Save(t *testing.T) {
	t.Run("Save should save all added components", func(t *testing.T) {
		path := t.TempDir() + "/test.db"
		saveTest(t, path, 0, 100)
		saveTest(t, path, 0, 100)
		saveTest(t, path, 100, 100)
		saveTest(t, path, 100, 100)
	})

	t.Run("Save should save all deleted components", func(t *testing.T) {
		path := t.TempDir() + "/test.db"
		// SAVE
		save := func() {
			db, err := testDB(t, path)
			if err != nil {
				t.Errorf("testDB failed: %v", err)
			}
			defer db.Close()

			eventBus := event.NewBus()
			ecs := ecsys.New(eventBus)
			// create system
			s := persistence.New(eventBus, ecs)
			for i := range 100 {
				x := i * 10
				y := i * 10
				position := position.New(entity.Entity(i), x, y)
				if _, err = ecs.AddPositionComponent(*position); err != nil {
					t.Errorf("AddPosition failed: %v", err)
				}
			}
			// delete component of entity 50
			position, err := ecs.Position(entity.Entity(50))
			if err != nil {
				t.Errorf("Position failed: %v", err)
			}
			if err = ecs.DeletePositionComponent(position); err != nil {
				t.Errorf("DeletePosition failed: %v", err)
			}
			// save all changed components
			if err := s.Save(db); err != nil {
				t.Errorf("Save failed: %v", err)
			}
		}

		// LOAD
		load := func() {
			db, err := testDB(t, path)
			if err != nil {
				t.Errorf("testDB failed: %v", err)
			}
			defer db.Close()

			// create new system
			eventBus := event.NewBus()
			ecs := ecsys.New(eventBus)
			s := persistence.New(eventBus, ecs)
			s.Load(db)
			// check ecs for saved component
			for i := range 100 {
				_, err := ecs.Position(entity.Entity(i))
				if i == 50 {
					if err == nil {
						t.Errorf("expected entity 50 to be deleted")
					}
				} else {
					if err != nil {
						t.Errorf("Position failed: %v", err)
					}
				}
			}
		}

		save()
		load()
	})
}

func TestSystem_Load(t *testing.T) {
	t.Run("Load should load all components", func(t *testing.T) {
		path := t.TempDir() + "/test.db"
		db, err := testDB(t, path)
		if err != nil {
			t.Errorf("testDB failed: %v", err)
		}
		defer db.Close()

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

		// load some data
		entity, err := ecs.CreateEntity(11, 22)
		if err != nil {
			t.Errorf("CreateEntity failed: %v", err)
		}

		position, err := ecs.Position(entity)
		if err != nil {
			t.Errorf("Position failed: %v", err)
		}

		// update position
		position.X = 33
		position.Y = 44

		if err := ecs.UpdatePositionComponent(position); err != nil {
			t.Errorf("UpdatePosition failed: %v", err)
		}

		// create system
		s := persistence.New(eventBus, ecs)
		if err := s.Load(db); err != nil {
			t.Errorf("Load failed: %v", err)
		}

		// check ecs for loaded components
		position, err = ecs.Position(entity)
		if err != nil {
			t.Errorf("Position failed: %v", err)
		}
		if position.X != 33 || position.Y != 44 {
			t.Errorf("Position failed: expected 11, 11, got %d, %d", position.X, position.Y)
		}
	})
}
