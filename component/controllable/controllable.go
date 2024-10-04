package controllable

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const Type = component.ComponentType("controllable")

var _ component.Component = &Controllable{}

// Controllable is a component that holds the controller of an entity.
type Controllable struct {
	*component.BaseComponent
}

func New(e entity.Entity) *Controllable {
	return &Controllable{
		BaseComponent: &component.BaseComponent{
			E: e,
			T: Type,
		},
	}
}

func init() {
	gob.Register(Controllable{})
}
