package ecsys

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
)

// ListPositions returns all positions.
func (s *ECS) ListPositions() []position.Position {
	c := s.stores.Position.List()
	r := make([]position.Position, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) ListControllables() []controllable.Controllable {
	c := s.stores.Controllable.List()
	r := make([]controllable.Controllable, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) ListHitboxesByEntity(e entity.Entity) []hitbox.Hitbox {
	c := s.stores.Hitbox.ListByEntity(e)
	r := make([]hitbox.Hitbox, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) ListRectangles() []shape.Rectangle {
	c := s.stores.Rectangle.List()
	r := make([]shape.Rectangle, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) ListRectanglesByEntity(e entity.Entity) []shape.Rectangle {
	c := s.stores.Rectangle.ListByEntity(e)
	r := make([]shape.Rectangle, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

// ListSprites returns all sprites.
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

// ListSkeletons returns all skeletons.
func (s *ECS) ListSkeletons() []skeleton.Skeleton {
	c := s.stores.Skeleton.List()
	r := make([]skeleton.Skeleton, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}
