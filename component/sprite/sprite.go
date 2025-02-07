package sprite

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

var _ component.Component = &Sprite{}

type Graphic string

const Type = component.Type("sprite")

// Sprite is a component that holds the sprite of an entity.
type Sprite struct {
	I       uint
	E       entity.Entity
	Tag     string  // Tag used to identify the sprite
	Graphic Graphic // Graphic of the sprite
}

func (p *Sprite) ID() uint              { return p.I }
func (p *Sprite) SetID(i uint)          { p.I = i }
func (p *Sprite) Type() component.Type  { return Type }
func (p *Sprite) Entity() entity.Entity { return p.E }

func New(e entity.Entity, tag string, graphic Graphic) *Sprite {
	return &Sprite{
		I:       0,
		E:       e,
		Tag:     tag,
		Graphic: graphic,
	}
}

func NewStore() *component.Store[*Sprite] {
	return component.NewStore[*Sprite](true, nil, nil, nil)
}
