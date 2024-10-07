package persistence

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/position"
)

var _ ComponentLifeCycle = &PositionLifeCycle{}

type PositionLifeCycle struct {
	repo    Repository[*position.Position]
	changed map[uint32]*position.Position
	deleted map[uint32]*position.Position
}

func NewPositionLifeCycle(repo Repository[*position.Position]) *PositionLifeCycle {
	return &PositionLifeCycle{
		repo:    repo,
		changed: make(map[uint32]*position.Position),
		deleted: make(map[uint32]*position.Position),
	}
}

func (l *PositionLifeCycle) Changed(e component.Event) error {
	if e.Deleted() {
		return nil
	}

	if _, ok := l.deleted[e.ComponentID()]; ok {
		return nil
	}

	p, ok := e.(position.Event)
	if !ok {
		return nil
	}

	l.changed[p.ComponentID()] = p.Position()
	return nil
}

func (l *PositionLifeCycle) Deleted(e component.Event) error {
	p, ok := e.(position.Event)
	if !ok {
		return nil
	}
	l.deleted[p.ComponentID()] = p.Position()
	return nil
}

func (l *PositionLifeCycle) Commit() error {
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
