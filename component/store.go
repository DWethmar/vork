package component

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/dwethmar/vork/entity"
)

// Predefined errors.
var (
	// ErrComponentNotFound is returned when a component is not found.
	ErrComponentNotFound = errors.New("component not found")
	// ErrEntityNotFound is returned when an entity is not found in the store.
	ErrEntityNotFound = errors.New("entity not found")
	// ErrUniqueComponentViolation is returned when a component is unique per entity.
	ErrUniqueComponentViolation = errors.New("only one component per entity allowed")
)

// Store holds components in memory and provides CRUD operations.
// It is generic over type C, which must implement the Component interface.
type Store[C Component] struct {
	mu              sync.RWMutex
	components      map[uint]C                          // components by ID
	entityIndex     map[entity.Entity]map[uint]struct{} // maps an entity to a set of component IDs
	nextID          uint
	uniquePerEntity bool
	onCreate        func(C) error
	onUpdate        func(C) error
	onDelete        func(C) error
}

// NewStore creates a new instance of Store for a specific component type.
// If uniquePerEntity is true, the store will enforce that only one component
// per entity can be added.
func NewStore[C Component](
	uniquePerEntity bool,
	onAdd func(C) error,
	onUpdate func(C) error,
	onDelete func(C) error,
) *Store[C] {
	return &Store[C]{
		components:      make(map[uint]C),
		entityIndex:     make(map[entity.Entity]map[uint]struct{}),
		nextID:          1,
		uniquePerEntity: uniquePerEntity,
		onCreate:        onAdd,
		onUpdate:        onUpdate,
		onDelete:        onDelete,
	}
}

// Add inserts a new component into the store.
// If the component ID is zero, a new unique ID is assigned.
// If uniquePerEntity is true, only one component per entity is allowed.
func (s *Store[C]) Add(c C) (uint, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ent := c.Entity()
	if s.uniquePerEntity {
		if comps, exists := s.entityIndex[ent]; exists && len(comps) > 0 {
			return 0, ErrUniqueComponentViolation
		}
	}

	// Assign a new ID if needed.
	if c.ID() == 0 {
		c.SetID(s.nextID)
		s.nextID++
	} else if _, exists := s.components[c.ID()]; exists {
		return 0, fmt.Errorf("component with ID %d already exists", c.ID())
	}

	// Save the component.
	s.components[c.ID()] = c

	// Update the entity index.
	if _, exists := s.entityIndex[ent]; !exists {
		s.entityIndex[ent] = make(map[uint]struct{})
	}
	s.entityIndex[ent][c.ID()] = struct{}{}

	// Call the onCreate handler if provided.
	if s.onCreate != nil {
		if err := s.onCreate(c); err != nil {
			return c.ID(), fmt.Errorf("on create handler failed: %w", err)
		}
	}

	return c.ID(), nil
}

// Get retrieves a component by its ID.
func (s *Store[C]) Get(id uint) (C, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	c, exists := s.components[id]
	if !exists {
		var zero C
		return zero, ErrComponentNotFound
	}
	return c, nil
}

// Update modifies an existing component in the store.
// If the component's associated entity changes, the entity index is updated.
// The unique-per-entity constraint is enforced as needed.
func (s *Store[C]) Update(c C) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	existing, exists := s.components[c.ID()]
	if !exists {
		return ErrComponentNotFound
	}

	// If the entity has changed, update the entity index.
	if existing.Entity() != c.Entity() {
		oldEnt := existing.Entity()
		if ids, ok := s.entityIndex[oldEnt]; ok {
			delete(ids, c.ID())
			if len(ids) == 0 {
				delete(s.entityIndex, oldEnt)
			}
		}
		// Enforce unique constraint on the new entity.
		if s.uniquePerEntity {
			if comps, exists := s.entityIndex[c.Entity()]; exists && len(comps) > 0 {
				return ErrUniqueComponentViolation
			}
		}
		if _, exists := s.entityIndex[c.Entity()]; !exists {
			s.entityIndex[c.Entity()] = make(map[uint]struct{})
		}
		s.entityIndex[c.Entity()][c.ID()] = struct{}{}
	}

	s.components[c.ID()] = c

	// Call the onUpdate handler if provided.
	if s.onUpdate != nil {
		if err := s.onUpdate(c); err != nil {
			return fmt.Errorf("on update handler failed: %w", err)
		}
	}

	return nil
}

// List returns all components in the store, sorted by their ID.
func (s *Store[C]) List() []C {
	s.mu.RLock()
	defer s.mu.RUnlock()

	components := make([]C, 0, len(s.components))
	for _, c := range s.components {
		components = append(components, c)
	}
	sort.Slice(components, func(i, j int) bool {
		return components[i].ID() < components[j].ID()
	})
	return components
}

// First retrieves the first component associated with an entity.
func (s *Store[C]) First(e entity.Entity) (C, error) {
	comps := s.All(e)
	if len(comps) == 0 {
		var zero C
		return zero, ErrEntityNotFound
	}
	return comps[0], nil
}

// ListByEntity retrieves all components associated with an entity, sorted by their ID.
func (s *Store[C]) All(e entity.Entity) []C {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ids, exists := s.entityIndex[e]
	if !exists {
		return nil
	}

	components := make([]C, 0, len(ids))
	for id := range ids {
		if comp, exists := s.components[id]; exists {
			components = append(components, comp)
		}
	}
	sort.Slice(components, func(i, j int) bool {
		return components[i].ID() < components[j].ID()
	})
	return components
}

// Delete removes a component by its ID.
func (s *Store[C]) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, exists := s.components[id]
	if !exists {
		return ErrComponentNotFound
	}

	delete(s.components, id)

	ent := c.Entity()
	if ids, ok := s.entityIndex[ent]; ok {
		delete(ids, id)
		if len(ids) == 0 {
			delete(s.entityIndex, ent)
		}
	}

	// Call the onDelete handler if provided.
	if s.onDelete != nil {
		if err := s.onDelete(c); err != nil {
			return fmt.Errorf("on delete handler failed: %w", err)
		}
	}

	return nil
}

// DeleteAll removes all components associated with an entity.
func (s *Store[C]) DeleteAll(e entity.Entity) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ids, exists := s.entityIndex[e]
	if !exists {
		return ErrEntityNotFound
	}

	for id := range ids {
		c := s.components[id]
		delete(s.components, id)
		if s.onDelete != nil {
			if err := s.onDelete(c); err != nil {
				return fmt.Errorf("on delete handler failed: %w", err)
			}
		}
	}
	delete(s.entityIndex, e)
	return nil
}
