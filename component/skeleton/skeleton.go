package skeleton

import (
	"encoding/gob"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/direction"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

func init() {
	gob.Register(Skeleton{})
}

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

type Store struct {
	eventBus *event.Bus
	cs       *component.Store[*Skeleton]
	nextID   uint // nextID is the next ID that will be used.
}

func NewStore(eventBus *event.Bus) *Store {
	return &Store{
		eventBus: eventBus,
		cs:       component.NewStore[*Skeleton](),
	}
}

func (s *Store) Add(c Skeleton) (uint, error) {
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

func (s *Store) Get(id uint) (*Skeleton, error) {
	return s.cs.Get(id)
}

func (s *Store) Update(c Skeleton) error {
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

func (s *Store) All() []*Skeleton {
	return s.cs.All()
}
