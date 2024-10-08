package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Position(e entity.Entity) (position.Position, error) {
	c, err := s.pos.FirstByEntity(e)
	if err != nil {
		return position.Position{}, fmt.Errorf("could not get position: %w", err)
	}
	return *c, nil
}
