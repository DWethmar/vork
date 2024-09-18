package render

import (
	"image/color"
	"log/slog"

	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/scene"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type System struct {
	logger  *slog.Logger
	scene   scene.Scene
	sprites *spritesheet.Sprites
}

func New(logger *slog.Logger, scene scene.Scene, sprites *spritesheet.Sprites) *System {
	return &System{
		logger:  logger,
		scene:   scene,
		sprites: sprites,
	}
}

// Draw draws the shapes.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, c := range s.scene.ComponentsByType(shape.Type) {
		var X, Y int64
		if c, ok := s.scene.Component(c.Entity(), position.Type); ok {
			if p, ok := c.(*position.Position); ok {
				X, Y = p.X, p.Y
			}
		}
		switch t := c.(type) {
		case *shape.Rectangle:
			vector.DrawFilledRect(screen, float32(X), float32(Y), float32(t.Width), float32(t.Height), color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, true)
		default:
			s.logger.Warn("unknown shape", "type", c.Type())
		}
	}
	for _, c := range s.scene.ComponentsByType(sprite.Type) {
		var X, Y int64
		if c, ok := s.scene.Component(c.Entity(), position.Type); ok {
			if p, ok := c.(*position.Position); ok {
				X, Y = p.X, p.Y
			}
		}
		switch t := c.(type) {
		case *sprite.Sprite:
			s.logger.Info("draw sprite", "sprite", t, "X", X, "Y", Y)
		default:
			s.logger.Warn("unknown sprite", "type", c.Type())
		}
	}
	return nil
}

func (s *System) Update() error {
	return nil
}
