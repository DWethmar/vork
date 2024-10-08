package sprite

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

var _ component.Component = &Sprite{}

type Graphic string

const Type = component.ComponentType("sprite")

// Sprite is a component that holds the sprite of an entity.
type Sprite struct {
	I       uint
	E       entity.Entity
	Graphic Graphic
}

func (p *Sprite) ID() uint                      { return p.I }
func (p *Sprite) SetID(i uint)                  { p.I = i }
func (p *Sprite) Type() component.ComponentType { return Type }
func (p *Sprite) Entity() entity.Entity         { return p.E }

func New(e entity.Entity, graphic Graphic) *Sprite {
	return &Sprite{
		I:       0,
		E:       e,
		Graphic: graphic,
	}
}
