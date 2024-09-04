package controllable

import "github.com/dwethmar/vork/entity"

const Type = "controller"

var _ entity.Component = &Controllable{}

// Controllable is a component that holds the controller of an entity.
type Controllable struct {
	I uint32
	E entity.Entity
}

func (c *Controllable) Entity() entity.Entity { return c.E }
func (c *Controllable) ID() uint32            { return c.I }
func (c *Controllable) SetID(i uint32)        { c.I = i }
func (c *Controllable) Type() string          { return Type }
