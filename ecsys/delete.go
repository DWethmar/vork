package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

func (s *ECS) deletePositionByEntity(e entity.Entity) error {
	c, err := s.GetPosition(e)
	if err != nil {
		return err
	}
	return s.DeletePosition(c)
}

func (s *ECS) deleteControllableByEntity(e entity.Entity) error {
	c, err := s.GetControllable(e)
	if err != nil {
		return err
	}
	return s.DeleteControllable(c)
}

func (s *ECS) deleteRectanglesByEntity(e entity.Entity) error {
	for _, c := range s.ListRectanglesByEntity(e) {
		if err := s.DeleteRectangle(c); err != nil {
			return err
		}
	}
	return nil
}

func (s *ECS) deleteSpritesByEntity(e entity.Entity) error {
	for _, sprite := range s.ListSpritesByEntity(e) {
		if err := s.DeleteSprite(sprite); err != nil {
			return err
		}
	}
	return nil
}

func (s *ECS) deleteSkeletonByEntity(e entity.Entity) error {
	c, err := s.GetSkeleton(e)
	if err != nil {
		return err
	}
	return s.DeleteSkeleton(c)
}

func deleteComponent[T component.Component](
	eventBus *event.Bus,
	c component.Component,
	store Store[T],
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
			return fmt.Errorf("could not publish delete event: %w", err)
		}
	}
	return nil
}

func (s *ECS) DeletePosition(c position.Position) error {
	err := deleteComponent(s.eventBus, &c, s.stores.Position, func(p *position.Position) event.Event {
		return position.NewDeletedEvent(*p)
	})
	if err != nil {
		return err
	}
	// Remove the entity from the hierarchy.
	for _, descendant := range s.hierarchy.Delete(c.Entity()) {
		if descendant == c.Entity() {
			continue
		}
		pos, pErr := s.GetPosition(descendant)
		if pErr != nil {
			return fmt.Errorf("could not get position: %w", err)
		}
		if pErr = s.DeletePosition(pos); pErr != nil {
			return fmt.Errorf("could not delete position: %w", err)
		}
	}
	return nil
}

func (s *ECS) DeleteVelocity(c velocity.Velocity) error {
	return deleteComponent(s.eventBus, &c, s.stores.Velocity, func(v *velocity.Velocity) event.Event {
		return velocity.NewDeletedEvent(*v)
	})
}

func (s *ECS) DeleteHitbox(c hitbox.Hitbox) error {
	return deleteComponent(s.eventBus, &c, s.stores.Hitbox, func(h *hitbox.Hitbox) event.Event {
		return hitbox.NewDeletedEvent(*h)
	})
}

func (s *ECS) DeleteControllable(c controllable.Controllable) error {
	return deleteComponent(s.eventBus, &c, s.stores.Controllable, func(ctr *controllable.Controllable) event.Event {
		return controllable.NewDeletedEvent(*ctr)
	})
}

func (s *ECS) DeleteRectangle(c shape.Rectangle) error {
	return deleteComponent(s.eventBus, &c, s.stores.Rectangle, nil)
}

func (s *ECS) DeleteSprite(c sprite.Sprite) error {
	return deleteComponent(s.eventBus, &c, s.stores.Sprite, nil)
}

func (s *ECS) DeleteSkeleton(c skeleton.Skeleton) error {
	return deleteComponent(s.eventBus, &c, s.stores.Skeleton, func(sk *skeleton.Skeleton) event.Event {
		return skeleton.NewDeletedEvent(*sk)
	})
}
