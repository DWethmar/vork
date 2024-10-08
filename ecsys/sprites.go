package ecsys

import (
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
)

// Sprites returns all sprites in the ECS.
func (s *ECS) Sprites() []sprite.Sprite {
	c := s.sprites.List()
	r := make([]sprite.Sprite, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

// SpritesByEntity returns all sprites for a given entity.
func (s *ECS) SpritesByEntity(e entity.Entity) []sprite.Sprite {
	c := s.sprites.ListByEntity(e)
	r := make([]sprite.Sprite, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}
