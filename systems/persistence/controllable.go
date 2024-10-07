package persistence

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
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

func (l *ControllableLifeCycle) Commit() error {
	for _, p := range l.changed {
		if err := l.repo.Save(p); err != nil {
			return err
		}
	}

	for _, p := range l.deleted {
		if err := l.repo.Delete(p.ID()); err != nil {
			return err
		}
	}

	return nil
}
