package hitbox

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/point"
)

const (
	Type = component.Type("hitbox")
)

// Hitbox is a component that holds the hitbox of an entity.
type Hitbox struct {
	I      uint          // ID
	E      entity.Entity // Entity
	Tag    string        // Tag is a string that describes the hitbox.
	Width  int
	Height int
	Offset point.Point
}

// New creates a new hitbox component.
func New(e entity.Entity, tag string, width, height int, offset point.Point) *Hitbox {
	return &Hitbox{
		I:      0,
		E:      e,
		Tag:    tag,
		Width:  width,
		Height: height,
		Offset: offset,
	}
}

func (h *Hitbox) SetID(i uint)          { h.I = i }
func (h *Hitbox) ID() uint              { return h.I }
func (h *Hitbox) Type() component.Type  { return Type }
func (h *Hitbox) Entity() entity.Entity { return h.E }

type Store struct {
	eventBus *event.Bus
	cs       *component.Store[*Hitbox]
	nextID   uint // nextID is the next ID that will be used.
}

func NewStore(eventBus *event.Bus) *Store {
	return &Store{
		eventBus: eventBus,
		cs:       component.NewStore[*Hitbox](),
	}
}

func (s *Store) Add(c Hitbox) (uint, error) {
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

func (s *Store) Get(id uint) (*Hitbox, error) {
	return s.cs.Get(id)
}

func (s *Store) Update(c Hitbox) error {
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
