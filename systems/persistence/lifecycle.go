package persistence

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/ecsys"
	bolt "go.etcd.io/bbolt"
)

// ComponentLifeCycle defines the interface for handling the lifecycle of components.
// It tracks when components change, allows saving changes to a database (via Commit),
// and enables loading components from a database (via Load).
type ComponentLifeCycle interface {
	Changed(c component.Component, deleted bool) error // Changed is called when a component has changed.
	Commit(tx *bolt.Tx) error                          // Commit saves all changes to the database.
	Load(tx *bolt.Tx) error                            // Load function added here
}

// GenericComponentLifeCycle is a generic implementation of the ComponentLifeCycle interface.
// It tracks changes and deletions of components of type T and interacts with a repository
// to persist those changes to a BoltDB database.
type GenericComponentLifeCycle[T component.Component] struct {
	repo                Repository[T]
	changed             map[uint]T
	deleted             map[uint]T
	componentMarkerFunc func(component.Component, map[uint]T) error // Function to mark the component as changed
	store               ecsys.Store[T]
}

// NewGenericComponentLifeCycle creates and initializes a new GenericComponentLifeCycle.
// It requires a repository for database interaction, a store for system state management,
// and a function to mark components as changed or deleted.
func NewGenericComponentLifeCycle[T component.Component](
	repo Repository[T],
	store ecsys.Store[T],
	componentMarkerFunc func(component.Component, map[uint]T) error,
) *GenericComponentLifeCycle[T] {
	return &GenericComponentLifeCycle[T]{
		repo:                repo,
		changed:             make(map[uint]T),
		deleted:             make(map[uint]T),
		componentMarkerFunc: componentMarkerFunc,
		store:               store,
	}
}

// Changed records a component's state change (or deletion) in the lifecycle.
// If the component is marked as deleted, it's added to the `deleted` map and removed from the `changed` map.
// If the component is modified, it is added to the `changed` map unless it is marked as deleted already.
func (l *GenericComponentLifeCycle[T]) Changed(e component.Component, deleted bool) error {
	if deleted {
		delete(l.changed, e.ID())
		if err := l.componentMarkerFunc(e, l.deleted); err != nil {
			return fmt.Errorf("could not mark component as deleted: %w", err)
		}
		return nil
	}
	if _, ok := l.deleted[e.ID()]; ok {
		return fmt.Errorf("component %d is marked as deleted and cannot be modified", e.ID())
	}
	if err := l.componentMarkerFunc(e, l.changed); err != nil {
		return fmt.Errorf("could not mark component as changed: %w", err)
	}
	return nil
}

// Commit saves all recorded changes to the database.
// It goes through both the `changed` and `deleted` maps, saving and removing components from the database
// as necessary, then clears these maps to reflect that all changes have been persisted.
func (l *GenericComponentLifeCycle[T]) Commit(tx *bolt.Tx) error {
	for _, p := range l.changed {
		if err := l.repo.Save(tx, p); err != nil {
			return fmt.Errorf("failed to save component: %w", err)
		}
	}
	for _, p := range l.deleted {
		if err := l.repo.Delete(tx, p.ID()); err != nil {
			return fmt.Errorf("failed to delete component: %w", err)
		}
	}
	l.changed = make(map[uint]T)
	l.deleted = make(map[uint]T)
	return nil
}

// Load retrieves components from the database using the repository and adds them to the store.
// This method is used to initialize the system's state from a persisted state stored in the database.
func (l *GenericComponentLifeCycle[T]) Load(tx *bolt.Tx) error {
	components, err := l.repo.List(tx)
	if err != nil {
		return err
	}
	for _, c := range components {
		if _, err = l.store.Add(c); err != nil {
			return err
		}
	}
	return nil
}
