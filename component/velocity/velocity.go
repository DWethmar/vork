package velocity

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/point"
)

const Type = component.Type("velocity")

var _ component.Component = &Velocity{}

// Velocity is a component that holds the velocity of an entity.
type Velocity struct {
	point.Point
	I uint          // ID
	E entity.Entity // Entity
}

func New(e entity.Entity, v point.Point) *Velocity {
	return &Velocity{
		I:     0,
		E:     e,
		Point: v,
	}
}

// Cords returns the x and y speed of the Velocity.
func (p *Velocity) Cords() (int, int)     { return p.X, p.Y }
func (p *Velocity) SetCords(x, y int)     { p.X, p.Y = x, y }
func (p *Velocity) ID() uint              { return p.I }
func (p *Velocity) SetID(i uint)          { p.I = i }
func (p *Velocity) Type() component.Type  { return Type }
func (p *Velocity) Entity() entity.Entity { return p.E }

func Empty() *Velocity {
	return &Velocity{}
}

func init() {
	gob.Register(Velocity{})
}
