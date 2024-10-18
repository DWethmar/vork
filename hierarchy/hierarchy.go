package hierarchy

import (
	"errors"
	"fmt"

	"github.com/dwethmar/vork/entity"
)

var (
	// ErrEntityNotFound is returned when an entity is not found in the store.
	ErrEntityNotFound = errors.New("entity not found")
	// ErrCyclicRelationship is returned when a cyclic relationship is detected.
	ErrCyclicRelationship = errors.New("cannot add parent as a child of its descendant")
)

type Hierarchy struct {
	root     entity.Entity
	parents  map[entity.Entity]entity.Entity   // child -> parent
	children map[entity.Entity][]entity.Entity // parent -> []children
}

func New(root entity.Entity) *Hierarchy {
	return &Hierarchy{
		root:     root,
		parents:  make(map[entity.Entity]entity.Entity),
		children: make(map[entity.Entity][]entity.Entity),
	}
}

type EntityPair struct {
	Parent entity.Entity
	Child  entity.Entity
}

// Clear removes all entities from the hierarchy.
func (h *Hierarchy) clear() {
	h.parents = make(map[entity.Entity]entity.Entity)
	h.children = make(map[entity.Entity][]entity.Entity)
}

func (h *Hierarchy) Build(pairs []EntityPair) error {
	h.clear()
	// Temporary tree to store relationships before hierarchy is built
	tree := make(map[entity.Entity][]entity.Entity)

	// Build the tree from the pairs
	for _, pair := range pairs {
		if _, exists := tree[pair.Parent]; !exists {
			tree[pair.Parent] = make([]entity.Entity, 0)
		}
		tree[pair.Parent] = append(tree[pair.Parent], pair.Child)
	}

	return h.addToParent(h.root, tree)
}

func (h *Hierarchy) addToParent(parent entity.Entity, tree map[entity.Entity][]entity.Entity) error {
	// Retrieve children of the current parent
	children, exists := tree[parent]
	if !exists {
		// If the parent has no children, nothing to do
		return nil
	}

	// Add each child to the hierarchy under the current parent
	for _, child := range children {
		// Add the child under the current parent
		if err := h.Add(parent, child); err != nil {
			return fmt.Errorf("error adding child %v to parent %v: %w", child, parent, err)
		}

		// Recursively add the children of this child
		if err := h.addToParent(child, tree); err != nil {
			return err
		}
	}

	return nil
}

func (h *Hierarchy) entityExists(e entity.Entity) bool {
	if e == h.root {
		return true
	}
	if _, exists := h.parents[e]; exists {
		return true
	}
	if _, exists := h.children[e]; exists {
		return true
	}
	return false
}

func (h *Hierarchy) Root() entity.Entity {
	return h.root
}

func (h *Hierarchy) Add(parent entity.Entity, child entity.Entity) error {
	// Check if the parent exists in the hierarchy
	if !h.entityExists(parent) {
		return fmt.Errorf("parent %v does not exist in the hierarchy", parent)
	}
	// Check if the child already has a parent
	if p, exists := h.parents[child]; exists {
		return fmt.Errorf("child %v already has parent %v", child, p)
	}
	// Check for cycles
	if h.hasPath(child, parent) {
		return ErrCyclicRelationship
	}
	// Add parent relationship
	h.parents[child] = parent
	h.children[parent] = append(h.children[parent], child)
	return nil
}

func (h *Hierarchy) Update(parent entity.Entity, child entity.Entity) error {
	// check if update is necessary
	if parent == h.parents[child] {
		return nil
	}
	// Check if the new parent exists
	if !h.entityExists(parent) {
		return fmt.Errorf("parent %v does not exist in the hierarchy", parent)
	}
	// Remove old parent relationship
	oldParent, exists := h.parents[child]
	if exists {
		h.removeChildFromParent(oldParent, child)
	}
	// Check for cycles
	if h.hasPath(child, parent) {
		return ErrCyclicRelationship
	}
	// Update parent
	h.parents[child] = parent
	h.children[parent] = append(h.children[parent], child)
	return nil
}

// Delete removes an entity and all its descendants from the hierarchy.
func (h *Hierarchy) Delete(child entity.Entity) []entity.Entity {
	// Check if the entity exists
	if !h.entityExists(child) {
		return nil
	}
	// Add the entity itself
	deletedEntities := []entity.Entity{child}
	// Collect all descendants recursively
	deletedEntities = append(deletedEntities, h.collectDescendants(child)...)
	// Remove from parents and children maps
	for _, e := range deletedEntities {
		// Remove from parent's children list
		parent, hasParent := h.parents[e]
		if hasParent {
			h.removeChildFromParent(parent, e)
			delete(h.parents, e)
		}
		// Remove from children map
		delete(h.children, e)
	}
	return deletedEntities
}

func (h *Hierarchy) Parent(child entity.Entity) (entity.Entity, error) {
	if child == h.root {
		return 0, fmt.Errorf("root %v has no parent", h.root)
	}
	parent, exists := h.parents[child]
	if !exists {
		return 0, fmt.Errorf("no parent found for child %v: %w", child, ErrEntityNotFound)
	}
	return parent, nil
}

func (h *Hierarchy) Children(parent entity.Entity) []entity.Entity {
	return h.children[parent]
}

// Helper method to check if there's a path from 'from' to 'to' (used for cycle detection).
func (h *Hierarchy) hasPath(from, to entity.Entity) bool {
	if from == to {
		return true
	}
	children := h.children[from]
	for _, child := range children {
		if h.hasPath(child, to) {
			return true
		}
	}
	return false
}

// Helper method to collect all descendants of an entity (used for recursive deletion).
func (h *Hierarchy) collectDescendants(e entity.Entity) []entity.Entity {
	children := h.children[e]
	if len(children) == 0 {
		return nil
	}
	descendants := make([]entity.Entity, 0, len(children))
	for _, child := range children {
		descendants = append(descendants, child)
		descendants = append(descendants, h.collectDescendants(child)...)
	}
	return descendants
}

// Helper method to remove a child from its parent's children list.
func (h *Hierarchy) removeChildFromParent(parent, child entity.Entity) {
	children := h.children[parent]
	for i, c := range children {
		if c == child {
			h.children[parent] = append(children[:i], children[i+1:]...)
			break
		}
	}
}
