package mock

import "github.com/dwethmar/vork/component"

type Repository struct {
	SaveFunc   func(c component.Component) error
	DeleteFunc func(t component.ComponentType, id uint32) error
	ListFunc   func(t string) []component.Component
}

func (r *Repository) Save(c component.Component) error {
	return r.SaveFunc(c)
}

func (r *Repository) Delete(t component.ComponentType, id uint32) error {
	return r.DeleteFunc(t, id)
}

func (r *Repository) List(t string) []component.Component {
	return r.ListFunc(t)
}
