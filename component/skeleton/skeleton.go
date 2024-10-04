package skeleton

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

const Type = component.ComponentType("skeleton")

// Skeleton is a component that describes a skeleton enemy.
type Skeleton struct {
	I uint32        // ID
	E entity.Entity // Entity
}

func New(e entity.Entity) *Skeleton {
	return &Skeleton{
		I: 0,
		E: e,
	}
}

func (s *Skeleton) ID() uint32                    { return s.I }
func (s *Skeleton) SetID(i uint32)                { s.I = i }
func (s *Skeleton) Type() component.ComponentType { return Type }
func (s *Skeleton) Entity() entity.Entity         { return s.E }

func init() {
	gob.Register(Skeleton{})
}
