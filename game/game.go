package game

import (
	"log/slog"

	"github.com/dwethmar/vork/scene/memory"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems"
	"github.com/dwethmar/vork/systems/controller"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game updates and draws the game.
type Game struct {
	systems []systems.System
}

// New creates a new game.
func New() (*Game, error) {
	l := slog.Default()
	scene := memory.New()
	addPlayer(scene, 10, 10)
	addEnemy(scene, 100, 100)
	sprites, err := spritesheet.New()
	if err != nil {
		return nil, err
	}
	return &Game{
		systems: []systems.System{
			controller.NewSystem(scene),
			render.New(l, scene, sprites),
		},
	}, nil
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) error {
	for _, s := range g.systems {
		if err := s.Draw(screen); err != nil {
			return err
		}
	}
	return nil
}

// Update updates the game.
func (g *Game) Update() error {
	for _, s := range g.systems {
		if err := s.Update(); err != nil {
			return err
		}
	}
	return nil
}
