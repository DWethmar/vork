package ecsys

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Controllable(e entity.Entity) (controllable.Controllable, error) {
	return s.contr.FirstByEntity(e)
}
func (s *ECS) Controllables() []controllable.Controllable           { return s.contr.List() }
func (s *ECS) UpdateControllable(c controllable.Controllable) error { return s.contr.Update(c) }
func (s *ECS) AddControllable(c controllable.Controllable) error    { return s.contr.Add(c) }
func (s *ECS) DeleteControllable(id uint32) error                   { return s.contr.Delete(id) }
