package persistence

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/ecsys"
	bolt "go.etcd.io/bbolt"
)

type ComponentLifeCycle interface {
	Changed(component.Event) error          // Changed is called when a component has changed.
	Deleted(component.Event) error          // Deleted is called when a component has been deleted.
	Commit(tx *bolt.Tx) error               // Commit saves all changes to the database.
	Load(tx *bolt.Tx, ecs *ecsys.ECS) error // Load function added here
}

type GenericComponentLifeCycle[T component.Component] struct {
	repo                Repository[T]
	changed             map[uint32]T
	deleted             map[uint32]T
	componentMarkerFunc func(component.Event, map[uint32]T) error // Function to mark the component as changed
	addFunc             func(T) (uint32, error)                   // Function to add the component to the ECS
}

func NewGenericComponentLifeCycle[T component.Component](
	repo Repository[T],
	addFunc func(T) (uint32, error),
	componentMarkerFunc func(component.Event, map[uint32]T) error,
) *GenericComponentLifeCycle[T] {
	return &GenericComponentLifeCycle[T]{
		repo:                repo,
		changed:             make(map[uint32]T),
		deleted:             make(map[uint32]T),
		componentMarkerFunc: componentMarkerFunc,
		addFunc:             addFunc,
	}
}

func (l *GenericComponentLifeCycle[T]) Changed(e component.Event) error {
	if e.Deleted() {
		return nil
	}
	if _, ok := l.deleted[e.ComponentID()]; ok {
		return fmt.Errorf("component %d is already deleted", e.ComponentID())
	}
	if err := l.componentMarkerFunc(e, l.changed); err != nil {
		return fmt.Errorf("could not mark component as changed: %w", err)
	}
	return nil
}

func (l *GenericComponentLifeCycle[T]) Deleted(e component.Event) error {
	delete(l.changed, e.ComponentID())
	if err := l.componentMarkerFunc(e, l.deleted); err != nil {
		return fmt.Errorf("could not mark component as deleted: %w", err)
	}
	return nil
}

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

	l.changed = make(map[uint32]T)
	l.deleted = make(map[uint32]T)

	return nil
}

func (l *GenericComponentLifeCycle[T]) Load(tx *bolt.Tx, ecs *ecsys.ECS) error {
	components, err := l.repo.List(tx)
	if err != nil {
		return err
	}

	for _, c := range components {
		if _, err := l.addFunc(c); err != nil {
			return err
		}
	}

	return nil
}
