package skeletons_test

import (
	"log/slog"
	"testing"

	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/store"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems/skeletons"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestNew(t *testing.T) {
	t.Run("New should create a new system and register event handlers", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())
		s := skeletons.New(slog.Default(), ecs, eventBus)
		if s == nil {
			t.Error("System should not be nil")
		}

		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 2 {
			t.Errorf("Expected 2 subscriptions, got %d", len(subscriptions))
		}

		e, err := ecs.CreateEntity(0, 0)
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}
		// should setup skeleton
		if err = eventBus.Publish(skeleton.NewCreatedEvent(skeleton.Skeleton{
			I: 1,
			E: e,
		})); err != nil {
			t.Errorf("Publish() error = %v", err)
		}

		// should have position
		if _, err = ecs.GetPosition(e); err != nil {
			t.Errorf("Expected position component, got %v", err)
		}

		// should have rectangle
		if _, err = ecs.GetRectangle(e); err != nil {
			t.Errorf("Expected shape component, got %v", err)
		}

		// should have sprite
		if r := ecs.ListSpritesByEntity(e); len(r) == 0 {
			t.Errorf("Expected sprite component, got %v", r)
		}
	})
}

func TestSystem_Draw(t *testing.T) {
	t.Run("Draw should not return an error", func(t *testing.T) {
		s := skeletons.New(slog.Default(), ecsys.New(event.NewBus(), store.NewStores()), event.NewBus())
		if err := s.Draw(&ebiten.Image{}); err != nil {
			t.Errorf("Draw() error = %v, wantErr %v", err, false)
		}
	})
}

func TestSystem_Update(t *testing.T) {
	t.Run("Update should not return an error", func(t *testing.T) {
		s := skeletons.New(slog.Default(), ecsys.New(event.NewBus(), store.NewStores()), event.NewBus())
		if err := s.Update(); err != nil {
			t.Errorf("Update() error = %v, wantErr %v", err, false)
		}
	})
}

func TestSystem_Close(t *testing.T) {
	t.Run("Close should not return an error and unsubscribe all event handlers", func(t *testing.T) {
		eventBus := event.NewBus()
		ecs := ecsys.New(eventBus, store.NewStores())

		s := skeletons.New(slog.Default(), ecs, eventBus)
		if err := s.Close(); err != nil {
			t.Errorf("Close() error = %v, wantErr %v", err, false)
		}

		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 0 {
			t.Errorf("Expected 0 subscriptions, got %d", len(subscriptions))
		}
	})
}
