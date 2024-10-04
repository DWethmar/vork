package ecsys

import (
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Rectangle(e entity.Entity) (shape.Rectangle, error) {
	c, err := s.rect.FirstByEntity(e)
	if err != nil {
		return shape.Rectangle{}, err
	}
	return *c, nil
}
func (s *ECS) Rectangles() []shape.Rectangle {
	c := s.rect.List()
	r := make([]shape.Rectangle, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}
func (s *ECS) UpdateRectangle(r shape.Rectangle) error {
	if err := s.rect.Update(&r); err != nil {
		return err
	}
	return nil
}
func (s *ECS) AddRectangle(r shape.Rectangle) (uint32, error) {
	id, err := s.rect.Add(&r)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ECS) DeleteRectangle(id uint32) error {
	return s.rect.Delete(id)
}
