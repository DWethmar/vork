// package entity is a package that holds the entity type.
package entity

// Entity is a type that represents an entity.
type Entity uint

type Store struct {
	nextID uint
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) CreateEntity() Entity {
	s.nextID++
	return Entity(s.nextID)
}
