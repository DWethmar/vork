package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/event"
)

func updateComponent[T component.Component](
	eventBus *event.Bus,
	c component.Component,
	store Store[T],
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
			return fmt.Errorf("could not publish update event: %w", err)
		}
	}
	return nil
}

func (s *ECS) UpdatePositionComponent(c position.Position) error {
	err := updateComponent(s.eventBus, &c, s.stores.Position, func(p *position.Position) event.Event {
		return position.NewUpdatedEvent(p)
	})
	if err != nil {
		return err
	}
	// Add the entity to the hierarchy.
	if err = s.hierarchy.Update(c.Parent, c.Entity()); err != nil {
		return fmt.Errorf("could not update entity in hierarchy: %w", err)
	}
	return nil
}

func (s *ECS) UpdateVelocityComponent(c velocity.Velocity) error {
	return updateComponent(s.eventBus, &c, s.stores.Velocity, func(p *velocity.Velocity) event.Event {
		return velocity.NewUpdatedEvent(p)
	})
}

func (s *ECS) UpdateControllableComponent(c controllable.Controllable) error {
	return updateComponent(s.eventBus, &c, s.stores.Controllable, func(ctr *controllable.Controllable) event.Event {
		return controllable.NewUpdatedEvent(ctr)
	})
}

func (s *ECS) UpdateRectangleComponent(c shape.Rectangle) error {
	return updateComponent(s.eventBus, &c, s.stores.Rectangle, nil)
}

func (s *ECS) UpdateSpriteComponent(c sprite.Sprite) error {
	return updateComponent(s.eventBus, &c, s.stores.Sprite, nil)
}

func (s *ECS) UpdateSkeletonComponent(c skeleton.Skeleton) error {
	return updateComponent(s.eventBus, &c, s.stores.Skeleton, func(sk *skeleton.Skeleton) event.Event {
		return skeleton.NewUpdatedEvent(sk)
	})
}
