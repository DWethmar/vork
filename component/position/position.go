package position

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const Type = component.ComponentType("position")

var _ component.Component = &Position{}

// Position is a component that holds the position of an entity.
type Position struct {
	*component.BaseComponent
	X, Y int64
}

func New(e entity.Entity, x, y int64) *Position {
	return &Position{
		BaseComponent: &component.BaseComponent{
			E: e,
			T: Type,
		},
		X: x,
		Y: y,
	}
}

func init() {
	gob.Register(Position{})
}
