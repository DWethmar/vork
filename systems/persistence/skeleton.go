package persistence

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/skeleton"
)

var _ ComponentLifeCycle = &PositionLifeCycle{}

type SkeletonLifeCycle struct {
	repo    Repository[*skeleton.Skeleton]
	changed map[uint32]*skeleton.Skeleton
	deleted map[uint32]*skeleton.Skeleton
}

func NewSkeletonLifeCycle(repo Repository[*skeleton.Skeleton]) *SkeletonLifeCycle {
	return &SkeletonLifeCycle{
		repo:    repo,
		changed: make(map[uint32]*skeleton.Skeleton),
		deleted: make(map[uint32]*skeleton.Skeleton),
	}
}

func (l *SkeletonLifeCycle) Changed(e component.Event) error {
	if e.Deleted() {
		return nil
	}

	if _, ok := l.deleted[e.ComponentID()]; ok {
		return nil
	}

	p, ok := e.(skeleton.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", l, e)
	}

	l.changed[p.ComponentID()] = p.Skeleton()
	return nil
}

func (l *SkeletonLifeCycle) Deleted(e component.Event) error {
	p, ok := e.(skeleton.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", l, e)
	}
	l.deleted[p.ComponentID()] = p.Skeleton()
	return nil
}

func (l *SkeletonLifeCycle) Commit() error {
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
