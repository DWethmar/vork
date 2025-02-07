package game

import (
	"errors"
	"fmt"
	"log/slog"

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
	logger *slog.Logger
	scene  Scene            // Current scene
	scenes map[string]Scene // All scenes
}

// New creates a new game.
func New(logger *slog.Logger) *Game {
	return &Game{
		logger: logger,
		scene:  nil,
		scenes: make(map[string]Scene),
	}
}

// SwitchScene switches to a scene.
func (g *Game) SwitchScene(name string) error {
	if scene, ok := g.scenes[name]; ok {
		g.scene = scene
		return nil
	}
	return fmt.Errorf("scene %q not found", name)
}

// AddScene adds a scene to the game.
func (g *Game) AddScene(scene Scene) error {
	if _, ok := g.scenes[scene.Name()]; ok {
		return ErrSceneAlreadyExists
	}
	g.scenes[scene.Name()] = scene
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 400
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}
	if err := g.scene.Draw(screen); err != nil {
		g.logger.Error("failed to draw scene", slog.String("scene", g.scene.Name()), slog.String("error", err.Error()))
	}
}

// Update updates the game.
func (g *Game) Update() error {
	if g.scene == nil {
		return nil
	}
	return g.scene.Update()
}
