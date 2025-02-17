package ecsys_test

import (
	"testing"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/point"
	"github.com/google/go-cmp/cmp"
)

func TestECS_UpdatePositionComponent(t *testing.T) {
	t.Run("should update a position component", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus(), ecsys.NewStores())

		parent, err := ecs.CreateEntity(ecs.Root(), point.Zero())
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		child1, err := ecs.CreateEntity(parent, point.Zero())
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		pos, err := ecs.GetPosition(child1)
		if err != nil {
			t.Errorf("Error getting position: %s", err)
		}

		if pos.Parent != parent {
			t.Errorf("Expected parent to be %d, got %d", parent, pos.Parent)
		}

		pos.Parent = ecs.Root()

		// move child 1 to root
		if err = ecs.UpdatePositionComponent(pos); err != nil {
			t.Errorf("Error updating position: %s", err)
		}

		// check if entity was added to hierarchy
		if diff := cmp.Diff(ecs.Children(0), []entity.Entity{parent, child1}); diff != "" {
			t.Errorf("Expected 1 child, got %s", diff)
		}
	})
}
