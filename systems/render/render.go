package render

import (
	"image/color"
	"log/slog"

	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type System struct {
	logger  *slog.Logger
	sprites *spritesheet.Sprites
	ecs     *systems.ECS
}

func New(
	logger *slog.Logger,
	sprites *spritesheet.Sprites,
	ecs *systems.ECS,
) *System {
	return &System{
		logger:  logger,
		sprites: sprites,
		ecs:     ecs,
	}
}

// Draw draws the shapes.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, r := range s.ecs.Rectangles() {
		var X, Y int64
		if c, err := s.ecs.Position(r.Entity()); err == nil {
			X, Y = c.X, c.Y
		} else {
			return err
		}
		// draw the rectangle
		vector.DrawFilledRect(screen, float32(X), float32(Y), float32(r.Width), float32(r.Height), color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, true)
	}
	return nil
}

func (s *System) Update() error {
	return nil
}
