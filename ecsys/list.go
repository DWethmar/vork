package ecsys

import (
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/entity"
)

// derefSlice dereferences a slice of pointers to a slice of values.
func derefSlice[T any](s []*T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = *v
	}
	return r
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) ListHitboxes(e entity.Entity) []hitbox.Hitbox {
	return derefSlice(s.stores.Hitbox.ListByEntity(e))
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) ListRectangles(e entity.Entity) []shape.Rectangle {
	return derefSlice(s.stores.Rectangle.ListByEntity(e))
}

// SpritesByEntity returns all sprites for a given entity.
func (s *ECS) ListSprites(e entity.Entity) []sprite.Sprite {
	return derefSlice(s.stores.Sprite.ListByEntity(e))
}
