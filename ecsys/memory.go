package ecsys

import (
	"fmt"
	"sort"
	"sync"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

// MemStore holds components in memory and provides CRUD operations.
// It is generic over type C, which must implement the Component interface.
type MemStore[C component.Component] struct {
	mu              sync.RWMutex
	components      []*C
	entityIndex     map[entity.Entity][]*C // Maps Entity ID to Components
	nextID          uint
	uniquePerEntity bool // Flag to enforce uniqueness per entity
}

// NewMemoryStore creates a new instance of Store for a specific component type.
// If uniquePerEntity is true, the store will enforce that only one component
// per entity can be added.
func NewMemStore[C component.Component](uniquePerEntity bool) *MemStore[C] {
	return &MemStore[C]{
		components:      []*C{},
		entityIndex:     make(map[entity.Entity][]*C),
		nextID:          1,
		uniquePerEntity: uniquePerEntity,
	}
}

// applyUniqueConstraint checks if a component violates the uniqueness constraint.
func (s *MemStore[C]) applyUniqueConstraint(e entity.Entity, id uint) error {
	if !s.uniquePerEntity {
		return nil
	}

	comps := s.entityIndex[e]
	if len(comps) == 0 {
		// No existing components; safe to add.
		return nil
	}

	if len(comps) == 1 && (*comps[0]).ID() == id {
		// Existing component has the same ID; it's an update.
		return nil
	}

	return ErrUniqueComponentViolation
}

// Add inserts a new component into the store.
// If the component ID is zero, it assigns a new unique ID.
// Enforces uniqueness per entity if the flag is set.
func (s *MemStore[C]) Add(c C) (uint, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entityID := c.Entity()

	if err := s.applyUniqueConstraint(entityID, c.ID()); err != nil {
		return 0, err
	}

	if c.ID() == 0 {
		c.SetID(s.nextID)
		s.nextID++
	} else {
		// Check if a component with this ID already exists using binary search
		index := s.searchComponentIndex(c.ID())
		if index < len(s.components) && (*s.components[index]).ID() == c.ID() {
			return 0, fmt.Errorf("component with ID %d already exists", c.ID())
		}
	}

	// Insert the component into the sorted slice
	s.insertComponentSorted(&c)

	// Add component to entityIndex
	s.entityIndex[entityID] = append(s.entityIndex[entityID], &c)

	return c.ID(), nil
}

// Get retrieves a component by its ID using binary search.
func (s *MemStore[C]) Get(id uint) (C, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	index := s.searchComponentIndex(id)
	if index < len(s.components) && (*s.components[index]).ID() == id {
		return *s.components[index], nil
	}
	var zero C
	return zero, ErrComponentNotFound
}

// Update modifies an existing component in the store using binary search.
func (s *MemStore[C]) Update(c C) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.applyUniqueConstraint(c.Entity(), c.ID()); err != nil {
		return err
	}

	index := s.searchComponentIndex(c.ID())
	if index < len(s.components) && (*s.components[index]).ID() == c.ID() {
		// Update the component in the components slice
		*s.components[index] = c
		return nil
	}
	return ErrComponentNotFound
}

// Delete removes a component by its ID using binary search.
func (s *MemStore[C]) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	index := s.searchComponentIndex(id)
	if index >= len(s.components) || (*s.components[index]).ID() != id {
		return ErrComponentNotFound
	}
	c := s.components[index]
	entityID := (*c).Entity()
	// Remove the component from the components slice
	s.components = append(s.components[:index], s.components[index+1:]...)
	// Remove component from entityIndex
	comps := s.entityIndex[entityID]
	for i, comp := range comps {
		if (*comp).ID() == id {
			// Remove the component from the slice
			s.entityIndex[entityID] = append(comps[:i], comps[i+1:]...)
			break
		}
	}
	// If the slice is empty after removal, delete the key from the map
	if len(s.entityIndex[entityID]) == 0 {
		delete(s.entityIndex, entityID)
	}
	return nil
}

// List returns all components in the store.
func (s *MemStore[C]) List() []C {
	s.mu.RLock()
	defer s.mu.RUnlock()
	components := make([]C, len(s.components))
	for i, compPtr := range s.components {
		components[i] = *compPtr
	}
	return components
}

// First retrieves the first component associated with an entity.
func (s *MemStore[C]) First(e entity.Entity) (C, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comps, exists := s.entityIndex[e]
	if !exists || len(comps) == 0 {
		var zero C
		return zero, ErrEntityNotFound
	}

	return *comps[0], nil
}

// ListByEntity retrieves all components associated with an entity.
func (s *MemStore[C]) ListByEntity(e entity.Entity) []C {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comps, exists := s.entityIndex[e]
	if !exists || len(comps) == 0 {
		return nil
	}

	components := make([]C, len(comps))
	for i, compPtr := range comps {
		components[i] = *compPtr
	}
	return components
}

// DeleteByEntity removes all components associated with an entity.
func (s *MemStore[C]) DeleteByEntity(e entity.Entity) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	comps, exists := s.entityIndex[e]
	if !exists {
		return ErrEntityNotFound // No components to delete for this entity
	}

	// Delete each component associated with the entity
	for _, compPtr := range comps {
		id := (*compPtr).ID()
		index := s.searchComponentIndex(id)
		if index < len(s.components) && (*s.components[index]).ID() == id {
			// Remove the component from the slice
			s.components = append(s.components[:index], s.components[index+1:]...)
		}
	}

	// Remove the entity from the index
	delete(s.entityIndex, e)
	return nil
}

// searchComponentIndex performs a binary search to find the index of a component with the given ID.
func (s *MemStore[C]) searchComponentIndex(id uint) int {
	return sort.Search(len(s.components), func(i int) bool {
		return (*s.components[i]).ID() >= id
	})
}

// insertComponentSorted inserts a component into the sorted slice.
func (s *MemStore[C]) insertComponentSorted(c *C) {
	index := s.searchComponentIndex((*c).ID())
	// Insert the component at the correct position
	s.components = append(s.components, c) // Append to increase the slice size
	copy(s.components[index+1:], s.components[index:])
	s.components[index] = c
}
