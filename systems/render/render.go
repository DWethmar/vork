package render

import (
	"fmt"
	"image/color"
	"log/slog"
	"sort"

	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/ecsys"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	zoomFactor = 1.1
	minZoom    = 0.5
	maxZoom    = 5.0
)

type MouseHandler func(x, y int)

// Sprite is a sprite.
type Sprite struct {
	Graphic          sprite.Graphic
	Img              *ebiten.Image
	OffsetX, OffsetY int // Offset from the entity's position
}

// System is the rendering system.
type System struct {
	logger       *slog.Logger
	sprites      map[sprite.Graphic]*Sprite
	ecs          *ecsys.ECS
	offsetX      int
	offsetY      int
	zoom         float64
	clickHandler MouseHandler
	hoverHandler MouseHandler
}

// Options are the options for the rendering system.
type Options struct {
	Logger       *slog.Logger
	Sprites      []Sprite
	ECS          *ecsys.ECS
	ClickHandler MouseHandler
	HoverHandler MouseHandler
}

// New creates a new rendering system.
func New(opts Options) *System {
	spriteMap := make(map[sprite.Graphic]*Sprite)
	for _, s := range opts.Sprites {
		spriteMap[s.Graphic] = &s
	}
	return &System{
		logger:       opts.Logger.With("system", "render"),
		sprites:      spriteMap,
		ecs:          opts.ECS,
		offsetX:      0,
		offsetY:      0,
		zoom:         1.0,
		clickHandler: opts.ClickHandler,
		hoverHandler: opts.HoverHandler,
	}
}

func (s *System) Init() error {
	return nil
}

// Close closes the system.
func (s *System) Close() error {
	return nil
}

// entityDraw holds the information necessary to draw an entity.
type entityDraw struct {
	Index    int
	DrawFunc func(screen *ebiten.Image)
}

// Draw draws the entities on the screen.
func (s *System) Draw(screen *ebiten.Image) error {
	if err := renderGrid(screen, s.offsetX, s.offsetY, s.zoom, true); err != nil {
		return err
	}

	entitiesToDraw := []entityDraw{}
	// Collect rectangles to draw
	for _, r := range s.ecs.ListRectangles() {
		var x, y int
		if c, err := s.ecs.GetPosition(r.Entity()); err == nil {
			x, y = c.Cords()
		} else {
			return err
		}

		// Add the drawing function for this rectangle
		entitiesToDraw = append(entitiesToDraw, entityDraw{
			Index: y,
			DrawFunc: func(screen *ebiten.Image) {
				// Apply zoom factor to position and size
				x := (float32(x) - float32(s.offsetX)) * float32(s.zoom)
				y := (float32(y) - float32(s.offsetY)) * float32(s.zoom)
				width := float32(r.Width) * float32(s.zoom)
				height := float32(r.Height) * float32(s.zoom)
				vector.DrawFilledRect(screen, x, y, width, height, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, true)
			},
		})
	}

	// Collect sprites to draw
	for _, spc := range s.ecs.ListSprites() {
		var x, y int
		if c, err := s.ecs.GetPosition(spc.Entity()); err == nil {
			x, y = c.Cords()
		} else {
			return err
		}
		spr, ok := s.sprites[spc.Graphic]
		if !ok {
			return fmt.Errorf("sprite not found: %s", spc.Graphic)
		}

		// Apply sprite offsets
		x += spr.OffsetX
		y += spr.OffsetY

		// Add the drawing function for this sprite
		entitiesToDraw = append(entitiesToDraw, entityDraw{
			Index: y,
			DrawFunc: func(screen *ebiten.Image) {
				// Apply zoom factor to position and scale
				op := &ebiten.DrawImageOptions{}
				// Scale the sprite
				op.GeoM.Scale(s.zoom, s.zoom)
				// Translate the sprite
				x := (float64(x) - float64(s.offsetX)) * s.zoom
				y := (float64(y) - float64(s.offsetY)) * s.zoom
				op.GeoM.Translate(x, y)
				screen.DrawImage(spr.Img, op)
			},
		})
	}

	// Sort entities by their Y value to render them correctly
	sort.Slice(entitiesToDraw, func(i, j int) bool {
		return entitiesToDraw[i].Index < entitiesToDraw[j].Index
	})

	// Draw sorted entities
	for _, entity := range entitiesToDraw {
		entity.DrawFunc(screen)
	}

	return nil
}

func (s *System) Update() error {
	// Handle zoom in/out with mouse wheel
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		if wheelY > 0 {
			s.zoom *= zoomFactor
		} else if wheelY < 0 {
			s.zoom /= zoomFactor
		}
		// Clamp zoom level to prevent extreme zooming
		if s.zoom < minZoom {
			s.zoom = minZoom
		} else if s.zoom > maxZoom {
			s.zoom = maxZoom
		}
	}

	// get the first controllable entity and center the camera on it
	if controllables := s.ecs.ListControllables(); len(controllables) > 0 {
		// Get the first controllable entity
		firstControllable := controllables[0]

		// Get the position of the controllable
		pt, err := s.ecs.GetAbsolutePosition(firstControllable.Entity())
		if err != nil {
			return err
		}

		x, y := pt.Cords()

		// Get the actual screen dimensions from Ebiten
		screenWidth, screenHeight := ebiten.WindowSize()
		// Calculate the offsets to center the controllable on the screen, accounting for zoom
		s.offsetX = int(float64(x) - (float64(screenWidth) / (2 * s.zoom)))
		s.offsetY = int(float64(y) - (float64(screenHeight) / (2 * s.zoom)))
	}

	// Handle mouse click
	if s.clickHandler != nil && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Apply zoom factor to mouse position
		x, y := s.applyZoom(ebiten.CursorPosition())
		s.clickHandler(x, y)
	}

	// Handle mouse hover
	if s.hoverHandler != nil {
		// Apply zoom factor to mouse position
		x, y := s.applyZoom(ebiten.CursorPosition())
		s.hoverHandler(x, y)
	}

	return nil
}

// applyZoom applies the zoom factor to the given position.
func (s *System) applyZoom(x, y int) (int, int) {
	// Apply zoom factor to position
	x = int(float64(x)/s.zoom) + s.offsetX
	y = int(float64(y)/s.zoom) + s.offsetY
	return x, y
}
