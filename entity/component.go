package entity

import "sort"

// Component is an interface that all components must implement.
type Component interface {
	ID() uint32
	SetID(uint32)
	Type() string
	Entity() Entity
}

// ComponentStore is a manager that holds all components.
type ComponentStore struct {
	lastID     uint32 // last assigned ID
	components map[string]map[Entity]Component
}

// NewComponentManager creates a new ComponentManager.
func NewComponentStore(lastID uint32) *ComponentStore {
	return &ComponentStore{
		components: make(map[string]map[Entity]Component),
	}
}

// Add adds a component to the manager.
// If the component has no ID, it will be assigned one.
func (cm *ComponentStore) Add(c Component) {
	componentType := c.Type()
	entity := c.Entity()
	if c.ID() == 0 {
		cm.lastID++
		c.SetID(cm.lastID)
	}
	if _, exists := cm.components[componentType]; !exists {
		cm.components[componentType] = make(map[Entity]Component)
	}
	cm.components[componentType][entity] = c
}

// Remove removes a component from the manager.
func (cm *ComponentStore) Remove(e Entity, componentType string) {
	if entityComponents, exists := cm.components[componentType]; exists {
		delete(entityComponents, e)
		if len(entityComponents) == 0 {
			delete(cm.components, componentType)
		}
	}
}

// Get returns a component from the manager.
func (cm *ComponentStore) Get(e Entity, componentType string) Component {
	if entityComponents, exists := cm.components[componentType]; exists {
		if component, exists := entityComponents[e]; exists {
			return component
		}
	}
	return nil
}

// List returns all components of a certain type.
func (cm *ComponentStore) List(componentType string) []Component {
	var components []Component
	if entityComponents, exists := cm.components[componentType]; exists {
		for _, component := range entityComponents {
			components = append(components, component)
		}
	}
	// sort components by ID
	sort.Slice(components, func(i, j int) bool {
		return components[i].ID() < components[j].ID()
	})
	return components
}
