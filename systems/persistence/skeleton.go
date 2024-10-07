package persistence

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
	bolt "go.etcd.io/bbolt"
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

func (l *SkeletonLifeCycle) Commit(tx *bolt.Tx) error {
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

func (l *SkeletonLifeCycle) Load(tx *bolt.Tx, ecs *ecsys.ECS) error {
	skeletons, err := l.repo.List(tx)
	if err != nil {
		return err
	}

	for _, c := range skeletons {
		if _, err := ecs.AddSkeleton(*c); err != nil {
			return err
		}
	}

	return nil
}
