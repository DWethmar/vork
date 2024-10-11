package hierarchy_test

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/component/store"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/hierarchy"
	"github.com/google/go-cmp/cmp"
)

// assertChildren checks if the children of a parent are equal to the expected children
func assertChildren(t *testing.T, h *hierarchy.Hierarchy, parent entity.Entity, expect []entity.Entity) {
	t.Helper()
	children, err := h.Children(parent)
	if err != nil {
		t.Errorf("Error getting children: %s", err)
	}

	if diff := cmp.Diff(children, expect); diff != "" {
		t.Errorf("Children should be equal: %s", diff)
	}
}

//nolint:gocognit
func TestNew(t *testing.T) {
	t.Run("New should create a new hierarchy", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)
		if h == nil {
			t.Error("Hierarchy should not be nil")
		}
	})

	t.Run("should return an error if a cyclic relationship is detected", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		hierarchy.New(root, eventBus, ecs)

		child1, err := ecs.CreateEntity(root, 0, 0)
		if err != nil {
			t.Error("Error creating entity")
		}

		child2, err := ecs.CreateEntity(child1, 0, 0)
		if err != nil {
			t.Error("Error creating entity")
		}

		child3, err := ecs.CreateEntity(child2, 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		// we now have root -> child1 -> child2 -> child3
		// try to add child3 as a parent of child1
		pos, err := ecs.GetPosition(child1)
		if err != nil {
			t.Error("Error getting position")
		}

		// this will create a cyclic relationship
		// child3 -> child2 -> child1 -> child3
		pos.Parent = child3

		if err = ecs.UpdatePositionComponent(pos); !errors.Is(err, hierarchy.ErrCyclicRelationship) {
			t.Error("Error should be ErrCyclicRelationship")
		}
	})

	t.Run("deleting a parent should remove the children from the hierarchy and ecs", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)

		parent, err := ecs.CreateEntity(root, 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		child1, err := ecs.CreateEntity(parent, 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		child2, err := ecs.CreateEntity(parent, 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		// tree should be root -> parent -> child1, child2

		// delete the parent
		if err = ecs.DeleteEntity(parent); err != nil {
			t.Errorf("Error deleting entity: %s", err)
		}

		// check if the children are removed from the hierarchy
		assertChildren(t, h, root, []entity.Entity{})

		// check if the children are removed from the ecs
		if _, err = ecs.GetPosition(child1); !errors.Is(err, store.ErrEntityNotFound) {
			t.Errorf("Error should not be nil: %s", err)
		}

		if _, err = ecs.GetPosition(child2); !errors.Is(err, store.ErrEntityNotFound) {
			t.Errorf("Error should not be nil: %s", err)
		}

		// check if the parent is removed from ecs
		if _, err = ecs.GetPosition(parent); !errors.Is(err, store.ErrEntityNotFound) {
			t.Errorf("Error should not be nil: %s", err)
		}
	})

	t.Run("updating a parent should update the children in the hierarchy", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)

		child1, err := ecs.CreateEntity(root, 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		child2, err := ecs.CreateEntity(child1, 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		child3, err := ecs.CreateEntity(child2, 0, 0)
		if err != nil {
			t.Errorf("Error creating entity: %s", err)
		}

		// tree should be root -> child1 -> child2 -> child3

		assertChildren(t, h, root, []entity.Entity{child1})
		assertChildren(t, h, child1, []entity.Entity{child2})
		assertChildren(t, h, child2, []entity.Entity{child3})

		// update child3's parent to child1
		pos, err := ecs.GetPosition(child3)
		if err != nil {
			t.Errorf("Error getting position: %s", err)
		}

		pos.Parent = child1
		if err = ecs.UpdatePositionComponent(pos); err != nil {
			t.Errorf("Error updating position: %s", err)
		}

		// tree should be root -> child1 -> child2, child3
		assertChildren(t, h, child1, []entity.Entity{child2, child3})
		// child2 should have no children
		assertChildren(t, h, child2, []entity.Entity{})
	})
}

func TestHierarchy_Close(t *testing.T) {
	t.Run("Close should close the hierarchy", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)
		if err := h.Close(); err != nil {
			t.Error("Hierarchy should not return an error")
		}

		// check if the subscriptions are removed
		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 0 {
			t.Error("Subscriptions should be empty")
		}
	})
}

func TestHierarchy_Children(t *testing.T) {
	t.Run("should add a child to the hierarchy", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)

		expect := []entity.Entity{}
		for range 10 {
			child, err := ecs.CreateEntity(root, 0, 0)
			if err != nil {
				t.Error("Error creating entity")
			}
			expect = append(expect, child)
		}

		// check if the child was added
		children, err := h.Children(root)
		if err != nil {
			t.Error("Error getting children")
		}

		if cmp.Diff(children, expect) != "" {
			t.Error("Children should be equal")
		}
	})

	t.Run("should return an error if the entity does not exist", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)

		_, err := h.Children(100)
		if err == nil {
			t.Error("Error should not be nil")
		}
	})

	t.Run("should return an empty list if the entity has no children", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)

		children, err := h.Children(root)
		if err != nil {
			t.Error("Error getting children")
		}

		if len(children) != 0 {
			t.Error("Children should be empty")
		}
	})
}

func TestHierarchy_Root(t *testing.T) {
	t.Run("Root should return the root entity", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		root := ecs.CreateEmptyEntity()
		h := hierarchy.New(root, eventBus, ecs)
		if h.Root() != root {
			t.Error("Root should be 0")
		}
	})
}

func TestHierarchy_GetAbsolutePosition(t *testing.T) {
	t.Run("GetAbsolutePosition should return the absolute position of an entity", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		h := hierarchy.New(ecs.CreateEmptyEntity(), eventBus, ecs)

		current := h.Root()
		for i := range 10 {
			child, err := ecs.CreateEntity(current, i+1, i+1)
			if err != nil {
				t.Error("Error creating entity")
			}
			current = child
		}

		x, y, err := h.GetAbsolutePosition(current)
		if err != nil {
			t.Errorf("Error getting absolute position: %s", err)
		}

		// current is 10 levels deep, so x and y should 1 + 2 + 4 ... + 10
		expectedX := 0
		expectedY := 0

		for i := range 10 {
			expectedX += i + 1
			expectedY += i + 1
		}

		if x != expectedX || y != expectedY {
			t.Errorf("Expected x: %d, y: %d, got x: %d, y: %d", expectedX, expectedY, x, y)
		}
	})
}
