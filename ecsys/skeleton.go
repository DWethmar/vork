package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Skeleton(e entity.Entity) (skeleton.Skeleton, error) {
	return s.sklt.FirstByEntity(e)
}

func (s *ECS) Skeletons() []skeleton.Skeleton {
	return s.sklt.List()
}

func (s *ECS) UpdateSkeleton(sk skeleton.Skeleton) error {
	if err := s.sklt.Update(sk); err != nil {
		return err
	}
	s.eventBus.Publish(&skeleton.UpdatedEvent{Skeleton: sk})
	return nil
}

func (s *ECS) AddSkeleton(sk skeleton.Skeleton) (uint32, error) {
	id, err := s.sklt.Add(sk)
	if err != nil {
		return 0, fmt.Errorf("could not add skeleton: %w", err)
	}
	s.eventBus.Publish(&skeleton.CreatedEvent{Skeleton: sk})
	return id, nil
}

func (s *ECS) DeleteSkeleton(sk skeleton.Skeleton) error {
	if err := s.sklt.Delete(sk.ID()); err != nil {
		return fmt.Errorf("could not delete skeleton: %w", err)
	}
	s.eventBus.Publish(&skeleton.DeletedEvent{Skeleton: sk})
	return nil
}
