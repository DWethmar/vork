package shape

import "github.com/dwethmar/vork/entity"

const Type = "shape"

// Shape is a component that holds the shape of an entity.
type Rectangle struct {
	I             uint32
	E             entity.Entity
	Width, Height int64
}

func (p *Rectangle) ID() uint32            { return p.I }
func (p *Rectangle) SetID(i uint32)        { p.I = i }
func (p *Rectangle) Type() string          { return Type }
func (p *Rectangle) Entity() entity.Entity { return p.E }
