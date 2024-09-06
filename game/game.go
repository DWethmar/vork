package game

import (
	"log/slog"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/scene/memory"
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
	scene := memory.New()

	e := scene.CreateEntity()
	scene.AddComponent(shape.NewRectangle(e, 10, 10))
	scene.AddComponent(position.New(e, 100, 100))
	scene.AddComponent(controllable.New(e))

	return &Game{
		systems: []systems.System{
			controller.NewSystem(scene),
			render.NewSystem(l, scene),
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
