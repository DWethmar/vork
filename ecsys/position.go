package ecsys

import (
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Position(e entity.Entity) (position.Position, error) { return s.pos.FirstByEntity(e) }
func (s *ECS) UpdatePosition(p position.Position) error            { return s.pos.Update(p) }
func (s *ECS) AddPosition(p position.Position) error               { return s.pos.Add(p) }
func (s *ECS) DeletePosition(id uint32) error                      { return s.pos.Delete(id) }
