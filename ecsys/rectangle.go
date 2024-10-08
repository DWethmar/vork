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
