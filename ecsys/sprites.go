package ecsys

import (
	"fmt"

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

// Sprite returns a sprite by its ID.
func (s *ECS) UpdateSprite(sp sprite.Sprite) error {
	if err := s.sprites.Update(&sp); err != nil {
		return fmt.Errorf("could not update sprite: %w", err)
	}
	return nil
}

// AddSprite adds a sprite to the ECS.
func (s *ECS) AddSprite(sp sprite.Sprite) (uint32, error) {
	id, err := s.sprites.Add(&sp)
	if err != nil {
		return 0, fmt.Errorf("could not add sprite: %v", err)
	}
	return id, nil
}

// DeleteSprite removes a sprite from the ECS.
func (s *ECS) DeleteSprite(id uint32) error {
	return s.sprites.Delete(id)
}
