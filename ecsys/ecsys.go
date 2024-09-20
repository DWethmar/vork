// package ecsys contains the Entity-Component-System architecture.
package ecsys

import (
	"sync"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

type ComponentStore interface {
	Delete(uint32) error
	DeleteByEntity(entity.Entity) error
}

type ControllableStore interface {
	ComponentStore
	Add(controllable.Controllable) error
	Get(uint32) (controllable.Controllable, error)
	FirstByEntity(entity.Entity) (controllable.Controllable, error)
	Update(controllable.Controllable) error
	List() []controllable.Controllable
}

type PositionStore interface {
	ComponentStore
	Add(position.Position) error
	Get(uint32) (position.Position, error)
	FirstByEntity(entity.Entity) (position.Position, error)
	Update(position.Position) error
	List() []position.Position
}

type RectanglesStore interface {
	ComponentStore
	Add(shape.Rectangle) error
	Get(uint32) (shape.Rectangle, error)
	FirstByEntity(entity.Entity) (shape.Rectangle, error)
	Update(shape.Rectangle) error
	List() []shape.Rectangle
}

type SpriteStore interface {
	ComponentStore
	Add(sprite.Sprite) error
	Get(uint32) (sprite.Sprite, error)
	ListByEntity(entity.Entity) ([]sprite.Sprite, error)
	Update(sprite.Sprite) error
	List() []sprite.Sprite
}

type SkeletonStore interface {
	ComponentStore
	Add(skeleton.Skeleton) error
	Get(uint32) (skeleton.Skeleton, error)
	FirstByEntity(entity.Entity) (skeleton.Skeleton, error)
	Update(skeleton.Skeleton) error
	List() []skeleton.Skeleton
}

// ECS is the Entity-Component-System architecture.
type ECS struct {
	mu           sync.RWMutex
	lastEntityID entity.Entity
	eventBus     *event.Bus
	// component stores
	pos     PositionStore
	contr   ControllableStore
	rect    RectanglesStore
	sprites SpriteStore
	sklt    SkeletonStore
}

func (s *ECS) CreateEntity() entity.Entity {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastEntityID++
	return s.lastEntityID
}

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

func (s *ECS) Position(e entity.Entity) (position.Position, error) { return s.pos.FirstByEntity(e) }
func (s *ECS) UpdatePosition(p position.Position) error            { return s.pos.Update(p) }
func (s *ECS) AddPosition(p position.Position) error               { return s.pos.Add(p) }
func (s *ECS) DeletePosition(id uint32) error                      { return s.pos.Delete(id) }

func (s *ECS) Controllable(e entity.Entity) (controllable.Controllable, error) {
	return s.contr.FirstByEntity(e)
}
func (s *ECS) Controllables() []controllable.Controllable           { return s.contr.List() }
func (s *ECS) UpdateControllable(c controllable.Controllable) error { return s.contr.Update(c) }
func (s *ECS) AddControllable(c controllable.Controllable) error    { return s.contr.Add(c) }
func (s *ECS) DeleteControllable(id uint32) error                   { return s.contr.Delete(id) }

func (s *ECS) Rectangle(e entity.Entity) (shape.Rectangle, error) { return s.rect.FirstByEntity(e) }
func (s *ECS) Rectangles() []shape.Rectangle                      { return s.rect.List() }
func (s *ECS) UpdateRectangle(r shape.Rectangle) error            { return s.rect.Update(r) }
func (s *ECS) AddRectangle(r shape.Rectangle) error               { return s.rect.Add(r) }
func (s *ECS) DeleteRectangle(id uint32) error                    { return s.rect.Delete(id) }

func (s *ECS) Sprites() []sprite.Sprite            { return s.sprites.List() }
func (s *ECS) UpdateSprite(sp sprite.Sprite) error { return s.sprites.Update(sp) }
func (s *ECS) AddSprite(sp sprite.Sprite) error    { return s.sprites.Add(sp) }
func (s *ECS) DeleteSprite(id uint32) error        { return s.sprites.Delete(id) }

func (s *ECS) Skeleton(e entity.Entity) (skeleton.Skeleton, error) {
	return s.sklt.FirstByEntity(e)
}

func (s *ECS) Skeletons() []skeleton.Skeleton {
	return s.sklt.List()
}

func (s *ECS) UpdateSkeleton(sk skeleton.Skeleton) error {
	if err := s.sklt.Update(sk); err != nil {
		return err
	}
	s.eventBus.Publish(skeleton.NewUpdatedEvent(sk))
	return nil
}

func (s *ECS) AddSkeleton(sk skeleton.Skeleton) error {
	if err := s.sklt.Add(sk); err != nil {
		return err
	}
	s.eventBus.Publish(skeleton.NewCreatedEvent(sk))
	return nil
}

func (s *ECS) DeleteSkeleton(sk skeleton.Skeleton) error {
	if err := s.sklt.Delete(sk.ID()); err != nil {
		return err
	}
	s.eventBus.Publish(skeleton.NewDeletedEvent(sk))
	return nil
}
