package mainmenu

import (
	"github.com/dwethmar/vork/game"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ game.Scene = &MainMenu{}

type MainMenu struct{}

// Draw implements game.Scene.
func (m *MainMenu) Draw(_ *ebiten.Image) error {
	panic("unimplemented")
}

// Name implements game.Scene.
func (m *MainMenu) Name() string {
	panic("unimplemented")
}

// Update implements game.Scene.
func (m *MainMenu) Update() error {
	panic("unimplemented")
}
