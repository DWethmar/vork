// package ecsys contains the Entity-Component-System architecture.
package ecsys_test

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus())
		if ecs == nil {
			t.Errorf("New() ecs = %v", ecs)
		}
	})
}

func TestECS_CreateEntity(t *testing.T) {
	t.Run("should create an entity with a position", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus())
		entity, err := ecs.CreateEntity(11, 22)
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}
		if entity == 0 {
			t.Errorf("CreateEntity() entity = %v", entity)
		}

		// Check if the entity has a position component
		pos, err := ecs.Position(entity)
		if err != nil {
			t.Errorf("GetPosition() error = %v", err)
		}

		expected := position.Position{
			I: 1,
			E: entity,
			X: 11,
			Y: 22,
		}

		if diff := cmp.Diff(pos, expected); diff != "" {
			t.Errorf("GetPosition() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestECS_DeleteEntity(t *testing.T) {
	t.Run("should delete an entity", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus())
		entity, err := ecs.CreateEntity(11, 22)
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}

		// add sprite component
		if _, err = ecs.AddSpriteComponent(sprite.Sprite{
			E:       entity,
			I:       1,
			Graphic: sprite.SkeletonDeath1,
		}); err != nil {
			t.Errorf("AddSprite() error = %v", err)
		}

		// check if the sprite has been added
		if l := len(ecs.SpritesByEntity(entity)); l != 1 {
			t.Errorf("SpritesByEntity() sprites = %v", l)
		}

		if err = ecs.DeleteEntity(entity); err != nil {
			t.Errorf("DeleteEntity() error = %v", err)
		}

		// Check if the entity has been deleted
		_, err = ecs.Position(entity)
		if !errors.Is(err, component.ErrEntityNotFound) {
			t.Errorf("GetPosition() error = %v", err)
		}

		// Check if the sprite has been deleted
		if l := len(ecs.SpritesByEntity(entity)); l != 0 {
			t.Errorf("SpritesByEntity() sprites = %v", l)
		}
	})
}
