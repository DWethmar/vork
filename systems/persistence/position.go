package persistence

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/ecsys"
	bolt "go.etcd.io/bbolt"
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

func (l *PositionLifeCycle) Commit(tx *bolt.Tx) error {
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

func (l *PositionLifeCycle) Load(tx *bolt.Tx, ecs *ecsys.ECS) error {
	positions, err := l.repo.List(tx)
	if err != nil {
		return err
	}

	for _, c := range positions {
		if _, err := ecs.AddPosition(*c); err != nil {
			return err
		}
	}

	return nil
}
