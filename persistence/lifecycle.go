package persistence

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/ecsys"
	bolt "go.etcd.io/bbolt"
)

type ComponentLifeCycle interface {
	Changed(c component.Component, deleted bool) error // Changed is called when a component has changed.
	Commit(tx *bolt.Tx) error                          // Commit saves all changes to the database.
	Load(tx *bolt.Tx) error                            // Load function added here
}

type GenericComponentLifeCycle[T component.Component] struct {
	repo                Repository[T]
	changed             map[uint]T
	deleted             map[uint]T
	componentMarkerFunc func(component.Component, map[uint]T) error // Function to mark the component as changed
	store               ecsys.Store[T]
}

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

// Changed is called when a component has changed.
func (l *GenericComponentLifeCycle[T]) Changed(e component.Component, deleted bool) error {
	if deleted {
		delete(l.changed, e.ID())
		if err := l.componentMarkerFunc(e, l.deleted); err != nil {
			return fmt.Errorf("could not mark component as deleted: %w", err)
		}
	} else {
		if _, ok := l.deleted[e.ID()]; ok {
			return fmt.Errorf("component %d is already deleted", e.ID())
		}
		if err := l.componentMarkerFunc(e, l.changed); err != nil {
			return fmt.Errorf("could not mark component as changed: %w", err)
		}
	}
	return nil
}

// Commit saves all changes to the database.
func (l *GenericComponentLifeCycle[T]) Commit(tx *bolt.Tx) error {
	for _, p := range l.changed {
		if err := l.repo.Save(tx, p); err != nil {
			return err
		}
	}
	for _, p := range l.deleted {
		if err := l.repo.Delete(tx, p.ID()); err != nil {
			return err
		}
	}
	l.changed = make(map[uint]T)
	l.deleted = make(map[uint]T)
	return nil
}

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
