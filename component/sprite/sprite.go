package sprite

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

var _ component.Component = &Sprite{}

type Graphic string

const Type = component.Type("sprite")

// Sprite is a component that holds the sprite of an entity.
type Sprite struct {
	I       uint
	E       entity.Entity
	Tag     string  // Tag used to identify the sprite
	Graphic Graphic // Graphic of the sprite
}

func (p *Sprite) ID() uint              { return p.I }
func (p *Sprite) SetID(i uint)          { p.I = i }
func (p *Sprite) Type() component.Type  { return Type }
func (p *Sprite) Entity() entity.Entity { return p.E }

func New(e entity.Entity, tag string, graphic Graphic) *Sprite {
	return &Sprite{
		I:       0,
		E:       e,
		Tag:     tag,
		Graphic: graphic,
	}
}

type Store struct {
	cs     *component.Store[*Sprite]
	nextID uint // nextID is the next ID that will be used.
}

func NewStore() *Store {
	return &Store{
		cs: component.NewStore[*Sprite](),
	}
}

func (s *Store) Add(c Sprite) (uint, error) {
	c.I = s.nextID
	s.nextID++
	id, err := s.cs.Add(&c)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Store) Get(id uint) (*Sprite, error) {
	return s.cs.Get(id)
}

func (s *Store) Update(c Sprite) error {
	if err := s.cs.Update(&c); err != nil {
		return err
	}
	return nil
}

func (s *Store) Delete(id uint) error {
	return s.cs.Delete(id)
}

func (s *Store) All() []*Sprite {
	return s.cs.All()
}
