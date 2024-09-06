package controller

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type System struct {
	scene scene.Scene
}

func NewSystem(scene scene.Scene) *System {
	return &System{
		scene: scene,
	}
}

func (s *System) Update() error {
	x, y := direction()
	for _, c := range s.scene.ComponentsByType(controllable.Type) {
		if p, ok := s.scene.Component(c.Entity(), position.Type); ok {
			if p, ok := p.(*position.Position); ok {
				p.X += int64(x)
				p.Y += int64(y)
			}
		}
	}
	return nil
}

func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}
