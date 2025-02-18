package controllable

import (
	"encoding/gob"
	"fmt"

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

func init() {
	gob.Register(Controllable{})
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

// Store is a store for controllable components.
type Store struct {
	eventBus *event.Bus
	cs       *component.Store[*Controllable]
	nextID   uint // Next ID to use
}

func (s *Store) Add(c *Controllable) (uint, error) {
	c.I = s.nextID
	s.nextID++
	id, err := s.cs.Add(c)
	if err != nil {
		return 0, fmt.Errorf("failed to add controllable component: %w", err)
	}
	if err := s.eventBus.Publish(NewCreatedEvent(*c)); err != nil {
		return 0, fmt.Errorf("failed to publish created event: %w", err)
	}
	return id, nil
}

func (s *Store) Get(id uint) (*Controllable, error) {
	return s.cs.Get(id)
}

func (s *Store) Update(c Controllable) error {
	if err := s.cs.Update(&c); err != nil {
		return fmt.Errorf("failed to update controllable component: %w", err)
	}
	if err := s.eventBus.Publish(NewUpdatedEvent(c)); err != nil {
		return fmt.Errorf("failed to publish updated event: %w", err)
	}
	return nil
}

func (s *Store) Delete(id uint) error {
	c, err := s.Get(id)
	if err != nil {
		return fmt.Errorf("failed to get controllable component: %w", err)
	}
	if err := s.cs.Delete(id); err != nil {
		return fmt.Errorf("failed to delete controllable component: %w", err)
	}
	if err := s.eventBus.Publish(NewDeletedEvent(*c)); err != nil {
		return fmt.Errorf("failed to publish deleted event: %w", err)
	}
	return nil
}

func (s *Store) List(e entity.Entity) []*Controllable {
	return s.cs.List(e)
}

func (s *Store) All() []*Controllable {
	return s.cs.All()
}

// NewStore creates a new store for controllable components.
func NewStore(eventBus *event.Bus) *Store {
	return &Store{
		eventBus: eventBus,
		cs:       component.NewStore[*Controllable](),
	}
}
