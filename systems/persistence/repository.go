package persistence

import "github.com/dwethmar/vork/component"

type Repository[T component.Component] interface {
	// Get returns an component by its ID.
	Get(id uint32) (T, error)
	// Save saves the given component.
	Save(c T) error
	// Delete removes an component by its ID.
	Delete(id uint32) error
	// List returns all components.
	List() ([]T, error)
}
