// package ecsys contains the Entity-Component-System architecture.
package ecsys

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

// componentTypes is a list of all component types used in the ECS.
// This list is used to initialize the component stores in the ECS.
// It is also used to ensure that all component types are accounted for when managing entities.
func componentTypes() []component.ComponentType {
	return []component.ComponentType{
		position.Type,
		controllable.Type,
		shape.RectangleType,
		// shape.CircleType, // not implemented yet
		sprite.Type,
		skeleton.Type,
	}
}

// BaseComponentStore defines a generic interface for managing any component type.
// T is the component type, such as position, sprite, etc.
type BaseComponentStore[T any] interface {
	Add(T) (uint32, error)              // Add a new component to the store.
	Get(uint32) (T, error)              // Get a component by its ID.
	Update(T) error                     // Update an existing component.
	List() []T                          // List all components in the store.
	Delete(uint32) error                // Delete a component by its ID.
	DeleteByEntity(entity.Entity) error // Delete all components associated with an entity.
}

// ControllableStore manages Controllable components, extending BaseComponentStore.
// Includes an additional method to get the first Controllable by an entity.
type ControllableStore interface {
	BaseComponentStore[*controllable.Controllable]
	FirstByEntity(entity.Entity) (*controllable.Controllable, error)
}

// PositionStore manages Position components, extending BaseComponentStore.
// Includes an additional method to get the first Position by an entity.
type PositionStore interface {
	BaseComponentStore[*position.Position]
	FirstByEntity(entity.Entity) (*position.Position, error)
}

// RectanglesStore manages Rectangle components (for shapes), extending BaseComponentStore.
// Includes an additional method to get the first Rectangle by an entity.
type RectanglesStore interface {
	BaseComponentStore[*shape.Rectangle]
	FirstByEntity(entity.Entity) (*shape.Rectangle, error)
}

// SpriteStore manages Sprite components, extending BaseComponentStore.
// Includes an additional method to list all sprites associated with an entity.
type SpriteStore interface {
	BaseComponentStore[*sprite.Sprite]
	ListByEntity(entity.Entity) []*sprite.Sprite
}

// SkeletonStore manages Skeleton components, extending BaseComponentStore.
// Includes an additional method to get the first Skeleton by an entity.
type SkeletonStore interface {
	BaseComponentStore[*skeleton.Skeleton]
	FirstByEntity(entity.Entity) (*skeleton.Skeleton, error)
}

// ECS is the main struct that manages entities and their associated components.
// It also provides access to various component stores (position, controllable, rectangle, sprite, skeleton)
// and integrates an event bus for handling in-game events.
type ECS struct {
	mu           sync.RWMutex
	lastEntityID entity.Entity
	eventBus     *event.Bus
	pos          PositionStore
	contr        ControllableStore
	rect         RectanglesStore
	sprites      SpriteStore
	sklt         SkeletonStore
}

// New creates a new ECS system, initializing it with the provided component stores and event bus.
// This function ensures that the ECS is ready to manage entities and components from the start.
func New(eventBus *event.Bus) *ECS {
	ecs := &ECS{
		lastEntityID: 0,
		eventBus:     eventBus,
	}

	// Initialize component stores for the ECS.
	for _, t := range componentTypes() {
		switch t {
		case position.Type:
			ecs.pos = component.NewStore[*position.Position](true)
		case controllable.Type:
			ecs.contr = component.NewStore[*controllable.Controllable](true)
		case shape.RectangleType:
			ecs.rect = component.NewStore[*shape.Rectangle](true)
		case sprite.Type:
			ecs.sprites = component.NewStore[*sprite.Sprite](false)
		case skeleton.Type:
			ecs.sklt = component.NewStore[*skeleton.Skeleton](true)
		default:
			panic(fmt.Sprintf("failed to initialize ECS because of an unknown component type %s", t))
		}
	}

	return ecs
}

// CreateEntity generates a new unique entity by incrementing the lastEntityID.
// It also creates a position component for the entity and adds it to the ECS.
func (s *ECS) CreateEntity(x, y int64) (entity.Entity, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastEntityID++

	// create position component
	pos := position.New(s.lastEntityID, x, y)
	if _, err := s.AddPositionComponent(*pos); err != nil {
		return 0, err
	}

	return s.lastEntityID, nil
}

func (s *ECS) DeleteEntity(e entity.Entity) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range componentTypes() {
		var err error
		switch t {
		case position.Type:
			err = s.pos.DeleteByEntity(e)
		case controllable.Type:
			err = s.contr.DeleteByEntity(e)
		case shape.RectangleType:
			err = s.rect.DeleteByEntity(e)
		case sprite.Type:
			err = s.sprites.DeleteByEntity(e)
		case skeleton.Type:
			err = s.sklt.DeleteByEntity(e)
		default:
			return fmt.Errorf("failed to delete entity because of an unknown component type %s", t)
		}
		if err != nil && !errors.Is(err, component.ErrEntityNotFound) {
			return err
		}
	}

	return nil
}
