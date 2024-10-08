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

// addComponent adds a component to the ECS.
func addComponent[T any](
	eventBus *event.Bus,
	c component.Component,
	store BaseComponentStore[T],
	eventCreator func(T) event.Event,
) (uint32, error) {
	comp, ok := c.(T)
	if !ok {
		return 0, fmt.Errorf("expected component of type %T, got %T", *new(T), c)
	}
	id, err := store.Add(comp)
	if err != nil {
		return 0, fmt.Errorf("could not add component: %w", err)
	}
	if eventCreator != nil {
		if err := eventBus.Publish(eventCreator(comp)); err != nil {
			return 0, fmt.Errorf("could not publish event: %w", err)
		}
	}
	return id, nil
}

func (s *ECS) AddPositionComponent(c position.Position) (uint32, error) {
	return addComponent(s.eventBus, &c, s.positionStore, func(p *position.Position) event.Event {
		return position.NewCreatedEvent(*p)
	})
}

func (s *ECS) AddControllableComponent(c controllable.Controllable) (uint32, error) {
	return addComponent(s.eventBus, &c, s.controllableStore, func(ctr *controllable.Controllable) event.Event {
		return controllable.NewCreatedEvent(*ctr)
	})
}

func (s *ECS) AddRectangleComponent(c shape.Rectangle) (uint32, error) {
	return addComponent(s.eventBus, &c, s.rectangleStore, nil)
}

func (s *ECS) AddSpriteComponent(c sprite.Sprite) (uint32, error) {
	return addComponent(s.eventBus, &c, s.spriteStore, nil)
}

func (s *ECS) AddSkeletonComponent(c skeleton.Skeleton) (uint32, error) {
	return addComponent(s.eventBus, &c, s.skeletonStore, func(sk *skeleton.Skeleton) event.Event {
		return skeleton.NewCreatedEvent(*sk)
	})
}
