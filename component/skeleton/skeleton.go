package skeleton

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/direction"
	"github.com/dwethmar/vork/entity"
)

const Type = component.ComponentType("skeleton")

type State int

const (
	Idle State = iota
	Moving
)

// Skeleton is a component that describes a skeleton enemy.
type Skeleton struct {
	I             uint          // ID
	E             entity.Entity // Entity
	State         State
	PrefX, PrefY  int // Preferred X and Y
	Facing        direction.Direction
	AnimationStep uint8
}

func New(e entity.Entity) *Skeleton {
	return &Skeleton{
		I:             0,
		E:             e,
		State:         Idle,
		Facing:        direction.South,
		AnimationStep: 0,
	}
}

func (s *Skeleton) ID() uint                      { return s.I }
func (s *Skeleton) SetID(i uint)                  { s.I = i }
func (s *Skeleton) Type() component.ComponentType { return Type }
func (s *Skeleton) Entity() entity.Entity         { return s.E }

func init() {
	gob.Register(Skeleton{})
}
