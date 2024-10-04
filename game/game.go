package game

import (
	"log/slog"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems"
	"github.com/dwethmar/vork/systems/controller"
	"github.com/dwethmar/vork/systems/render"
	"github.com/dwethmar/vork/systems/skeletons"
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

	positionStore := component.NewStore[position.Position](true)
	controllableStore := component.NewStore[controllable.Controllable](true)
	rectangleStore := component.NewStore[shape.Rectangle](true)
	spriteStore := component.NewStore[sprite.Sprite](false)
	skeletonStore := component.NewStore[skeleton.Skeleton](true)

	eventBus := event.NewBus()

	ecs := ecsys.New(
		eventBus,
		positionStore,
		controllableStore,
		rectangleStore,
		spriteStore,
		skeletonStore,
	)

	systems := []systems.System{
		controller.New(l, ecs),
		render.New(l, Sprites(spritesheet), ecs),
		skeletons.New(l, ecs, eventBus),
	}

	addPlayer(ecs, 10, 10)
	addEnemy(ecs, 100, 100)

	return &Game{
		scene: NewScene(systems),
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
