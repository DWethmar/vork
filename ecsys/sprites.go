package ecsys

import "github.com/dwethmar/vork/component/sprite"

func (s *ECS) Sprites() []sprite.Sprite            { return s.sprites.List() }
func (s *ECS) UpdateSprite(sp sprite.Sprite) error { return s.sprites.Update(sp) }
func (s *ECS) AddSprite(sp sprite.Sprite) error    { return s.sprites.Add(sp) }
func (s *ECS) DeleteSprite(id uint32) error        { return s.sprites.Delete(id) }
