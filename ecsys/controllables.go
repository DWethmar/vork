package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Controllable(e entity.Entity) (controllable.Controllable, error) {
	c, err := s.contr.FirstByEntity(e)
	if err != nil {
		return controllable.Controllable{}, fmt.Errorf("could not get controllable: %v", err)
	}
	return *c, nil
}
func (s *ECS) Controllables() []controllable.Controllable {
	c := s.contr.List()
	r := make([]controllable.Controllable, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) UpdateControllable(c controllable.Controllable) error {
	if err := s.contr.Update(&c); err != nil {
		return fmt.Errorf("could not update controllable: %v", err)
	}
	if err := s.eventBus.Publish(&controllable.UpdatedEvent{Controllable: c}); err != nil {
		return fmt.Errorf("could not publish event: %v", err)
	}
	return nil
}

func (s *ECS) AddControllable(c controllable.Controllable) (uint32, error) {
	id, err := s.contr.Add(&c)
	if err != nil {
		return 0, fmt.Errorf("could not add controllable: %v", err)
	}
	if err := s.eventBus.Publish(&controllable.CreatedEvent{Controllable: c}); err != nil {
		return 0, fmt.Errorf("could not publish event: %v", err)
	}
	return id, nil
}

func (s *ECS) DeleteControllable(id uint32) error { return s.contr.Delete(id) }
