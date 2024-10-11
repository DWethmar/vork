package ecsys

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) ListPositions() []position.Position {
	c := s.stores.Position.List()
	r := make([]position.Position, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) ListControllables() []controllable.Controllable {
	c := s.stores.Controllable.List()
	r := make([]controllable.Controllable, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) ListRectangles() []shape.Rectangle {
	c := s.stores.Rectangle.List()
	r := make([]shape.Rectangle, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) ListRectanglesByEntity(e entity.Entity) []shape.Rectangle {
	c := s.stores.Rectangle.ListByEntity(e)
	r := make([]shape.Rectangle, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) ListSprites() []sprite.Sprite {
	c := s.stores.Sprite.List()
	r := make([]sprite.Sprite, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

// SpritesByEntity returns all sprites for a given entity.
func (s *ECS) ListSpritesByEntity(e entity.Entity) []sprite.Sprite {
	c := s.stores.Sprite.ListByEntity(e)
	r := make([]sprite.Sprite, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) ListSkeletons() []skeleton.Skeleton {
	c := s.stores.Skeleton.List()
	r := make([]skeleton.Skeleton, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}
