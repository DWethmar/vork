package skeleton

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/direction"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

const Type = component.Type("skeleton")

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

func Empty() *Skeleton {
	return &Skeleton{}
}

func (s *Skeleton) ID() uint              { return s.I }
func (s *Skeleton) SetID(i uint)          { s.I = i }
func (s *Skeleton) Type() component.Type  { return Type }
func (s *Skeleton) Entity() entity.Entity { return s.E }

func NewStore(eventBus *event.Bus) *component.Store[*Skeleton] {
	return component.NewStore[*Skeleton](
		true,
		func(c *Skeleton) error {
			return eventBus.Publish(NewCreatedEvent(*c))
		},
		func(c *Skeleton) error {
			return eventBus.Publish(NewUpdatedEvent(*c))
		},
		func(c *Skeleton) error {
			return eventBus.Publish(NewDeletedEvent(*c))
		},
	)
}

func init() {
	gob.Register(Skeleton{})
}
