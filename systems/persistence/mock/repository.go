package mock

import "github.com/dwethmar/vork/component"

type Repository[T component.Component] struct {
	GetFunc    func(id uint) (T, error)
	SaveFunc   func(c T) error
	DeleteFunc func(id uint) error
	ListFunc   func() ([]T, error)
}

func (r *Repository[T]) Get(id uint) (T, error) {
	return r.GetFunc(id)
}

func (r *Repository[T]) Save(c T) error {
	return r.SaveFunc(c)
}

func (r *Repository[T]) Delete(id uint) error {
	return r.DeleteFunc(id)
}

func (r *Repository[T]) List() ([]T, error) {
	return r.ListFunc()
}
