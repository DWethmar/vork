package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Position(e entity.Entity) (position.Position, error) {
	return s.pos.FirstByEntity(e)
}

func (s *ECS) UpdatePosition(p position.Position) error {
	if err := s.pos.Update(p); err != nil {
		return fmt.Errorf("could not update position: %v", err)
	}
	s.eventBus.Publish(&position.UpdatedEvent{Position: p})
	return nil
}

func (s *ECS) AddPosition(p position.Position) (uint32, error) {
	id, err := s.pos.Add(p)
	if err != nil {
		return 0, fmt.Errorf("could not add position: %v", err)
	}
	s.eventBus.Publish(&position.CreatedEvent{Position: p})
	return id, nil
}

func (s *ECS) DeletePosition(p position.Position) error {
	if err := s.pos.Delete(p.ID()); err != nil {
		return fmt.Errorf("could not delete position: %v", err)
	}
	s.eventBus.Publish(&position.DeletedEvent{Position: p})
	return nil
}
