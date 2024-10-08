package component

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/dwethmar/vork/entity"
)

var (
	// ErrComponentNotFound is returned when a component is not found.
	ErrComponentNotFound = errors.New("component not found")
	// ErrEntityNotFound is returned when an entity is not found in the store.
	ErrEntityNotFound = errors.New("entity not found")
)

// Store holds components in memory and provides CRUD operations.
// It is generic over type C, which must implement the Component interface.
type Store[C Component] struct {
	mu              sync.RWMutex
	components      []*C
	entityIndex     map[entity.Entity][]*C // Maps Entity ID to Components
	nextID          uint32
	uniquePerEntity bool // Flag to enforce uniqueness per entity
}

// NewStore creates a new instance of Store for a specific component type.
// If uniquePerEntity is true, the store will enforce that only one component
// per entity can be added.
func NewStore[C Component](uniquePerEntity bool) *Store[C] {
	return &Store[C]{
		components:      []*C{},
		entityIndex:     make(map[entity.Entity][]*C),
		nextID:          1,
		uniquePerEntity: uniquePerEntity,
	}
}

// Add inserts a new component into the store.
// If the component ID is zero, it assigns a new unique ID.
// Enforces uniqueness per entity if the flag is set.
func (s *Store[C]) Add(c C) (uint32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entityID := c.Entity()

	if s.uniquePerEntity {
		if comps, exists := s.entityIndex[entityID]; exists && len(comps) > 0 {
			return 0, fmt.Errorf("component for entity ID %d already exists with component ID %d", entityID, (*comps[0]).ID())
		}
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
func (s *Store[C]) Get(id uint32) (C, error) {
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
func (s *Store[C]) Update(c C) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := s.searchComponentIndex(c.ID())
	if index < len(s.components) && (*s.components[index]).ID() == c.ID() {
		// Update the component in the components slice
		*s.components[index] = c
		return nil
	}
	return fmt.Errorf("component with ID %d not found", c.ID())
}

// Delete removes a component by its ID using binary search.
func (s *Store[C]) Delete(id uint32) error {
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
func (s *Store[C]) List() []C {
	s.mu.RLock()
	defer s.mu.RUnlock()

	components := make([]C, len(s.components))
	for i, compPtr := range s.components {
		components[i] = *compPtr
	}
	return components
}

// FirstByEntity retrieves the first component associated with an entity.
func (s *Store[C]) FirstByEntity(e entity.Entity) (C, error) {
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
func (s *Store[C]) ListByEntity(e entity.Entity) []C {
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
func (s *Store[C]) DeleteByEntity(e entity.Entity) error {
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
func (s *Store[C]) searchComponentIndex(id uint32) int {
	return sort.Search(len(s.components), func(i int) bool {
		return (*s.components[i]).ID() >= id
	})
}

// insertComponentSorted inserts a component into the sorted slice.
func (s *Store[C]) insertComponentSorted(c *C) {
	index := s.searchComponentIndex((*c).ID())
	// Insert the component at the correct position
	s.components = append(s.components, c) // Append to increase the slice size
	copy(s.components[index+1:], s.components[index:])
	s.components[index] = c
}
