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
	offsetX int
	offsetY int
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
		offsetX: 0,
		offsetY: 0,
	}
}

// entityDraw holds the information necessary to draw an entity.
type entityDraw struct {
	X, Y     int64
	DrawFunc func(screen *ebiten.Image)
}

// Draw draws the entities on the screen.
func (s *System) Draw(screen *ebiten.Image) error {
	entitiesToDraw := []entityDraw{}

	// Draw checkered background
	for x := 0; x < 800; x += 40 {
		for y := 0; y < 600; y += 40 {
			dx := float32(x - s.offsetX)
			dy := float32(y - s.offsetY)
			if (x/40+y/40)%2 == 0 {
				vector.DrawFilledRect(screen, dx, dy, 40, 40, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, true)
			} else {
				vector.DrawFilledRect(screen, dx, dy, 40, 40, color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}, true)
			}
		}
	}

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
				// Subtract the offset to correctly position the entity relative to the camera
				vector.DrawFilledRect(screen, float32(X)-float32(s.offsetX), float32(Y)-float32(s.offsetY), float32(r.Width), float32(r.Height), color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, true)
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
		// Apply offset to the sprite
		X += int64(spr.OffsetX)
		Y += int64(spr.OffsetY)

		// Add the drawing function for this sprite
		entitiesToDraw = append(entitiesToDraw, entityDraw{
			X: X, Y: Y,
			DrawFunc: func(screen *ebiten.Image) {
				op := &ebiten.DrawImageOptions{}
				// Subtract the offset to correctly position the sprite relative to the camera
				op.GeoM.Translate(float64(X)-float64(s.offsetX), float64(Y)-float64(s.offsetY))
				screen.DrawImage(spr.Img, op)
			},
		})
	}

	// Sort entities by their Y value to render them correctly
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
	controllables := s.ecs.Controllables()
	if len(controllables) > 0 {
		// Get the first controllable entity
		firstControllable := controllables[0]

		// Get the position of the controllable
		pos, err := s.ecs.Position(firstControllable.Entity())
		if err != nil {
			return err
		}

		// Get the actual screen dimensions from Ebiten
		screenWidth, screenHeight := ebiten.WindowSize()

		// Calculate the offsets to center the controllable on the screen
		s.offsetX = int(float64(pos.X) - float64(screenWidth)/2)
		s.offsetY = int(float64(pos.Y) - float64(screenHeight)/2)
	}

	return nil
}
