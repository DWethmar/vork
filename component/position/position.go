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
	I    uint          // ID
	E    entity.Entity // Entity
	X, Y int
}

func New(e entity.Entity, x, y int) *Position {
	return &Position{
		I: 0,
		E: e,
		X: x,
		Y: y,
	}
}

func (p *Position) ID() uint                      { return p.I }
func (p *Position) SetID(i uint)                  { p.I = i }
func (p *Position) Type() component.ComponentType { return Type }
func (p *Position) Entity() entity.Entity         { return p.E }

func init() {
	gob.Register(Position{})
}
