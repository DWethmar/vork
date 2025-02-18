package velocity

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/point"
)

const Type = component.Type("velocity")

var _ component.Component = &Velocity{}

// Velocity is a component that holds the velocity of an entity.
type Velocity struct {
	point.Point
	I uint          // ID
	E entity.Entity // Entity
}

func New(e entity.Entity, v point.Point) *Velocity {
	return &Velocity{
		I:     0,
		E:     e,
		Point: v,
	}
}

// Cords returns the x and y speed of the Velocity.
func (p *Velocity) Cords() (int, int)     { return p.X, p.Y }
func (p *Velocity) SetCords(x, y int)     { p.X, p.Y = x, y }
func (p *Velocity) ID() uint              { return p.I }
func (p *Velocity) SetID(i uint)          { p.I = i }
func (p *Velocity) Type() component.Type  { return Type }
func (p *Velocity) Entity() entity.Entity { return p.E }

func Empty() *Velocity {
	return &Velocity{}
}

func init() {
	gob.Register(Velocity{})
}

type Store struct {
	eventBus *event.Bus
	cs       *component.Store[*Velocity]
	nextID   uint // nextID is the next ID that will be used.
}

func NewStore(eventBus *event.Bus) *Store {
	return &Store{
		eventBus: eventBus,
		cs:       component.NewStore[*Velocity](),
	}
}

func (s *Store) Add(c Velocity) (uint, error) {
	c.I = s.nextID
	s.nextID++
	id, err := s.cs.Add(&c)
	if err != nil {
		return 0, err
	}
	if err := s.eventBus.Publish(NewCreatedEvent(c)); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Store) Get(id uint) (*Velocity, error) {
	return s.cs.Get(id)
}

func (s *Store) Update(c Velocity) error {
	if err := s.cs.Update(&c); err != nil {
		return err
	}
	if err := s.eventBus.Publish(NewUpdatedEvent(c)); err != nil {
		return err
	}
	return nil
}

func (s *Store) Delete(id uint) error {
	c, err := s.Get(id)
	if err != nil {
		return err
	}
	if err := s.cs.Delete(id); err != nil {
		return err
	}
	if err := s.eventBus.Publish(NewDeletedEvent(*c)); err != nil {
		return err
	}
	return nil
}

func (s *Store) List(e entity.Entity) []*Velocity {
	return s.cs.List(e)
}

func (s *Store) All() []*Velocity {
	return s.cs.All()
}
