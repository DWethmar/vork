package memory

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/scene"
)

var _ scene.Scene = &Memory{}

// CompKey is a key for a component.
type CompKey struct {
	E entity.Entity
	T component.ComponentType
}

type Memory struct {
	lastEntity       entity.Entity
	lastComponentID  uint32
	entities         map[entity.Entity][]component.Component
	components       map[CompKey]component.Component
	componentsByType map[component.ComponentType][]component.Component
}

func New() *Memory {
	return &Memory{
		lastEntity:       0,
		lastComponentID:  0,
		entities:         make(map[entity.Entity][]component.Component),
		components:       make(map[CompKey]component.Component),
		componentsByType: make(map[component.ComponentType][]component.Component),
	}
}

// Create creates a new entity and stores it in the memory store.
func (m *Memory) CreateEntity() entity.Entity {
	m.lastEntity++
	newEntity := m.lastEntity
	return newEntity
}

// Delete removes an entity and all its components from the memory store.
func (m *Memory) DeleteEntity(e entity.Entity) error {
	if _, exists := m.entities[e]; !exists {
		return fmt.Errorf("entity with ID %d does not exist", e)
	}

	// Retrieve all component IDs for the entity
	var componentIDs []uint32
	for _, comp := range m.entities[e] {
		componentIDs = append(componentIDs, comp.ID())
	}

	// Delete all components associated with the entity
	if err := m.DeleteComponents(e, componentIDs...); err != nil {
		return fmt.Errorf("failed to delete components for entity with ID %d: %w", e, err)
	}

	// Remove the entity from the entities map
	delete(m.entities, e)

	return nil
}

// Component retrieves a component by entity and component type from the memory store.
func (m *Memory) Component(e entity.Entity, t component.ComponentType) (component.Component, bool) {
	key := CompKey{E: e, T: t}
	comp, exists := m.components[key]
	return comp, exists
}

// Components retrieves all components associated with a given entity.
func (m *Memory) Components(e entity.Entity) ([]component.Component, error) {
	comps, exists := m.entities[e]
	if !exists {
		return nil, fmt.Errorf("entity with ID %d does not exist", e)
	}
	return comps, nil
}

// ComponentsByType retrieves all components of a given type.
func (m *Memory) ComponentsByType(t component.ComponentType) []component.Component {
	return m.componentsByType[t]
}

// AddComponent adds a new component to the specified entity and returns the component ID.
func (m *Memory) AddComponent(c component.Component) uint32 {
	e := c.Entity()
	// If the entity does not exist, create a new entry in the entities map
	if _, exists := m.entities[e]; !exists {
		m.entities[e] = []component.Component{}
	}

	// Increment the component ID counter and assign the ID to the component
	m.lastComponentID++
	c.SetID(m.lastComponentID)

	// Add to the entity's component list
	m.entities[e] = append(m.entities[e], c)

	// Add to the components map
	key := CompKey{E: e, T: c.Type()}
	m.components[key] = c

	// Add to the componentsByType map
	m.componentsByType[c.Type()] = append(m.componentsByType[c.Type()], c)

	return c.ID()
}

// DeleteComponents removes components with specified IDs from an entity.
func (m *Memory) DeleteComponents(e entity.Entity, ids ...uint32) error {
	if _, exists := m.entities[e]; !exists {
		return fmt.Errorf("entity with ID %d does not exist", e)
	}

	for _, id := range ids {
		found := false
		for i := len(m.entities[e]) - 1; i >= 0; i-- {
			comp := m.entities[e][i]
			if comp.ID() == id {
				// Remove from entity's component list
				m.entities[e] = append(m.entities[e][:i], m.entities[e][i+1:]...)

				// Remove from components map
				key := CompKey{E: e, T: comp.Type()}
				delete(m.components, key)

				// Remove from componentsByType
				compType := comp.Type()
				for j, c := range m.componentsByType[compType] {
					if c.ID() == id {
						m.componentsByType[compType] = append(m.componentsByType[compType][:j], m.componentsByType[compType][j+1:]...)
						break
					}
				}

				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("component with ID %d does not exist for entity with ID %d", id, e)
		}
	}
	return nil
}
