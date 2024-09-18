package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dwethmar/vork/component"
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
type Store[C component.Component] struct {
	mu              sync.RWMutex
	components      map[uint32]C               // Maps Component ID to Component
	entityIndex     map[entity.Entity][]uint32 // Maps Entity ID to Component IDs
	nextID          uint32
	uniquePerEntity bool // Flag to enforce uniqueness per entity
}

// New creates a new instance of Store for a specific component type.
// If uniquePerEntity is true, the store will enforce that only one component
// per entity can be added.
func New[C component.Component](uniquePerEntity bool) *Store[C] {
	return &Store[C]{
		components:      make(map[uint32]C),
		entityIndex:     make(map[entity.Entity][]uint32),
		nextID:          1,
		uniquePerEntity: uniquePerEntity,
	}
}

// Add inserts a new component into the store.
// If the component ID is zero, it assigns a new unique ID.
// Enforces uniqueness per entity if the flag is set.
func (s *Store[C]) Add(c C) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	entityID := c.Entity()

	if s.uniquePerEntity {
		if compIDs, exists := s.entityIndex[entityID]; exists && len(compIDs) > 0 {
			return fmt.Errorf("component for entity ID %d already exists with component ID %d", entityID, compIDs[0])
		}
	}

	if c.ID() == 0 {
		c.SetID(s.nextID)
		s.nextID++
	} else {
		if _, exists := s.components[c.ID()]; exists {
			return fmt.Errorf("component with ID %d already exists", c.ID())
		}
	}

	s.components[c.ID()] = c

	// Add component ID to entityIndex
	s.entityIndex[entityID] = append(s.entityIndex[entityID], c.ID())

	return nil
}

// Get retrieves a component by its ID.
func (s *Store[C]) Get(id uint32) (C, error) {
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
func (s *Store[C]) Update(c C) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.components[c.ID()]; !exists {
		return fmt.Errorf("component with ID %d not found", c.ID())
	}
	s.components[c.ID()] = c
	return nil
}

func (s *Store[C]) Delete(id uint32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, exists := s.components[id]
	if !exists {
		return ErrComponentNotFound
	}
	delete(s.components, id)

	entityID := c.Entity()

	// Remove component ID from entityIndex
	compIDs := s.entityIndex[entityID]
	for i, compID := range compIDs {
		if compID == id {
			// Remove the component ID from the slice
			s.entityIndex[entityID] = append(compIDs[:i], compIDs[i+1:]...)
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

	components := make([]C, 0, len(s.components))
	for _, c := range s.components {
		components = append(components, c)
	}
	return components
}

// FirstByEntity retrieves a component by its associated entity.
// If multiple components are associated with the entity, it returns the first one.
func (s *Store[C]) FirstByEntity(e entity.Entity) (C, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	compIDs, exists := s.entityIndex[e]
	if !exists || len(compIDs) == 0 {
		var zero C
		return zero, ErrEntityNotFound
	}

	// Return the first component associated with the entity
	c, exists := s.components[compIDs[0]]
	if !exists {
		var zero C
		return zero, fmt.Errorf("existing entity ID %d with missing component ID %d", e, compIDs[0])
	}
	return c, nil
}

// ListByEntity retrieves all components associated with an entity.
// If no components are associated with the entity, it returns an empty slice.
func (s *Store[C]) ListByEntity(e entity.Entity) ([]C, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	compIDs, exists := s.entityIndex[e]
	if !exists || len(compIDs) == 0 {
		return nil, ErrEntityNotFound
	}

	components := make([]C, 0, len(compIDs))
	for _, compID := range compIDs {
		if c, exists := s.components[compID]; exists {
			components = append(components, c)
		} else {
			return nil, fmt.Errorf("existing entity ID %d with missing component ID %d", e, compID)
		}
	}
	return components, nil
}

func (s *Store[C]) DeleteByEntity(e entity.Entity) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	compIDs, exists := s.entityIndex[e]
	if !exists {
		return ErrEntityNotFound // No components to delete for this entity
	}

	// Delete each component associated with the entity
	for _, compID := range compIDs {
		delete(s.components, compID)
	}

	// Remove the entity from the index
	delete(s.entityIndex, e)

	return nil
}
