package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) GetPosition(e entity.Entity) (position.Position, error) {
	c, err := s.stores.Position.First(e)
	if err != nil {
		return position.Position{}, fmt.Errorf("could not get position of entity %d: %w", e, err)
	}
	return *c, nil
}

func (s *ECS) GetVelocity(e entity.Entity) (velocity.Velocity, error) {
	c, err := s.stores.Velocity.First(e)
	if err != nil {
		return velocity.Velocity{}, fmt.Errorf("could not get velocity of entity %d: %w", e, err)
	}
	return *c, nil
}

func (s *ECS) GetControllable(e entity.Entity) (controllable.Controllable, error) {
	c, err := s.stores.Controllable.First(e)
	if err != nil {
		return controllable.Controllable{}, fmt.Errorf("could not get controllable of entity %d: %w", e, err)
	}
	return *c, nil
}

func (s *ECS) GetSkeleton(e entity.Entity) (skeleton.Skeleton, error) {
	c, err := s.stores.Skeleton.First(e)
	if err != nil {
		return skeleton.Skeleton{}, fmt.Errorf("could not get skeleton of entity %d: %w", e, err)
	}
	return *c, nil
}
