package hierarchy_test

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/hierarchy"
	"github.com/dwethmar/vork/point"
	"github.com/google/go-cmp/cmp"
)

// assertChildren checks if the children of a parent are equal to the expected children.
func assertChildren(t *testing.T, h *hierarchy.Hierarchy, parent entity.Entity, expect []entity.Entity) {
	t.Helper()
	children := h.Children(parent)
	if diff := cmp.Diff(children, expect); diff != "" {
		t.Errorf("expected children of %v to be %v, got %v", parent, expect, children)
	}
}

//nolint:gocognit
func TestNew(t *testing.T) {
	t.Run("New should create a new hierarchy", func(t *testing.T) {
		root := entity.Entity(0)
		h := hierarchy.New(root)
		if h == nil {
			t.Error("Hierarchy should not be nil")
		}
	})

	t.Run("should return an error if a cyclic relationship is detected", func(t *testing.T) {
		root := entity.Entity(0)
		h := hierarchy.New(root)

		child1 := entity.Entity(1)
		child2 := entity.Entity(2)
		child3 := entity.Entity(3)

		// add root -> child1 -> child2 -> child3
		if err := h.Add(root, child1); err != nil {
			t.Errorf("Error adding child1: %s", err)
		}

		if err := h.Add(child1, child2); err != nil {
			t.Errorf("Error adding child2: %s", err)
		}

		if err := h.Add(child2, child3); err != nil {
			t.Errorf("Error adding child3: %s", err)
		}

		// we now have root -> child1 -> child2 -> child3
		// try to add child3 as a parent of child1
		if err := h.Update(child3, child1); !errors.Is(err, hierarchy.ErrCyclicRelationship) {
			t.Error("Error should be ErrCyclicRelationship")
		}
	})

	t.Run("deleting a parent should remove the children from the hierarchy", func(t *testing.T) {
		root := entity.Entity(0)
		h := hierarchy.New(root)

		child1 := entity.Entity(1)
		child2 := entity.Entity(2)
		child3 := entity.Entity(3)

		if err := h.Add(root, child1); err != nil {
			t.Errorf("Error adding child1: %s", err)
		}

		if err := h.Add(child1, child2); err != nil {
			t.Errorf("Error adding child2: %s", err)
		}

		if err := h.Add(child1, child3); err != nil {
			t.Errorf("Error adding child3: %s", err)
		}

		// tree should be child1 -> child2, child3

		// delete child1, parent of child2 and child3
		deletedChildren := h.Delete(child1)

		// check if the children are returned
		expectedChildren := []entity.Entity{child1, child2, child3}
		if diff := cmp.Diff(deletedChildren, expectedChildren); diff != "" {
			t.Errorf("Children should be equal: %s", diff)
		}

		// check if the children are removed from the hierarchy
		rootChildren := h.Children(root)
		if len(rootChildren) != 0 {
			t.Errorf("Root should have no children but has %v", rootChildren)
		}

		// check if the children are removed from the ecs
		if _, err := h.Parent(child2); !errors.Is(err, hierarchy.ErrEntityNotFound) {
			t.Errorf("Child2 should not have a parent: %s", err)
		}

		if _, err := h.Parent(child3); !errors.Is(err, hierarchy.ErrEntityNotFound) {
			t.Errorf("Child3 should not have a parent: %s", err)
		}

		// check if the parent is removed from the hierarchy
		if _, err := h.Parent(child1); !errors.Is(err, hierarchy.ErrEntityNotFound) {
			t.Errorf("Child1 should not have a parent: %s", err)
		}
	})

	t.Run("updating a parent should update the children in the hierarchy", func(t *testing.T) {
		root := entity.Entity(0)
		h := hierarchy.New(root)

		child1 := entity.Entity(1)
		child2 := entity.Entity(2)
		child3 := entity.Entity(3)
		child4 := entity.Entity(4)

		if err := h.Add(root, child1); err != nil {
			t.Errorf("Error adding child1: %s", err)
		}

		if err := h.Add(child1, child2); err != nil {
			t.Errorf("Error adding child2: %s", err)
		}

		if err := h.Add(child2, child3); err != nil {
			t.Errorf("Error adding child3: %s", err)
		}

		if err := h.Add(child2, child4); err != nil {
			t.Errorf("Error adding child4: %s", err)
		}

		// tree should be root -> child1 -> child2 -> child3, child4

		assertChildren(t, h, root, []entity.Entity{child1})
		assertChildren(t, h, child1, []entity.Entity{child2})
		assertChildren(t, h, child2, []entity.Entity{child3, child4})

		// update child3's parent to child1
		if err := h.Update(child1, child4); err != nil {
			t.Errorf("Error updating child4: %s", err)
		}

		// tree should be root -> child1 -> child2, child4 -> child3
		assertChildren(t, h, child2, []entity.Entity{child3})
		assertChildren(t, h, child1, []entity.Entity{child2, child4})
	})
}

func TestHierarchy_Built(t *testing.T) {
	root := entity.Entity(0)
	p := []hierarchy.EntityPair{
		{
			Parent: root,
			Child:  entity.Entity(1),
		},
		{
			Parent: root,
			Child:  entity.Entity(2),
		},
		{
			Parent: entity.Entity(1),
			Child:  entity.Entity(3),
		},
		{
			Parent: entity.Entity(2),
			Child:  entity.Entity(4),
		},
	}

	h := hierarchy.New(root)
	if err := h.Build(p); err != nil {
		t.Error(err)
	}

	assertChildren(t, h, root, []entity.Entity{
		entity.Entity(1),
		entity.Entity(2),
	})

	assertChildren(t, h, entity.Entity(1), []entity.Entity{
		entity.Entity(3),
	})

	assertChildren(t, h, entity.Entity(2), []entity.Entity{
		entity.Entity(4),
	})
}

func TestHierarchy_Children(t *testing.T) {
	t.Run("should add a child to the hierarchy", func(t *testing.T) {
		eventBus := event.NewBus()
		root := entity.Entity(0)
		ecs := ecsys.New(eventBus, ecsys.NewStores())

		expect := []entity.Entity{}
		for range 10 {
			child, err := ecs.CreateEntity(root, point.New(0, 0))
			if err != nil {
				t.Error("Error creating entity")
			}
			expect = append(expect, child)
		}

		// check if the child was added
		children := ecs.Children(root)
		if cmp.Diff(children, expect) != "" {
			t.Error("Children should be equal")
		}
	})

	t.Run("should return an empty list if the entity has no children", func(t *testing.T) {
		root := entity.Entity(0)
		h := hierarchy.New(root)
		children := h.Children(root)
		if len(children) != 0 {
			t.Error("Children should be empty")
		}
	})
}

func TestHierarchy_Root(t *testing.T) {
	t.Run("Root should return the root entity", func(t *testing.T) {
		root := entity.Entity(0)
		h := hierarchy.New(root)
		if h.Root() != root {
			t.Error("Root should be 0")
		}
	})
}
