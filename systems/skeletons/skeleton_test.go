package skeletons_test

import (
	"log/slog"
	"testing"

	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems/skeletons"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestNew(t *testing.T) {
	t.Run("New should create a new system and register event handlers", func(t *testing.T) {
		eventBus := event.NewBus()

		positionStore := position.NewStore(eventBus)
		skeletonStore := skeleton.NewStore(eventBus)
		rectanleStore := shape.NewRectangleStore()
		spriteStore := sprite.NewStore()
		hitboxStore := hitbox.NewStore(eventBus)
		velocityStore := velocity.NewStore(eventBus)

		s := skeletons.New(slog.Default(), positionStore, skeletonStore, rectanleStore, spriteStore, hitboxStore, velocityStore, eventBus)
		if s == nil {
			t.Error("System should not be nil")
		}

		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 2 {
			t.Errorf("Expected 2 subscriptions, got %d", len(subscriptions))
		}

		sk := skeleton.New(1)

		// should setup skeleton
		if err := eventBus.Publish(skeleton.NewCreatedEvent(*sk)); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// should have position
		if _, err := positionStore.Get(1); err != nil {
			t.Errorf("Expected position component, got %v", err)
		}

		// should have rectangle
		if len(rectanleStore.All()) == 0 {
			t.Errorf("Expected rectangle component, got 0")
		}

		// should have sprite
		if len(spriteStore.All()) == 0 {
			t.Errorf("Expected sprite component, got 0")
		}
	})
}

func TestSystem_Draw(t *testing.T) {
	t.Run("Draw should not return an error", func(t *testing.T) {
		eventBus := event.NewBus()
		positionStore := position.NewStore(eventBus)
		skeletonStore := skeleton.NewStore(eventBus)
		rectanleStore := shape.NewRectangleStore()
		spriteStore := sprite.NewStore()
		hitboxStore := hitbox.NewStore(eventBus)
		velocityStore := velocity.NewStore(eventBus)
		s := skeletons.New(slog.Default(), positionStore, skeletonStore, rectanleStore, spriteStore, hitboxStore, velocityStore, eventBus)

		if s == nil {
			t.Error("System should not be nil")
		}
		if err := s.Draw(&ebiten.Image{}); err != nil {
			t.Errorf("Draw() error = %v, wantErr %v", err, false)
		}
	})
}

func TestSystem_Update(t *testing.T) {
	t.Run("Update should not return an error", func(t *testing.T) {
		eventBus := event.NewBus()
		positionStore := position.NewStore(eventBus)
		skeletonStore := skeleton.NewStore(eventBus)
		rectanleStore := shape.NewRectangleStore()
		spriteStore := sprite.NewStore()
		hitboxStore := hitbox.NewStore(eventBus)
		velocityStore := velocity.NewStore(eventBus)
		s := skeletons.New(slog.Default(), positionStore, skeletonStore, rectanleStore, spriteStore, hitboxStore, velocityStore, eventBus)

		if err := s.Update(); err != nil {
			t.Errorf("Update() error = %v, wantErr %v", err, false)
		}
	})
}

func TestSystem_Close(t *testing.T) {
	t.Run("Close should not return an error and unsubscribe all event handlers", func(t *testing.T) {
		eventBus := event.NewBus()
		positionStore := position.NewStore(eventBus)
		skeletonStore := skeleton.NewStore(eventBus)
		rectanleStore := shape.NewRectangleStore()
		spriteStore := sprite.NewStore()
		hitboxStore := hitbox.NewStore(eventBus)
		velocityStore := velocity.NewStore(eventBus)
		s := skeletons.New(slog.Default(), positionStore, skeletonStore, rectanleStore, spriteStore, hitboxStore, velocityStore, eventBus)
		if err := s.Close(); err != nil {
			t.Errorf("Close() error = %v, wantErr %v", err, false)
		}

		subscriptions := eventBus.Subscriptions()
		if len(subscriptions) != 0 {
			t.Errorf("Expected 0 subscriptions, got %d", len(subscriptions))
		}
	})
}
