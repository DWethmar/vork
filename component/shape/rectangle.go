package shape

import (
	"image/color"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const RectangleType = component.ComponentType("shape-rectangle")

var _ component.Component = &Rectangle{}

// Shape is a component that holds the shape of an entity.
type Rectangle struct {
	I             uint32        // ID
	E             entity.Entity // Entity
	Width, Height int64
	Color         color.RGBA
}

func (p *Rectangle) ID() uint32                    { return p.I }
func (p *Rectangle) SetID(i uint32)                { p.I = i }
func (p *Rectangle) Type() component.ComponentType { return RectangleType }
func (p *Rectangle) Entity() entity.Entity         { return p.E }

func NewRectangle(e entity.Entity, width, height int64, color color.RGBA) *Rectangle {
	return &Rectangle{
		I:      0,
		E:      e,
		Width:  width,
		Height: height,
		Color:  color,
	}
}
