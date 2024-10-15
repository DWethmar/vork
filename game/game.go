package game

import (
	"errors"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ErrSceneAlreadyExists = errors.New("scene already exists")
)

// Scene is a game scene. A scene is a part of the game, like a menu, the gameplay, etc.
type Scene interface {
	Name() string
	Draw(screen *ebiten.Image) error
	Update() error
	Close() error
}

// Game updates and draws the game.
type Game struct {
	scene  Scene            // Current scene
	scenes map[string]Scene // All scenes
}

// New creates a new game.
func New() (*Game, error) {
	return &Game{
		scene:  nil,
		scenes: make(map[string]Scene),
	}, nil
}

// SwitchScene switches to a scene.
func (g *Game) SwitchScene(name string) error {
	if scene, ok := g.scenes[name]; ok {
		g.scene = scene
		return nil
	}
	return fmt.Errorf("scene %s not found", name)
}

// AddScene adds a scene to the game.
func (g *Game) AddScene(scene Scene) error {
	if _, ok := g.scenes[scene.Name()]; ok {
		return ErrSceneAlreadyExists
	}
	g.scenes[scene.Name()] = scene
	return nil
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) error {
	if g.scene == nil {
		return nil
	}
	return g.scene.Draw(screen)
}

// Update updates the game.
func (g *Game) Update() error {
	if g.scene == nil {
		return nil
	}
	return g.scene.Update()
}
