package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Skeleton(e entity.Entity) (skeleton.Skeleton, error) {
	c, err := s.sklt.FirstByEntity(e)
	if err != nil {
		return skeleton.Skeleton{}, fmt.Errorf("could not get skeleton: %v", err)
	}
	return *c, nil
}

func (s *ECS) Skeletons() []skeleton.Skeleton {
	c := s.sklt.List()
	r := make([]skeleton.Skeleton, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) UpdateSkeleton(sk skeleton.Skeleton) error {
	if err := s.sklt.Update(&sk); err != nil {
		return err
	}
	if err := s.eventBus.Publish(skeleton.NewUpdatedEvent(sk)); err != nil {
		return fmt.Errorf("could not publish event: %w", err)
	}
	return nil
}

func (s *ECS) AddSkeleton(sk skeleton.Skeleton) (uint32, error) {
	id, err := s.sklt.Add(&sk)
	if err != nil {
		return 0, fmt.Errorf("could not add skeleton: %w", err)
	}
	if err := s.eventBus.Publish(skeleton.NewCreatedEvent(sk)); err != nil {
		return 0, fmt.Errorf("could not publish event: %w", err)
	}
	return id, nil
}

func (s *ECS) DeleteSkeleton(sk skeleton.Skeleton) error {
	if err := s.sklt.Delete(sk.ID()); err != nil {
		return fmt.Errorf("could not delete skeleton: %w", err)
	}
	if err := s.eventBus.Publish(skeleton.NewDeletedEvent(sk)); err != nil {
		return fmt.Errorf("could not publish event: %w", err)
	}
	return nil
}
