package position

import "github.com/dwethmar/vork/entity"

const Type = "position"

var _ entity.Component = &Position{}

// Position is a component that holds the position of an entity.
type Position struct {
	I    uint32
	E    entity.Entity
	X, Y int64
}

func (p *Position) ID() uint32            { return p.I }
func (p *Position) SetID(i uint32)        { p.I = i }
func (p *Position) Type() string          { return Type }
func (p *Position) Entity() entity.Entity { return p.E }
