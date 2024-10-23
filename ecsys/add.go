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

// addComponent adds a component to the ECS. It returns the ID of the component.
func addComponent[T component.Component](
	ecs *ECS,
	c component.Component,
	store Store[T],
	eventCreator func(T) event.Event,
) (uint, error) {
	comp, ok := c.(T)
	if !ok {
		return 0, fmt.Errorf("expected component of type %T, got %T", *new(T), c)
	}
	id, err := store.Add(comp)
	if err != nil {
		return 0, fmt.Errorf("could not add component of type %T: %w", *new(T), err)
	}
	if eventCreator != nil {
		if err = ecs.eventBus.Publish(eventCreator(comp)); err != nil {
			return 0, fmt.Errorf("could not publish add event: %w", err)
		}
	}
	// Update the lastEntityID if the entity is higher than the current lastEntityID.
	if c.Entity() > ecs.lastEntityID {
		ecs.lastEntityID = c.Entity()
	}
	return id, nil
}

// AddPosition adds a position component to the ECS.
func (s *ECS) AddPosition(c position.Position) (uint, error) {
	id, err := addComponent(s, &c, s.stores.Position, func(p *position.Position) event.Event {
		return position.NewCreatedEvent(*p)
	})
	if err != nil {
		return 0, fmt.Errorf("could not add position component: %w", err)
	}
	// Add the entity to the hierarchy.
	if err = s.hierarchy.Add(c.Parent, c.Entity()); err != nil {
		return 0, fmt.Errorf("could not add entity to hierarchy: %w", err)
	}
	return id, nil
}

func (s *ECS) AddVelocity(c velocity.Velocity) (uint, error) {
	return addComponent(s, &c, s.stores.Velocity, func(v *velocity.Velocity) event.Event {
		return velocity.NewCreatedEvent(*v)
	})
}

func (s *ECS) AddControllable(c controllable.Controllable) (uint, error) {
	return addComponent(s, &c, s.stores.Controllable, func(ctr *controllable.Controllable) event.Event {
		return controllable.NewCreatedEvent(*ctr)
	})
}

func (s *ECS) AddRectangle(c shape.Rectangle) (uint, error) {
	return addComponent(s, &c, s.stores.Rectangle, nil)
}

func (s *ECS) AddSprite(c sprite.Sprite) (uint, error) {
	return addComponent(s, &c, s.stores.Sprite, nil)
}

func (s *ECS) AddSkeleton(c skeleton.Skeleton) (uint, error) {
	return addComponent(s, &c, s.stores.Skeleton, func(sk *skeleton.Skeleton) event.Event {
		return skeleton.NewCreatedEvent(*sk)
	})
}
