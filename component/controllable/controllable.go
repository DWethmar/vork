package controllable

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

const Type = component.Type("controllable")

var _ component.Component = &Controllable{}

// Controllable is a component that holds the controller of an entity.
type Controllable struct {
	I uint          // ID
	E entity.Entity // Entity
}

func New(e entity.Entity) *Controllable {
	return &Controllable{
		I: 0,
		E: e,
	}
}

func Empty() *Controllable {
	return &Controllable{}
}

func (c *Controllable) ID() uint              { return c.I }
func (c *Controllable) SetID(i uint)          { c.I = i }
func (c *Controllable) Type() component.Type  { return Type }
func (c *Controllable) Entity() entity.Entity { return c.E }

// NewStore creates a new store for controllable components.
func NewStore(eventBus *event.Bus) *component.Store[*Controllable] {
	return component.NewStore[*Controllable](
		true,
		func(c *Controllable) error {
			return eventBus.Publish(NewCreatedEvent(*c))
		},
		func(c *Controllable) error {
			return eventBus.Publish(NewUpdatedEvent(*c))
		},
		func(c *Controllable) error {
			return eventBus.Publish(NewDeletedEvent(*c))
		},
	)
}

func init() {
	gob.Register(Controllable{})
}
