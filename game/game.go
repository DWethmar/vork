package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Name() string
	Draw(screen *ebiten.Image) error
	Update() error
}

// Game updates and draws the game.
type Game struct {
	scene  Scene            // Current scene
	scenes map[string]Scene // All scenes
}

// New creates a new game.
func New(scenes map[string]Scene, startingScene string) (*Game, error) {
	return &Game{
		scene:  scenes[startingScene],
		scenes: scenes,
	}, nil
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) error {
	return g.scene.Draw(screen)
}

// Update updates the game.
func (g *Game) Update() error {
	return g.scene.Update()
}
