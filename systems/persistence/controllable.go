package persistence

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/ecsys"
	bolt "go.etcd.io/bbolt"
)

var _ ComponentLifeCycle = &PositionLifeCycle{}

type ControllableLifeCycle struct {
	repo    Repository[*controllable.Controllable]
	changed map[uint32]*controllable.Controllable
	deleted map[uint32]*controllable.Controllable
}

func NewControllableLifeCycle(repo Repository[*controllable.Controllable]) *ControllableLifeCycle {
	return &ControllableLifeCycle{
		repo:    repo,
		changed: make(map[uint32]*controllable.Controllable),
		deleted: make(map[uint32]*controllable.Controllable),
	}
}

func (l *ControllableLifeCycle) Changed(e component.Event) error {
	if e.Deleted() {
		return nil
	}

	if _, ok := l.deleted[e.ComponentID()]; ok {
		return nil
	}

	p, ok := e.(controllable.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", l, e)
	}

	l.changed[p.ComponentID()] = p.Controllable()
	return nil
}

func (l *ControllableLifeCycle) Deleted(e component.Event) error {
	p, ok := e.(controllable.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", l, e)
	}
	l.deleted[p.ComponentID()] = p.Controllable()
	return nil
}

func (l *ControllableLifeCycle) Commit(tx *bolt.Tx) error {
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

	return nil
}

func (l *ControllableLifeCycle) Load(tx *bolt.Tx, ecs *ecsys.ECS) error {
	components, err := l.repo.List(tx)
	if err != nil {
		return err
	}

	for _, c := range components {
		if _, err := ecs.AddControllable(*c); err != nil {
			return err
		}
	}

	return nil
}
