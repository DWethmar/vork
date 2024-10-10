package render_test

import (
	"log/slog"
	"testing"

	"github.com/dwethmar/vork/component/store"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		logger := slog.Default()
		sprites := []render.Sprite{}
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		clickHandler := func(_, _ int) {}
		got := render.New(logger, sprites, ecs, clickHandler)
		if got == nil {
			t.Errorf("New() = nil, want System")
		}
	})
}

func TestSystem_Close(t *testing.T) {
	t.Run("Close", func(t *testing.T) {
		logger := slog.Default()
		sprites := []render.Sprite{}
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		clickHandler := func(_, _ int) {}
		s := render.New(logger, sprites, ecs, clickHandler)
		if err := s.Close(); err != nil {
			t.Errorf("Close() = %v, want nil", err)
		}
	})
}

func TestSystem_Draw(t *testing.T) {
	t.Run("Draw", func(t *testing.T) {
		logger := slog.Default()
		sprites := []render.Sprite{}
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		clickHandler := func(_, _ int) {}
		s := render.New(logger, sprites, ecs, clickHandler)

		screen := ebiten.NewImage(100, 100)

		if err := s.Draw(screen); err != nil {
			t.Errorf("Draw() = %v, want nil", err)
		}
	})
}

func TestSystem_Update(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		logger := slog.Default()
		sprites := []render.Sprite{}
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		clickHandler := func(_, _ int) {}
		s := render.New(logger, sprites, ecs, clickHandler)

		if err := s.Update(); err != nil {
			t.Errorf("Update() = %v, want nil", err)
		}
	})
}
