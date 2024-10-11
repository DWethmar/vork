package hierarchy

import (
	"errors"
	"fmt"
	"slices"

	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

var (
	// ErrEntityNotFound is returned when an entity is not found in the store.
	ErrEntityNotFound = errors.New("entity not found")
	// ErrCyclicRelationship is returned when a cyclic relationship is detected.
	ErrCyclicRelationship = errors.New("cannot add parent as a child of its descendant")
)

// Hierarchy build a tree of entities based on their parent-child relationships.
// It listens to position events to update the hierarchy when entities are created or deleted.
type Hierarchy struct {
	ecs           *ecsys.ECS
	eventBus      *event.Bus
	root          entity.Entity
	tree          map[entity.Entity][]entity.Entity
	parent        map[entity.Entity]entity.Entity
	subscriptions []int
}

func New(root entity.Entity, eventBus *event.Bus, ecs *ecsys.ECS) *Hierarchy {
	h := &Hierarchy{
		ecs:      ecs,
		eventBus: eventBus,
		root:     root,
		tree:     make(map[entity.Entity][]entity.Entity),
		parent:   make(map[entity.Entity]entity.Entity),
	}

	// add root to the hierarchy
	h.tree[root] = []entity.Entity{}

	h.subscriptions = append(h.subscriptions, eventBus.Subscribe(
		event.MatchAny(position.CreatedEventType, position.UpdatedEventType, position.DeletedEventType),
		h.eventHandler),
	)

	return h
}

// Close unsubscribes the hierarchy from the event bus.
func (h *Hierarchy) Close() error {
	for _, sub := range h.subscriptions {
		h.eventBus.Unsubscribe(sub)
	}
	return nil
}

// eventHandler handles position events.
// It updates the hierarchy when entities are created or deleted.
func (h *Hierarchy) eventHandler(e event.Event) error {
	switch t := e.(type) {
	case *position.CreatedEvent:
		return h.add(t.Position())
	case *position.UpdatedEvent:
		return h.update(t.Position())
	case *position.DeletedEvent:
		return h.remove(t.Position())
	}
	return nil
}

func (h *Hierarchy) add(p *position.Position) error {
	// Prevent cycles
	if err := h.detectCyclicRelationship(p.Entity(), p.Parent); err != nil {
		return err
	}

	// Update child's parent
	h.parent[p.Entity()] = p.Parent

	// Add child to parent's list
	children, ok := h.tree[p.Parent]
	if !ok {
		children = []entity.Entity{}
	}
	h.tree[p.Parent] = append(children, p.Entity())

	return nil
}

func (h *Hierarchy) update(p *position.Position) error {
	// Prevent cycles
	if err := h.detectCyclicRelationship(p.Entity(), p.Parent); err != nil {
		return err
	}

	// Remove child from current parent's list
	oldParent := h.parent[p.Entity()]
	children, ok := h.tree[oldParent]
	if !ok {
		return errors.New("current parent entity not found")
	}

	index := slices.IndexFunc(children, func(e entity.Entity) bool {
		return e == p.Entity()
	})

	if index == -1 {
		return errors.New("child entity not found")
	}

	// Remove child from old parent's list
	h.tree[oldParent] = append(children[:index], children[index+1:]...)

	// Update child's parent
	h.parent[p.Entity()] = p.Parent

	// Add child to new parent's list
	newChildren, ok := h.tree[p.Parent]
	if !ok {
		newChildren = []entity.Entity{}
	}
	h.tree[p.Parent] = append(newChildren, p.Entity())

	return nil
}

func (h *Hierarchy) remove(p *position.Position) error {
	// delete children
	for _, child := range h.tree[p.Entity()] {
		if err := h.ecs.DeleteEntity(child); err != nil {
			return fmt.Errorf("failed to delete entity %v: %w", child, err)
		}
	}

	// Get the list of siblings for the parent
	siblings, ok := h.tree[p.Parent]
	if !ok {
		return errors.New("parent entity not found")
	}

	// Find the index of the child in the parent's children list
	index := slices.IndexFunc(siblings, func(e entity.Entity) bool {
		return e == p.Entity()
	})

	if index == -1 {
		return errors.New("child entity not found in parent's list")
	}

	// Remove the child from the parent's children list
	h.tree[p.Parent] = append(siblings[:index], siblings[index+1:]...)

	// Remove the parent-child relationship
	delete(h.parent, p.Entity())

	// Clean up the tree map if the child has no children
	delete(h.tree, p.Entity())

	return nil
}

// Children returns the direct descendants of the parent entity.
func (h *Hierarchy) Children(e entity.Entity) ([]entity.Entity, error) {
	children, ok := h.tree[e]
	if !ok {
		return nil, ErrEntityNotFound
	}
	r := make([]entity.Entity, len(children))
	copy(r, children)
	return r, nil
}

func (h *Hierarchy) Root() entity.Entity {
	return h.root
}

// detectCyclicRelationship checks if adding the 'parent' to the 'entity' would create a cycle.
// This happens if the entity is already an ancestor of the parent.
func (h *Hierarchy) detectCyclicRelationship(entity, parent entity.Entity) error {
	current := parent

	// Traverse upwards through the hierarchy to check if 'entity' is an ancestor of 'parent'
	for current != h.root {
		if current == entity {
			return ErrCyclicRelationship // A cycle is detected
		}

		// Move up to the parent of the current entity
		parentEntity, ok := h.parent[current]
		if !ok {
			break // No more parents, we've reached the root
		}
		current = parentEntity
	}

	return nil // No cycle detected
}

func (h *Hierarchy) GetAbsolutePosition(e entity.Entity) (int, int, error) {
	x, y := 0, 0
	current := e
	root := h.Root()
	for {
		// skip root entity
		if current == root {
			break
		}
		// Get the position component
		pos, err := h.ecs.GetPosition(current)
		if err != nil {
			return 0, 0, fmt.Errorf("position not found for entity %v: %w", current, err)
		}

		// Accumulate positions
		x += pos.X
		y += pos.Y

		// Move to parent
		parent, ok := h.parent[current]
		if !ok {
			break // Reached the root
		}
		current = parent
	}

	return x, y, nil
}
