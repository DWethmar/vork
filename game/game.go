package game

import (
	"log/slog"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/store/memory"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems"
	"github.com/dwethmar/vork/systems/controller"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game updates and draws the game.
type Game struct {
	scene *Scene
}

// New creates a new game.
func New() (*Game, error) {
	l := slog.Default()
	spritesheet, err := spritesheet.New()
	if err != nil {
		return nil, err
	}

	positionStore := memory.New[position.Position](true)
	controllableStore := memory.New[controllable.Controllable](true)
	rectangleStore := memory.New[shape.Rectangle](true)
	spriteStore := memory.New[sprite.Sprite](false)

	ecs := systems.NewECS(
		positionStore,
		controllableStore,
		rectangleStore,
		spriteStore,
	)

	addPlayer(ecs, 10, 10)
	addEnemy(ecs, 100, 100)

	return &Game{
		scene: NewScene([]systems.System{
			controller.New(l, ecs),
			render.New(l, Sprites(spritesheet), ecs),
		}),
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
