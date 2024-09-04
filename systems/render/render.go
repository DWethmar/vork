package render

import (
	"image/color"
	"log/slog"

	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/entity/position"
	"github.com/dwethmar/vork/entity/shape"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type System struct {
	logger *slog.Logger
	cs     *entity.ComponentStore
}

func NewSystem(logger *slog.Logger, cs *entity.ComponentStore) *System {
	return &System{
		logger: logger,
		cs:     cs,
	}
}

// Draw draws the shapes.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, c := range s.cs.List(shape.Type) {
		var X, Y int64
		if p := s.cs.Get(c.Entity(), position.Type); p != nil {
			p := p.(*position.Position)
			X, Y = p.X, p.Y
		}
		switch t := c.(type) {
		case *shape.Rectangle:
			vector.DrawFilledRect(screen, float32(X), float32(Y), float32(t.Width), float32(t.Height), color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, true)
		default:
			s.logger.Warn("unknown shape", "type", c.Type())
		}
	}
	return nil
}

func (s *System) Update() error {
	return nil
}
