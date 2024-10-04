// package ecsys contains the Entity-Component-System architecture.
package ecsys

import (
	"errors"
	"sync"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

// ErrNotFound is the error returned when a component is not found in the store.
var (
	ErrNotFound = errors.New("not found")
)

// BaseComponentStore defines a generic interface for managing any component type.
// T is the component type, such as position, sprite, etc.
type BaseComponentStore[T any] interface {
	Add(T) error                        // Add a new component to the store.
	Get(uint32) (T, error)              // Get a component by its ID.
	Update(T) error                     // Update an existing component.
	List() []T                          // List all components in the store.
	Delete(uint32) error                // Delete a component by its ID.
	DeleteByEntity(entity.Entity) error // Delete all components associated with an entity.
}

// ControllableStore manages Controllable components, extending BaseComponentStore.
// Includes an additional method to get the first Controllable by an entity.
type ControllableStore interface {
	BaseComponentStore[controllable.Controllable]
	FirstByEntity(entity.Entity) (controllable.Controllable, error)
}

// PositionStore manages Position components, extending BaseComponentStore.
// Includes an additional method to get the first Position by an entity.
type PositionStore interface {
	BaseComponentStore[position.Position]
	FirstByEntity(entity.Entity) (position.Position, error)
}

// RectanglesStore manages Rectangle components (for shapes), extending BaseComponentStore.
// Includes an additional method to get the first Rectangle by an entity.
type RectanglesStore interface {
	BaseComponentStore[shape.Rectangle]
	FirstByEntity(entity.Entity) (shape.Rectangle, error)
}

// SpriteStore manages Sprite components, extending BaseComponentStore.
// Includes an additional method to list all sprites associated with an entity.
type SpriteStore interface {
	BaseComponentStore[sprite.Sprite]
	ListByEntity(entity.Entity) ([]sprite.Sprite, error)
}

// SkeletonStore manages Skeleton components, extending BaseComponentStore.
// Includes an additional method to get the first Skeleton by an entity.
type SkeletonStore interface {
	BaseComponentStore[skeleton.Skeleton]
	FirstByEntity(entity.Entity) (skeleton.Skeleton, error)
}

// ECS is the main struct that manages entities and their associated components.
// It also provides access to various component stores (position, controllable, rectangle, sprite, skeleton)
// and integrates an event bus for handling in-game events.
type ECS struct {
	mu           sync.RWMutex  // Mutex to ensure thread-safe entity and component operations.
	lastEntityID entity.Entity // Tracks the last created entity ID to ensure unique entity creation.
	eventBus     *event.Bus    // Event bus to handle in-game events and communication between systems.
	// Component stores for managing different types of components.
	pos     PositionStore
	contr   ControllableStore
	rect    RectanglesStore
	sprites SpriteStore
	sklt    SkeletonStore
}

// CreateEntity generates a new unique entity by incrementing the lastEntityID.
// This method is thread-safe due to the use of a write lock.
func (s *ECS) CreateEntity() entity.Entity {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastEntityID++
	return s.lastEntityID
}

// New creates a new ECS system, initializing it with the provided component stores and event bus.
// This function ensures that the ECS is ready to manage entities and components from the start.
func New(
	eventBus *event.Bus,
	positions PositionStore,
	controllables ControllableStore,
	rectangles RectanglesStore,
	sprites SpriteStore,
	sklt SkeletonStore,
) *ECS {
	return &ECS{
		lastEntityID: 0,
		eventBus:     eventBus,
		pos:          positions,
		contr:        controllables,
		rect:         rectangles,
		sprites:      sprites,
		sklt:         sklt,
	}
}
