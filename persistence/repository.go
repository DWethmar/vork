package persistence

import (
	"github.com/dwethmar/vork/component"
	bolt "go.etcd.io/bbolt"
)

// Repository is a generic interface for a component repository.
type Repository[T component.Component] interface {
	// Get returns an component by its ID.
	Get(tx *bolt.Tx, id uint) (T, error)
	// Save saves the given component.
	Save(tx *bolt.Tx, c T) error
	// Delete removes an component by its ID.
	Delete(tx *bolt.Tx, id uint) error
	// List returns all components.
	List(tx *bolt.Tx) ([]T, error)
}
