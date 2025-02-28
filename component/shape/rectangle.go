package shape

import (
	"fmt"
	"image/color"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const RectangleType = component.Type("shape-rectangle")

var _ component.Component = &Rectangle{}

// Shape is a component that holds the shape of an entity.
type Rectangle struct {
	I             uint          // ID
	E             entity.Entity // Entity
	Width, Height int64
	Color         color.RGBA
}

func (p *Rectangle) ID() uint              { return p.I }
func (p *Rectangle) SetID(i uint)          { p.I = i }
func (p *Rectangle) Type() component.Type  { return RectangleType }
func (p *Rectangle) Entity() entity.Entity { return p.E }

func NewRectangle(e entity.Entity, width, height int64, color color.RGBA) *Rectangle {
	return &Rectangle{
		I:      0,
		E:      e,
		Width:  width,
		Height: height,
		Color:  color,
	}
}

type RectangleStore struct {
	cs     *component.Store[*Rectangle]
	nextID uint
}

func NewRectangleStore() *RectangleStore {
	return &RectangleStore{
		cs:     component.NewStore[*Rectangle](),
		nextID: 1,
	}
}

func (s *RectangleStore) Add(c Rectangle) (uint, error) {
	c.I = s.nextID
	s.nextID++
	id, err := s.cs.Add(&c)
	if err != nil {
		return 0, fmt.Errorf("failed to add rectangle component: %w", err)
	}
	return id, nil
}

func (s *RectangleStore) Get(id uint) (*Rectangle, error) {
	return s.cs.Get(id)
}

func (s *RectangleStore) Update(c Rectangle) error {
	if err := s.cs.Update(&c); err != nil {
		return fmt.Errorf("failed to update rectangle component: %w", err)
	}
	return nil
}

func (s *RectangleStore) Delete(id uint) error {
	if err := s.cs.Delete(id); err != nil {
		return fmt.Errorf("failed to remove rectangle component: %w", err)
	}
	return nil
}

func (s *RectangleStore) All() []*Rectangle {
	return s.cs.All()
}
