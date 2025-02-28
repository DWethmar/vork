package render_test

import (
	"log/slog"
	"testing"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		eventBus := event.NewBus()
		got := render.New(render.Options{
			Logger:            slog.Default(),
			Sprites:           []render.Sprite{},
			PositionsStore:    position.NewStore(eventBus),
			RectanglesStore:   shape.NewRectangleStore(),
			SpriteStore:       sprite.NewStore(),
			ControllableStore: controllable.NewStore(eventBus),
			ClickHandler:      func(_, _ int) {},
		})
		if got == nil {
			t.Errorf("New() = nil, want System")
		}
	})
}

func TestSystem_Close(t *testing.T) {
	t.Run("Close", func(t *testing.T) {
		eventBus := event.NewBus()
		s := render.New(render.Options{
			Logger:            slog.Default(),
			Sprites:           []render.Sprite{},
			PositionsStore:    position.NewStore(eventBus),
			RectanglesStore:   shape.NewRectangleStore(),
			SpriteStore:       sprite.NewStore(),
			ControllableStore: controllable.NewStore(eventBus),
			ClickHandler:      func(_, _ int) {},
		})
		if err := s.Close(); err != nil {
			t.Errorf("Close() = %v, want nil", err)
		}
	})
}

func TestSystem_Draw(t *testing.T) {
	t.Run("Draw", func(t *testing.T) {
		eventBus := event.NewBus()
		s := render.New(render.Options{
			Logger:            slog.Default(),
			Sprites:           []render.Sprite{},
			PositionsStore:    position.NewStore(eventBus),
			RectanglesStore:   shape.NewRectangleStore(),
			SpriteStore:       sprite.NewStore(),
			ControllableStore: controllable.NewStore(eventBus),
			ClickHandler:      func(_, _ int) {},
		})

		screen := ebiten.NewImage(100, 100)

		if err := s.Draw(screen); err != nil {
			t.Errorf("Draw() = %v, want nil", err)
		}
	})
}

func TestSystem_Update(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		eventBus := event.NewBus()
		s := render.New(render.Options{
			Logger:            slog.Default(),
			Sprites:           []render.Sprite{},
			PositionsStore:    position.NewStore(eventBus),
			RectanglesStore:   shape.NewRectangleStore(),
			SpriteStore:       sprite.NewStore(),
			ControllableStore: controllable.NewStore(eventBus),
			ClickHandler:      func(_, _ int) {},
		})
		if err := s.Update(); err != nil {
			t.Errorf("Update() = %v, want nil", err)
		}
	})
}
