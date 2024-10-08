package skeletons_test

import (
	"log/slog"
	"testing"

	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems/skeletons"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestNew(t *testing.T) {
	t.Run("New should create a new system and register event handlers", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus)
		s := skeletons.New(slog.Default(), ecs, eventBus)
		if s == nil {
			t.Error("System should not be nil")
		}

		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 1 {
			t.Errorf("Expected 1 subscriptions, got %d", len(subscriptions))
		}

		e := ecs.CreateEntity()
		// should setup skeleton
		eventBus.Publish(skeleton.NewCreatedEvent(skeleton.Skeleton{
			I: 1,
			E: e,
		}))

		// shoudl have position
		if _, err := ecs.Position(e); err != nil {
			t.Errorf("Expected position component, got %v", err)
		}

		// should have rectangle
		if _, err := ecs.Rectangle(e); err != nil {
			t.Errorf("Expected shape component, got %v", err)
		}

		// should have sprite
		if r := ecs.SpritesByEntity(e); len(r) == 0 {
			t.Errorf("Expected sprite component, got %v", r)
		}
	})
}

func TestSystem_Draw(t *testing.T) {
	t.Run("Draw should not return an error", func(t *testing.T) {
		s := skeletons.New(slog.Default(), ecsys.New(event.NewBus()), event.NewBus())
		if err := s.Draw(&ebiten.Image{}); err != nil {
			t.Errorf("Draw() error = %v, wantErr %v", err, false)
		}
	})
}

func TestSystem_Update(t *testing.T) {
	t.Run("Update should not return an error", func(t *testing.T) {
		s := skeletons.New(slog.Default(), ecsys.New(event.NewBus()), event.NewBus())
		if err := s.Update(); err != nil {
			t.Errorf("Update() error = %v, wantErr %v", err, false)
		}
	})
}
