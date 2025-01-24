package ecsys

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
)

// AllPositions returns all positions.
func (s *ECS) AllPositions() []position.Position {
	return derefSlice(s.stores.Position.List())
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) AllControllables() []controllable.Controllable {
	return derefSlice(s.stores.Controllable.List())
}

// ListPositions returns all positions.
func (s *ECS) AllHitboxes() []hitbox.Hitbox {
	return derefSlice(s.stores.Hitbox.List())
}

// ListPositionsByEntity returns all positions for a given entity.
func (s *ECS) AllRectangles() []shape.Rectangle {
	return derefSlice(s.stores.Rectangle.List())
}

// AllSkeletons returns all skeletons.
func (s *ECS) AllSkeletons() []skeleton.Skeleton {
	return derefSlice(s.stores.Skeleton.List())
}

// AllSprites returns all sprites.
func (s *ECS) AllSprites() []sprite.Sprite {
	return derefSlice(s.stores.Sprite.List())
}
