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
	mu          sync.RWMutex
	components  map[uint]C                          // components by ID
	entityIndex map[entity.Entity]map[uint]struct{} // maps an entity to a set of component IDs
}

// NewStore creates a new instance of Store for a specific component type.
// If uniquePerEntity is true, the store will enforce that only one component
// per entity can be added.
func NewStore[C Component]() *Store[C] {
	return &Store[C]{
		components:  make(map[uint]C),
		entityIndex: make(map[entity.Entity]map[uint]struct{}),
	}
}

// Add inserts a new component into the store.
// If the component ID is zero, a new unique ID is assigned.
// If uniquePerEntity is true, only one component per entity is allowed.
func (s *Store[C]) Add(c C) (uint, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	e := c.Entity()
	// Assign a new ID if needed.
	if c.ID() == 0 {
		return 0, fmt.Errorf("component ID must be set")
	} else if _, exists := s.components[c.ID()]; exists {
		return 0, fmt.Errorf("component with ID %d already exists", c.ID())
	}

	// Save the component.
	s.components[c.ID()] = c

	// Update the entity index.
	if _, exists := s.entityIndex[e]; !exists {
		s.entityIndex[e] = make(map[uint]struct{})
	}
	s.entityIndex[e][c.ID()] = struct{}{}

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
		if _, exists := s.entityIndex[c.Entity()]; !exists {
			s.entityIndex[c.Entity()] = make(map[uint]struct{})
		}
		s.entityIndex[c.Entity()][c.ID()] = struct{}{}
	}

	s.components[c.ID()] = c

	return nil
}

// List returns all components in the store, sorted by their ID.
func (s *Store[C]) All() []C {
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
	comps := s.List(e)
	if len(comps) == 0 {
		var zero C
		return zero, ErrEntityNotFound
	}
	return comps[0], nil
}

// ListByEntity retrieves all components associated with an entity, sorted by their ID.
func (s *Store[C]) List(e entity.Entity) []C {
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
		delete(s.components, id)
	}
	delete(s.entityIndex, e)
	return nil
}
