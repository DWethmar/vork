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
	"github.com/dwethmar/vork/point"
)

type Hierarchy interface {
	Add(parent entity.Entity, child entity.Entity) error
	Update(parent entity.Entity, child entity.Entity) error
	// Delete removes an entity and all its children from the hierarchy and returns the list the child descendants.
	Delete(child entity.Entity) []entity.Entity
	Parent(child entity.Entity) (entity.Entity, error)
	Children(parent entity.Entity) []entity.Entity
	Root() entity.Entity
}

// componentTypes is a list of all component types used in the ECS.
// This list is used to initialize the component stores in the ECS.
// It is also used to ensure that all component types are accounted for when managing entities.
func componentTypes() []component.Type {
	return []component.Type{
		position.Type,
		controllable.Type,
		shape.RectangleType,
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
	stores       *Stores
	hierarchy    Hierarchy
}

// New creates a new ECS system, initializing it with the provided component stores and event bus.
// This function ensures that the ECS is ready to manage entities and components from the start.
func New(eventBus *event.Bus, s *Stores, h Hierarchy) *ECS {
	return &ECS{
		lastEntityID: h.Root(),
		eventBus:     eventBus,
		stores:       s,
		hierarchy:    h,
	}
}

// CreateEntity generates a new unique entity by incrementing the lastEntityID.
// It also creates a position component for the entity and adds it to the ECS.
func (s *ECS) CreateEntity(parent entity.Entity, p point.Point) (entity.Entity, error) {
	e := s.CreateEmptyEntity()
	pos := position.New(parent, e, p)
	if _, err := s.AddPositionComponent(*pos); err != nil {
		return 0, err
	}
	return e, nil
}

// CreateEmptyEntity generates a new unique entity by incrementing the last entity ID.
func (s *ECS) CreateEmptyEntity() entity.Entity {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastEntityID++
	return s.lastEntityID
}

// DeleteEntity removes an entity and all its associated components from the ECS.
func (s *ECS) DeleteEntity(e entity.Entity) error {
	// Check for errors and return the first one that is not a "not found" error.
	for _, err := range s.deleteAllEntityComponents(e) {
		if err != nil && !errors.Is(err, ErrEntityNotFound) {
			return fmt.Errorf("failed to delete entity: %w", err)
		}
	}
	return nil
}

// deleteAllEntityComponents removes an entity and all its associated components from the ECS.
func (s *ECS) deleteAllEntityComponents(e entity.Entity) []error {
	var errs []error
	for _, t := range componentTypes() {
		switch t {
		case position.Type:
			if err := s.deletePositionByEntity(e); err != nil {
				errs = append(errs, err)
			}
		case controllable.Type:
			if err := s.deleteControllableByEntity(e); err != nil {
				errs = append(errs, err)
			}
		case shape.RectangleType:
			if err := s.deleteRectanglesByEntity(e); err != nil {
				errs = append(errs, err)
			}
		case sprite.Type:
			if err := s.deleteSpritesByEntity(e); err != nil {
				errs = append(errs, err)
			}
		case skeleton.Type:
			if err := s.deleteSkeletonByEntity(e); err != nil {
				errs = append(errs, err)
			}
		default:
			errs = append(errs, fmt.Errorf("unknown component type: %v", t))
		}
	}
	return errs
}

func (s *ECS) GetAbsolutePosition(e entity.Entity) (point.Point, error) {
	if e == s.hierarchy.Root() {
		return point.Point{}, nil
	}
	parent, err := s.hierarchy.Parent(e)
	if err != nil {
		return point.Point{}, err
	}
	pos, err := s.GetPosition(e)
	if err != nil {
		return point.Point{}, err
	}
	p, err := s.GetAbsolutePosition(parent)
	if err != nil {
		return point.Point{}, err
	}
	return p.Add(pos.Cords()), nil
}
