package game

import (
	"log/slog"

	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/entity/controllable"
	"github.com/dwethmar/vork/entity/position"
	"github.com/dwethmar/vork/entity/shape"
	"github.com/dwethmar/vork/systems"
	"github.com/dwethmar/vork/systems/controller"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	systems []systems.System
}

func New() *Game {
	l := slog.Default()
	cs := entity.NewComponentStore(0)

	cs.Add(&shape.Rectangle{
		E:      1,
		Width:  32,
		Height: 32,
	})

	cs.Add(&position.Position{
		E: 1,
		X: 100,
		Y: 100,
	})

	cs.Add(&controllable.Controllable{
		E: 1,
	})

	return &Game{
		systems: []systems.System{
			controller.NewSystem(cs),
			render.NewSystem(l, cs),
		},
	}
}

func (g *Game) Draw(screen *ebiten.Image) error {
	for _, s := range g.systems {
		if err := s.Draw(screen); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Update() error {
	for _, s := range g.systems {
		if err := s.Update(); err != nil {
			return err
		}
	}
	return nil
}
