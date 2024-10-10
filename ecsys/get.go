package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) GetPosition(e entity.Entity) (position.Position, error) {
	c, err := s.stores.Position.FirstByEntity(e)
	if err != nil {
		return position.Position{}, fmt.Errorf("could not get position of entity %d: %w", e, err)
	}
	return *c, nil
}

func (s *ECS) GetControllable(e entity.Entity) (controllable.Controllable, error) {
	c, err := s.stores.Controllable.FirstByEntity(e)
	if err != nil {
		return controllable.Controllable{}, fmt.Errorf("could not get controllable of entity %d: %w", e, err)
	}
	return *c, nil
}

func (s *ECS) GetRectangle(e entity.Entity) (shape.Rectangle, error) {
	c, err := s.stores.Rectangle.FirstByEntity(e)
	if err != nil {
		return shape.Rectangle{}, fmt.Errorf("could not get rectangle of entity %d: %w", e, err)
	}
	return *c, nil
}

func (s *ECS) GetSkeleton(e entity.Entity) (skeleton.Skeleton, error) {
	c, err := s.stores.Skeleton.FirstByEntity(e)
	if err != nil {
		return skeleton.Skeleton{}, fmt.Errorf("could not get skeleton of entity %d: %w", e, err)
	}
	return *c, nil
}
