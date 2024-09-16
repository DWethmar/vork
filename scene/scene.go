package scene

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

// Scene is a collection of entities and components.
type Scene interface {
	// CreateEntity creates a new entity and returns it.
	CreateEntity() entity.Entity
	// DeleteEntity deletes an entity and all its components.
	DeleteEntity(entity.Entity) error
	// Component returns a component for an entity.
	Component(entity.Entity, component.ComponentType) (component.Component, bool)
	// Components returns all components for an entity.
	Components(entity.Entity) ([]component.Component, error)
	// ComponentsByType returns all components of a certain type.
	ComponentsByType(component.ComponentType) []component.Component
	// AddComponent adds a component to an entity.
	AddComponent(component.Component) uint32
	// DeleteComponents deletes components from an entity.
	DeleteComponents(entity.Entity, ...uint32) error
}
