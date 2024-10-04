package persistence

import "github.com/dwethmar/vork/component"

type Repository interface {
	// Save saves the given entity.
	Save(c component.Component) error
	// Delete removes an component by its ID.
	Delete(t component.ComponentType, id uint32) error
	// List returns all entities of the given type.
	List(t string) []component.Component
}
