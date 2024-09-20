package shape

import (
	"image/color"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const Type = component.ComponentType("shape")

// Shape is a component that holds the shape of an entity.
type Rectangle struct {
	*component.BaseComponent
	Width, Height int64
	Color         color.RGBA
}

func NewRectangle(e entity.Entity, width, height int64, color color.RGBA) *Rectangle {
	return &Rectangle{
		BaseComponent: &component.BaseComponent{
			E: e,
			T: Type,
		},
		Width:  width,
		Height: height,
		Color:  color,
	}
}
