package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/event"
)

func updateComponent[T any](
	eventBus *event.Bus,
	c component.Component,
	store BaseComponentStore[T],
	eventCreator func(T) event.Event,
) error {
	comp, ok := c.(T)
	if !ok {
		return fmt.Errorf("expected component of type %T, got %T", *new(T), c)
	}
	if err := store.Update(comp); err != nil {
		return fmt.Errorf("could not update component: %w", err)
	}
	if eventCreator != nil {
		if err := eventBus.Publish(eventCreator(comp)); err != nil {
			return fmt.Errorf("could not publish event: %w", err)
		}
	}
	return nil
}

func (s *ECS) UpdatePositionComponent(c position.Position) error {
	return updateComponent(s.eventBus, &c, s.positionStore, func(p *position.Position) event.Event {
		return position.NewUpdatedEvent(*p)
	})
}

func (s *ECS) UpdateControllableComponent(c controllable.Controllable) error {
	return updateComponent(s.eventBus, &c, s.controllableStore, func(ctr *controllable.Controllable) event.Event {
		return controllable.NewUpdatedEvent(*ctr)
	})
}

func (s *ECS) UpdateRectangleComponent(c shape.Rectangle) error {
	return updateComponent(s.eventBus, &c, s.rectangleStore, nil)
}

func (s *ECS) UpdateSpriteComponent(c sprite.Sprite) error {
	return updateComponent(s.eventBus, &c, s.spriteStore, nil)
}

func (s *ECS) UpdateSkeletonComponent(c skeleton.Skeleton) error {
	return updateComponent(s.eventBus, &c, s.skeletonStore, func(sk *skeleton.Skeleton) event.Event {
		return skeleton.NewUpdatedEvent(*sk)
	})
}
