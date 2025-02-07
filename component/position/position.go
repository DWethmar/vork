package position

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/point"
)

const Type = component.Type("position")

var _ component.Component = &Position{}

// Position is a component that holds the position of an entity.
type Position struct {
	point.Point
	I      uint          // ID
	E      entity.Entity // Entity
	Parent entity.Entity
}

func New(parent entity.Entity, e entity.Entity, coord point.Point) *Position {
	return &Position{
		I:      0,
		E:      e,
		Parent: parent,
		Point:  coord,
	}
}

// Cords returns the x and y coordinates of the position.
func (p *Position) Cords() (int, int)     { return p.X, p.Y }
func (p *Position) SetCords(x, y int)     { p.X, p.Y = x, y }
func (p *Position) ID() uint              { return p.I }
func (p *Position) SetID(i uint)          { p.I = i }
func (p *Position) Type() component.Type  { return Type }
func (p *Position) Entity() entity.Entity { return p.E }

func Empty() *Position {
	return &Position{}
}

// NewStore creates a new store for position components.
func NewStore(eventBus *event.Bus) *component.Store[*Position] {
	return component.NewStore[*Position](
		true,
		func(c *Position) error {
			return eventBus.Publish(NewCreatedEvent(*c))
		},
		func(c *Position) error {
			return eventBus.Publish(NewUpdatedEvent(*c))
		},
		func(c *Position) error {
			return eventBus.Publish(NewDeletedEvent(*c))
		},
	)
}

func init() {
	gob.Register(Position{})
}
