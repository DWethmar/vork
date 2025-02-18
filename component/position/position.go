package position

import (
	"fmt"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/point"
)

const Type = component.Type("position")
const Root = entity.Entity(0)

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

type Store struct {
	eventBus  *event.Bus
	cs        *component.Store[*Position]
	nextID    uint // nextID is the next ID that will be used.
	hierarchy *Hierarchy
}

func NewStore(eventBus *event.Bus) *Store {
	return &Store{
		eventBus:  eventBus,
		cs:        component.NewStore[*Position](),
		nextID:    0,
		hierarchy: NewHierarchy(Root),
	}
}

func (s *Store) Add(c Position) (uint, error) {
	c.I = s.nextID
	s.nextID++
	id, err := s.cs.Add(&c)
	if err != nil {
		return 0, err
	}
	// Add the entity to the hierarchy.
	if err = s.hierarchy.Add(c.Parent, c.Entity()); err != nil {
		return 0, fmt.Errorf("could not add entity to hierarchy: %w", err)
	}
	if err := s.eventBus.Publish(NewCreatedEvent(c)); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Store) Get(id uint) (*Position, error) {
	return s.cs.Get(id)
}

func (s *Store) Update(c Position) error {
	if err := s.cs.Update(&c); err != nil {
		return fmt.Errorf("could not update component: %w", err)
	}
	if err := s.hierarchy.Update(c.Parent, c.Entity()); err != nil {
		return fmt.Errorf("could not update entity in hierarchy: %w", err)
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
		return fmt.Errorf("could not delete component: %w", err)
	}

	// Remove the entity from the hierarchy. // todo: this should be done in the component store.
	s.hierarchy.Delete(c.Entity())

	if err := s.eventBus.Publish(NewDeletedEvent(*c)); err != nil {
		return err
	}
	return nil
}

func (s *Store) List(e entity.Entity) []*Position {
	return s.cs.List(e)
}

func (s *Store) AbsolutePosition(e entity.Entity) (point.Point, error) {
	if e == s.hierarchy.Root() {
		return point.Point{
			X: 0,
			Y: 0,
		}, nil
	}
	parent, err := s.hierarchy.Parent(e)
	if err != nil {
		return point.Point{}, err
	}
	pos, err := s.cs.First(e)
	if err != nil {
		return point.Point{}, err
	}
	p, err := s.AbsolutePosition(parent)
	if err != nil {
		return point.Point{}, err
	}
	return p.Add(pos.Cords()), nil
}

func (s *Store) All() []*Position {
	return s.cs.All()
}
