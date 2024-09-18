package scene

import (
	"github.com/dwethmar/vork/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene is a collection of entities and components.
type Scene struct {
	systems []systems.System
}

// Draw draws the game.
func (s *Scene) Draw(screen *ebiten.Image) error {
	for _, sys := range s.systems {
		if err := sys.Draw(screen); err != nil {
			return err
		}
	}
	return nil
}

// Update updates the game.
func (s *Scene) Update() error {
	for _, sys := range s.systems {
		if err := sys.Update(); err != nil {
			return err
		}
	}
	return nil
}

func New(s []systems.System) *Scene {
	return &Scene{
		systems: s,
	}
}
