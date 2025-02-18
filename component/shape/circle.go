package shape

import (
	"fmt"
	"image/color"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const CircleType = component.Type("shape-circle")

var _ component.Component = &Rectangle{}

// Shape is a component that holds the shape of an entity.
type Circle struct {
	I      uint          // ID
	E      entity.Entity // Entity
	Radius int64
	Color  color.RGBA
}

func (p *Circle) ID() uint              { return p.I }
func (p *Circle) SetID(i uint)          { p.I = i }
func (p *Circle) Type() component.Type  { return CircleType }
func (p *Circle) Entity() entity.Entity { return p.E }

func NewCircle(e entity.Entity, radius int64, color color.RGBA) *Circle {
	return &Circle{
		I:      0,
		E:      e,
		Radius: radius,
		Color:  color,
	}
}

type CircleStore struct {
	cs     *component.Store[*Circle]
	nextID uint // Next ID to use
}

func NewCircleStore() *CircleStore {
	return &CircleStore{
		cs: component.NewStore[*Circle](),
	}
}

func (s *CircleStore) Add(c Circle) (uint, error) {
	c.I = s.nextID
	s.nextID++
	id, err := s.cs.Add(&c)
	if err != nil {
		return 0, fmt.Errorf("failed to add circle component: %w", err)
	}
	return id, nil
}

func (s *CircleStore) Get(id uint) (*Circle, error) {
	return s.cs.Get(id)
}

func (s *CircleStore) Update(c Circle) error {
	if err := s.cs.Update(&c); err != nil {
		return fmt.Errorf("failed to update circle component: %w", err)
	}
	return nil
}

func (s *CircleStore) Delete(id uint) error {
	return s.cs.Delete(id)
}

func (s *CircleStore) All() []*Circle {
	return s.cs.All()
}
