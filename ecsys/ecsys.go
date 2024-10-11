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
func (s *ECS) CreateEntity(parent entity.Entity, x, y int) (entity.Entity, error) {
	e := s.CreateEmptyEntity()
	pos := position.New(parent, e, x, y)
	if _, err := s.AddPositionComponent(*pos); err != nil {
		return 0, err
	}
	return e, nil
}

// CreateEmptyEntity generates a new unique entity by incrementing the lastEntityID.
func (s *ECS) CreateEmptyEntity() entity.Entity {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastEntityID++
	return s.lastEntityID
}

// DeleteEntity removes an entity and all its associated components from the ECS.
func (s *ECS) DeleteEntity(e entity.Entity) error {
	// Check for errors and return the first one that is not a "not found" error.
	for _, err := range s.deleteEntity(e) {
		if err != nil && !errors.Is(err, store.ErrEntityNotFound) {
			return fmt.Errorf("failed to delete entity: %w", err)
		}
	}
	return nil
}

// deleteEntity removes an entity and all its associated components from the ECS.
func (s *ECS) deleteEntity(e entity.Entity) []error {
	errs := make([]error, 0)
	for _, t := range componentTypes() {
		switch t {
		case position.Type:
			if err := s.deletePosition(e); err != nil {
				errs = append(errs, err)
			}
		case controllable.Type:
			if err := s.deleteControllable(e); err != nil {
				errs = append(errs, err)
			}
		case shape.RectangleType:
			if err := s.deleteRectangle(e); err != nil {
				errs = append(errs, err)
			}
		case sprite.Type:
			if err := s.deleteSprite(e); err != nil {
				errs = append(errs, err)
			}
		case skeleton.Type:
			if err := s.deleteSkeleton(e); err != nil {
				errs = append(errs, err)
			}
		default:
			errs = append(errs, fmt.Errorf("unknown component type: %v", t))
		}
	}

	return errs
}

func (s *ECS) deletePosition(e entity.Entity) error {
	c, err := s.GetPosition(e)
	if err != nil {
		return err
	}
	return s.DeletePosition(c)
}

func (s *ECS) deleteControllable(e entity.Entity) error {
	c, err := s.GetControllable(e)
	if err != nil {
		return err
	}
	return s.DeleteControllable(c)
}

func (s *ECS) deleteRectangle(e entity.Entity) error {
	for _, c := range s.ListRectanglesByEntity(e) {
		if err := s.DeleteRectangle(c); err != nil {
			return err
		}
	}
	return nil
}

func (s *ECS) deleteSprite(e entity.Entity) error {
	for _, sprite := range s.ListSpritesByEntity(e) {
		if err := s.DeleteSprite(sprite); err != nil {
			return err
		}
	}
	return nil
}

func (s *ECS) deleteSkeleton(e entity.Entity) error {
	c, err := s.GetSkeleton(e)
	if err != nil {
		return err
	}
	return s.DeleteSkeleton(c)
}
