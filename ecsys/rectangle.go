package ecsys

import (
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Rectangle(e entity.Entity) (shape.Rectangle, error) { return s.rect.FirstByEntity(e) }
func (s *ECS) Rectangles() []shape.Rectangle                      { return s.rect.List() }
func (s *ECS) UpdateRectangle(r shape.Rectangle) error            { return s.rect.Update(r) }
func (s *ECS) AddRectangle(r shape.Rectangle) (uint32, error)     { return s.rect.Add(r) }
func (s *ECS) DeleteRectangle(id uint32) error                    { return s.rect.Delete(id) }
