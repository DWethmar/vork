package position

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/point"
)

const Type = component.Type("position")

var _ component.Component = &Position{}

// Position is a component that holds the position of an entity.
type Position struct {
	I      uint          // ID
	E      entity.Entity // Entity
	Parent entity.Entity
	Point  point.Point
}

func New(parent entity.Entity, e entity.Entity, point point.Point) *Position {
	return &Position{
		I:      0,
		E:      e,
		Parent: parent,
		Point:  point,
	}
}

// Cords returns the x and y coordinates of the position.
func (p *Position) Cords() (int, int) { return p.Point.X, p.Point.Y }
func (p *Position) SetCords(x, y int) { p.Point.X, p.Point.Y = x, y }

func Empty() *Position {
	return &Position{}
}

func (p *Position) ID() uint              { return p.I }
func (p *Position) SetID(i uint)          { p.I = i }
func (p *Position) Type() component.Type  { return Type }
func (p *Position) Entity() entity.Entity { return p.E }

func init() {
	gob.Register(Position{})
}
