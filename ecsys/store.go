package ecsys

import (
	"errors"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/entity"
)

var (
	// ErrComponentNotFound is returned when a component is not found.
	ErrComponentNotFound = errors.New("component not found")
	// ErrEntityNotFound is returned when an entity is not found in the store.
	ErrEntityNotFound = errors.New("entity not found")
	// ErrUniqueComponentViolation is returned when a component is unique per entity.
	ErrUniqueComponentViolation = errors.New("only one component per entity allowed")
)

// Store defines a generic interface for managing any component type.
// T is the component type, such as position, sprite, etc.
type Store[T component.Component] interface {
	Add(T) (uint, error)           // Add a new component to the store.
	Get(uint) (T, error)           // Get a component by its ID.
	Update(T) error                // Update an existing component.
	List() []T                     // List all components in the store.
	Delete(uint) error             // Delete a component by its ID.
	DeleteAll(entity.Entity) error // Delete all components associated with an entity.
}

// ControllableStore manages Controllable components
// Includes an additional method to get the first Controllable by an entity.
type ControllableStore interface {
	Store[*controllable.Controllable]
	First(entity.Entity) (*controllable.Controllable, error)
}

// PositionStore manages Position components
// Includes an additional method to get the first Position by an entity.
type PositionStore interface {
	Store[*position.Position]
	First(entity.Entity) (*position.Position, error)
}

// VelocityStore manages Velocity components.
type VelocityStore interface {
	Store[*velocity.Velocity]
	First(entity.Entity) (*velocity.Velocity, error)
}

// HitboxStore manages Velocity components.
type HitboxStore interface {
	Store[*hitbox.Hitbox]
	ListByEntity(entity.Entity) []*hitbox.Hitbox
}

// RectanglesStore manages Rectangle components (for shapes)
// Includes an additional method to get the first Rectangle by an entity.
type RectanglesStore interface {
	Store[*shape.Rectangle]
	ListByEntity(entity.Entity) []*shape.Rectangle
}

// SpriteStore manages Sprite components
// Includes an additional method to list all sprites associated with an entity.
type SpriteStore interface {
	Store[*sprite.Sprite]
	ListByEntity(entity.Entity) []*sprite.Sprite
}

// SkeletonStore manages Skeleton components
// Includes an additional method to get the first Skeleton by an entity.
type SkeletonStore interface {
	Store[*skeleton.Skeleton]
	First(entity.Entity) (*skeleton.Skeleton, error)
}

// Stores is a collection of component stores used in the ECS.
type Stores struct {
	Controllable ControllableStore
	Position     PositionStore
	Velocity     VelocityStore
	Hitbox       HitboxStore
	Rectangle    RectanglesStore
	Sprite       SpriteStore
	Skeleton     SkeletonStore
}

// NewStores creates a new set of component stores.
func NewStores() *Stores {
	return &Stores{
		Controllable: NewMemStore[*controllable.Controllable](true),
		Position:     NewMemStore[*position.Position](true),
		Velocity:     NewMemStore[*velocity.Velocity](true),
		Hitbox:       NewMemStore[*hitbox.Hitbox](false),
		Rectangle:    NewMemStore[*shape.Rectangle](true),
		Sprite:       NewMemStore[*sprite.Sprite](false),
		Skeleton:     NewMemStore[*skeleton.Skeleton](true),
	}
}
