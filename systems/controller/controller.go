package controller

import (
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/entity/controllable"
	"github.com/dwethmar/vork/entity/position"
	"github.com/hajimehoshi/ebiten/v2"
)

type System struct {
	cs *entity.ComponentStore
}

func NewSystem(cs *entity.ComponentStore) *System {
	return &System{
		cs: cs,
	}
}

func (s *System) Update() error {
	x, y := direction()
	for _, c := range s.cs.List(controllable.Type) {
		p := s.cs.Get(c.Entity(), position.Type)
		if p, ok := p.(*position.Position); ok {
			p.X += int64(x)
			p.Y += int64(y)
		}
	}
	return nil
}

func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}
