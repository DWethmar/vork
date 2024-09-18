package systems

import (
	"sync"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
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

// ECS is the Entity-Component-System architecture.
type ECS struct {
	mu           sync.RWMutex
	lastEntityID entity.Entity
	// component stores
	positions     PositionStore
	controllables ControllableStore
	rectangles    RectanglesStore
	sprites       SpriteStore
}

func (s *ECS) CreateEntity() entity.Entity {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastEntityID++
	return s.lastEntityID
}

func NewECS(
	positions PositionStore,
	controllables ControllableStore,
	rectangles RectanglesStore,
	sprites SpriteStore,
) *ECS {
	return &ECS{
		lastEntityID:  0,
		positions:     positions,
		controllables: controllables,
		rectangles:    rectangles,
		sprites:       sprites,
	}
}

func (s *ECS) Position(e entity.Entity) (position.Position, error) {
	return s.positions.FirstByEntity(e)
}
func (s *ECS) UpdatePosition(p position.Position) error { return s.positions.Update(p) }
func (s *ECS) AddPosition(p position.Position) error    { return s.positions.Add(p) }
func (s *ECS) DeletePosition(id uint32) error           { return s.positions.Delete(id) }

func (s *ECS) Controllable(e entity.Entity) (controllable.Controllable, error) {
	return s.controllables.FirstByEntity(e)
}
func (s *ECS) Controllables() []controllable.Controllable           { return s.controllables.List() }
func (s *ECS) UpdateControllable(c controllable.Controllable) error { return s.controllables.Update(c) }
func (s *ECS) AddControllable(c controllable.Controllable) error    { return s.controllables.Add(c) }
func (s *ECS) DeleteControllable(id uint32) error                   { return s.controllables.Delete(id) }

func (s *ECS) Rectangle(e entity.Entity) (shape.Rectangle, error) {
	return s.rectangles.FirstByEntity(e)
}
func (s *ECS) Rectangles() []shape.Rectangle           { return s.rectangles.List() }
func (s *ECS) UpdateRectangle(r shape.Rectangle) error { return s.rectangles.Update(r) }
func (s *ECS) AddRectangle(r shape.Rectangle) error    { return s.rectangles.Add(r) }
func (s *ECS) DeleteRectangle(id uint32) error         { return s.rectangles.Delete(id) }

func (s *ECS) Sprites(e entity.Entity) ([]sprite.Sprite, error) {
	return s.sprites.ListByEntity(e)
}
func (s *ECS) UpdateSprite(sp sprite.Sprite) error { return s.sprites.Update(sp) }
func (s *ECS) AddSprite(sp sprite.Sprite) error    { return s.sprites.Add(sp) }
func (s *ECS) DeleteSprite(id uint32) error        { return s.sprites.Delete(id) }
