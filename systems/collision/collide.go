package collision

import (
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
)

// BoundingBox is a rectangle.
type BoundingBox struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

// getBoundingBox returns the bounding box of a hitbox.
func getBoundingBox(pos position.Position, hb *hitbox.Hitbox) BoundingBox {
	left := pos.X + hb.Offset.X
	top := pos.Y + hb.Offset.Y
	right := left + hb.Width
	bottom := top + hb.Height

	return BoundingBox{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
	}
}

// boxesOverlap returns true if two bounding boxes overlap.
func boxesOverlap(a, b BoundingBox) bool {
	// If one rectangle is on the left side of the other
	if a.Right <= b.Left || b.Right <= a.Left {
		return false
	}
	// If one rectangle is above the other
	if a.Bottom <= b.Top || b.Bottom <= a.Top {
		return false
	}
	return true
}
