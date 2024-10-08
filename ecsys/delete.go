package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/store"
	"github.com/dwethmar/vork/event"
)

func deleteComponent[T any](
	eventBus *event.Bus,
	c component.Component,
	store store.Store[T],
	eventCreator func(T) event.Event,
) error {
	comp, ok := c.(T)
	if !ok {
		return fmt.Errorf("expected component of type %T, got %T", *new(T), c)
	}
	if err := store.Delete(c.ID()); err != nil {
		return fmt.Errorf("could not update component: %w", err)
	}
	if eventCreator != nil {
		if err := eventBus.Publish(eventCreator(comp)); err != nil {
			return fmt.Errorf("could not publish event: %w", err)
		}
	}
	return nil
}

func (s *ECS) DeletePositionComponent(c position.Position) error {
	return deleteComponent(s.eventBus, &c, s.stores.Position, func(p *position.Position) event.Event {
		return position.NewDeletedEvent(*p)
	})
}

func (s *ECS) DeleteControllableComponent(c controllable.Controllable) error {
	return deleteComponent(s.eventBus, &c, s.stores.Controllable, func(ctr *controllable.Controllable) event.Event {
		return controllable.NewDeletedEvent(*ctr)
	})
}

func (s *ECS) DeleteRectangleComponent(c shape.Rectangle) error {
	return deleteComponent(s.eventBus, &c, s.stores.Rectangle, nil)
}

func (s *ECS) DeleteSpriteComponent(c sprite.Sprite) error {
	return deleteComponent(s.eventBus, &c, s.stores.Sprite, nil)
}

func (s *ECS) DeleteSkeletonComponent(c skeleton.Skeleton) error {
	return deleteComponent(s.eventBus, &c, s.stores.Skeleton, func(sk *skeleton.Skeleton) event.Event {
		return skeleton.NewDeletedEvent(*sk)
	})
}
