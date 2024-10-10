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
	"github.com/dwethmar/vork/component/store"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

// componentTypes is a list of all component types used in the ECS.
// This list is used to initialize the component stores in the ECS.
// It is also used to ensure that all component types are accounted for when managing entities.
func componentTypes() []component.Type {
	return []component.Type{
		position.Type,
		controllable.Type,
		shape.RectangleType,
		// shape.CircleType, // not implemented yet
		sprite.Type,
		skeleton.Type,
	}
}

// ECS is the main struct that manages entities and their associated components.
// It also provides access to various component stores (position, controllable, rectangle, sprite, skeleton)
// and integrates an event bus for handling in-game events.
type ECS struct {
	mu sync.RWMutex
	// lastEntityID is the last entity ID that was created. It is used to generate new entity IDs.
	// When adding a component with an entity ID higher than lastEntityID, lastEntityID is updated.
	lastEntityID entity.Entity
	eventBus     *event.Bus
	stores       *store.Stores
}

// New creates a new ECS system, initializing it with the provided component stores and event bus.
// This function ensures that the ECS is ready to manage entities and components from the start.
func New(eventBus *event.Bus, s *store.Stores) *ECS {
	return &ECS{
		lastEntityID: 0,
		eventBus:     eventBus,
		stores:       s,
	}
}

// CreateEntity generates a new unique entity by incrementing the lastEntityID.
// It also creates a position component for the entity and adds it to the ECS.
func (s *ECS) CreateEntity(x, y int) (entity.Entity, error) {
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
	stores := s.stores
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range componentTypes() {
		var err error
		switch t {
		case position.Type:
			err = stores.Position.DeleteByEntity(e)
		case controllable.Type:
			err = stores.Controllable.DeleteByEntity(e)
		case shape.RectangleType:
			err = stores.Rectangle.DeleteByEntity(e)
		case sprite.Type:
			err = stores.Sprite.DeleteByEntity(e)
		case skeleton.Type:
			err = stores.Skeleton.DeleteByEntity(e)
		default:
			return fmt.Errorf("failed to delete entity because of an unknown component type %s", t)
		}
		if err != nil && !errors.Is(err, store.ErrEntityNotFound) {
			return fmt.Errorf("failed to delete entity: %w", err)
		}
	}

	return nil
}
