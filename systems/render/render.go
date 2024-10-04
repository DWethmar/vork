package render

import (
	"fmt"
	"image/color"
	"log/slog"
	"sort"

	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/ecsys"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Sprite is a sprite.
type Sprite struct {
	Graphic          sprite.Graphic
	Img              *ebiten.Image
	OffsetX, OffsetY int
}

// System is the rendering system.
type System struct {
	logger  *slog.Logger
	sprites map[sprite.Graphic]*Sprite
	ecs     *ecsys.ECS
}

func New(
	logger *slog.Logger,
	sprites []Sprite,
	ecs *ecsys.ECS,
) *System {
	spriteMap := make(map[sprite.Graphic]*Sprite)
	for _, s := range sprites {
		spriteMap[s.Graphic] = &s
	}
	return &System{
		logger:  logger,
		sprites: spriteMap,
		ecs:     ecs,
	}
}

// entityDraw holds the information necessary to draw an entity.
type entityDraw struct {
	X, Y     int64
	DrawFunc func(screen *ebiten.Image)
}

// Draw draws the shapes.
func (s *System) Draw(screen *ebiten.Image) error {
	entitiesToDraw := []entityDraw{}

	// Collect rectangles to draw
	for _, r := range s.ecs.Rectangles() {
		var X, Y int64
		if c, err := s.ecs.Position(r.Entity()); err == nil {
			X, Y = c.X, c.Y
		} else {
			return err
		}

		// Add the drawing function for this rectangle
		entitiesToDraw = append(entitiesToDraw, entityDraw{
			X: X, Y: Y,
			DrawFunc: func(screen *ebiten.Image) {
				vector.DrawFilledRect(screen, float32(X), float32(Y), float32(r.Width), float32(r.Height), color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, true)
			},
		})
	}

	// Collect sprites to draw
	for _, spc := range s.ecs.Sprites() {
		var X, Y int64
		if c, err := s.ecs.Position(spc.Entity()); err == nil {
			X, Y = c.X, c.Y
		} else {
			return err
		}
		spr, ok := s.sprites[spc.Graphic]
		if !ok {
			return fmt.Errorf("sprite not found: %s", spc.Graphic)
		}
		// apply offset
		X += int64(spr.OffsetX)
		Y += int64(spr.OffsetY)

		// Add the drawing function for this sprite
		entitiesToDraw = append(entitiesToDraw, entityDraw{
			X: X, Y: Y,
			DrawFunc: func(screen *ebiten.Image) {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(X), float64(Y))
				screen.DrawImage(spr.Img, op)
			},
		})
	}

	// Sort entities by their Y value
	sort.Slice(entitiesToDraw, func(i, j int) bool {
		return entitiesToDraw[i].Y < entitiesToDraw[j].Y
	})

	// Draw sorted entities
	for _, entity := range entitiesToDraw {
		entity.DrawFunc(screen)
	}

	return nil
}

func (s *System) Update() error {
	return nil
}
