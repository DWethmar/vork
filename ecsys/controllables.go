package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Controllable(e entity.Entity) (controllable.Controllable, error) {
	c, err := s.controllableStore.FirstByEntity(e)
	if err != nil {
		return controllable.Controllable{}, fmt.Errorf("could not get controllable: %w", err)
	}
	return *c, nil
}
func (s *ECS) Controllables() []controllable.Controllable {
	c := s.controllableStore.List()
	r := make([]controllable.Controllable, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}

func (s *ECS) DeleteControllable(id uint) error { return s.controllableStore.Delete(id) }
