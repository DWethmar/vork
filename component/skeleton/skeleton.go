package skeleton

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const Type = component.ComponentType("skeleton")

// Skeleton is a component that describes a skeleton enemy.
type Skeleton struct {
	*component.BaseComponent
	PrevX, PrevY int64
}

func New(e entity.Entity) *Skeleton {
	return &Skeleton{
		BaseComponent: &component.BaseComponent{
			E: e,
			T: Type,
		},
	}
}
