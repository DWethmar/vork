package ecsys_test

import (
	"testing"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/hierarchy"
	"github.com/google/go-cmp/cmp"
)

func TestECS_AddPositionComponent(t *testing.T) {
	t.Run("should add a position component", func(t *testing.T) {
		h := hierarchy.New(0)
		ecs := ecsys.New(event.NewBus(), ecsys.NewStores(), h)
		e, err := ecs.CreateEntity(h.Root(), 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		// check if entity was added to hierarchy
		if diff := cmp.Diff(h.Children(0), []entity.Entity{e}); diff != "" {
			t.Errorf("Expected 1 child, got %d", len(diff))
		}
	})
}
