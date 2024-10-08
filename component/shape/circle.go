package shape

import (
	"image/color"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const CircleType = component.ComponentType("shape-circle")

var _ component.Component = &Rectangle{}

// Shape is a component that holds the shape of an entity.
type Circle struct {
	I      uint          // ID
	E      entity.Entity // Entity
	Radius int64
	Color  color.RGBA
}

func (p *Circle) ID() uint                      { return p.I }
func (p *Circle) SetID(i uint)                  { p.I = i }
func (p *Circle) Type() component.ComponentType { return CircleType }
func (p *Circle) Entity() entity.Entity         { return p.E }

func NewCircle(e entity.Entity, radius int64, color color.RGBA) *Circle {
	return &Circle{
		I:      0,
		E:      e,
		Radius: radius,
		Color:  color,
	}
}
