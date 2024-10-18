package ecsys_test

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/hierarchy"
	"github.com/dwethmar/vork/point"
)

func TestECS_DeletePosition(t *testing.T) {
	t.Run("should delete a position component", func(t *testing.T) {
		h := hierarchy.New(0)
		ecs := ecsys.New(event.NewBus(), ecsys.NewStores(), h)

		child1, err := ecs.CreateEntity(h.Root(), point.Zero())
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		child2, err := ecs.CreateEntity(child1, point.Zero())
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		// get position of child1
		pos, err := ecs.GetPosition(child1)
		if err != nil {
			t.Errorf("Error getting position: %s", err)
		}

		// delete position
		if err = ecs.DeletePosition(pos); err != nil {
			t.Errorf("Error deleting position: %s", err)
		}

		// check if entity was added to hierarchy
		if c := h.Children(0); len(c) != 0 {
			t.Errorf("Expected 0 children, got children %v", c)
		}

		_, err = ecs.GetPosition(child2)
		if !errors.Is(err, ecsys.ErrEntityNotFound) {
			t.Errorf("Expected error, got %s", err)
		}
	})
}
