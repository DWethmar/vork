package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Position(e entity.Entity) (position.Position, error) {
	c, err := s.pos.FirstByEntity(e)
	if err != nil {
		return position.Position{}, fmt.Errorf("could not get position: %v", err)
	}
	return *c, nil
}

func (s *ECS) UpdatePosition(p position.Position) error {
	if err := s.pos.Update(&p); err != nil {
		return fmt.Errorf("could not update position: %v", err)
	}
	if err := s.eventBus.Publish(position.NewUpdatedEvent(p)); err != nil {
		return fmt.Errorf("could not publish position update event: %v", err)
	}
	return nil
}

func (s *ECS) AddPosition(p position.Position) (uint32, error) {
	id, err := s.pos.Add(&p)
	if err != nil {
		return 0, fmt.Errorf("could not add position: %v", err)
	}
	if err := s.eventBus.Publish(position.NewCreatedEvent(p)); err != nil {
		return 0, fmt.Errorf("could not publish position create event: %v", err)
	}
	return id, nil
}

func (s *ECS) DeletePosition(p position.Position) error {
	if err := s.pos.Delete(p.ID()); err != nil {
		return fmt.Errorf("could not delete position: %v", err)
	}
	if err := s.eventBus.Publish(position.NewDeletedEvent(p)); err != nil {
		return fmt.Errorf("could not publish position delete event: %v", err)
	}
	return nil
}
