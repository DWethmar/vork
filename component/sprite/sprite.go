package sprite

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

type Graphic string

const Type = component.ComponentType("sprite")

// Sprite is a component that holds the sprite of an entity.
type Sprite struct {
	*component.BaseComponent
	Graphic Graphic
}

func New(e entity.Entity, graphic Graphic) *Sprite {
	return &Sprite{
		BaseComponent: &component.BaseComponent{
			E: e,
		},
		Graphic: graphic,
	}
}
