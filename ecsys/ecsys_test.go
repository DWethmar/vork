// package ecsys contains the Entity-Component-System architecture.
package ecsys_test

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/point"
	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus(), ecsys.NewStores())
		if ecs == nil {
			t.Errorf("New() ecs = %v", ecs)
		}
	})
}

func TestECS_CreateEntity(t *testing.T) {
	t.Run("should create an entity with a position", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus(), ecsys.NewStores())
		entity, err := ecs.CreateEntity(entity.Entity(0), point.New(11, 22))
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}
		if entity == 0 {
			t.Errorf("CreateEntity() entity = %v", entity)
		}

		// Check if the entity has a position component
		pos, err := ecs.GetPosition(entity)
		if err != nil {
			t.Errorf("GetPosition() error = %v", err)
		}

		expected := position.Position{
			I:     1,
			E:     entity,
			Point: point.New(11, 22),
		}

		if diff := cmp.Diff(pos, expected); diff != "" {
			t.Errorf("GetPosition() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestECS_DeleteEntity(t *testing.T) {
	t.Run("should delete an entity", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus(), ecsys.NewStores())
		entity, err := ecs.CreateEntity(entity.Entity(0), point.New(11, 22))
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}

		// add sprite component
		if _, err = ecs.AddSprite(sprite.Sprite{
			E:       entity,
			I:       1,
			Graphic: sprite.SkeletonDeath1,
		}); err != nil {
			t.Errorf("AddSprite() error = %v", err)
		}

		// check if the sprite has been added
		if l := len(ecs.ListSpritesByEntity(entity)); l != 1 {
			t.Errorf("SpritesByEntity() sprites = %v", l)
		}

		if err = ecs.DeleteEntity(entity); err != nil {
			t.Errorf("DeleteEntity() error = %v", err)
		}

		// Check if the entity has been deleted
		_, err = ecs.GetPosition(entity)
		if !errors.Is(err, ecsys.ErrEntityNotFound) {
			t.Errorf("GetPosition() error = %v", err)
		}

		// Check if the sprite has been deleted
		if l := len(ecs.ListSpritesByEntity(entity)); l != 0 {
			t.Errorf("SpritesByEntity() sprites = %v", l)
		}
	})
}

func TestECS_CreateEmptyEntity(t *testing.T) {
	ecs := ecsys.New(event.NewBus(), ecsys.NewStores())
	for i := range 100 {
		if e := ecs.CreateEmptyEntity(); e != entity.Entity(i+1) {
			t.Errorf("expected entity %d, got %d", i+1, e)
		}
	}
}

func TestECS_GetAbsolutePosition(t *testing.T) {
	t.Run("should get the absolute position of an entity", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus(), ecsys.NewStores())

		child1, err := ecs.CreateEntity(entity.Entity(0), point.New(11, 22))
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}

		child2, err := ecs.CreateEntity(child1, point.New(1, 2))
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}

		child3, err := ecs.CreateEntity(child2, point.New(3, 4))
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}

		p, err := ecs.GetAbsolutePosition(child3)
		if err != nil {
			t.Errorf("GetAbsolutePosition() error = %v", err)
		}

		expectedX := 15
		expectedY := 28

		if p.X != expectedX {
			t.Errorf("expected x = %v, got %v", expectedX, p.X)
		}

		if p.Y != expectedY {
			t.Errorf("expected y = %v, got %v", expectedY, p.Y)
		}
	})
}
