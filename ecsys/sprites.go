package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/sprite"
)

func (s *ECS) Sprites() []sprite.Sprite {
	c := s.sprites.List()
	r := make([]sprite.Sprite, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}
func (s *ECS) UpdateSprite(sp sprite.Sprite) error {
	if err := s.sprites.Update(&sp); err != nil {
		return fmt.Errorf("could not update sprite: %v", err)
	}
	return nil
}

func (s *ECS) AddSprite(sp sprite.Sprite) (uint32, error) {
	id, err := s.sprites.Add(&sp)
	if err != nil {
		return 0, fmt.Errorf("could not add sprite: %v", err)
	}
	return id, nil
}

func (s *ECS) DeleteSprite(id uint32) error {
	return s.sprites.Delete(id)
}
