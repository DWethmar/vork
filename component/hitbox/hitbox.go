package hitbox

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/point"
)

const (
	Type = component.Type("hitbox")
)

// Hitbox is a component that holds the hitbox of an entity.
type Hitbox struct {
	I      uint          // ID
	E      entity.Entity // Entity
	Tag    string        // Tag is a string that describes the hitbox.
	Width  int
	Height int
	Offset point.Point
}

// New creates a new hitbox component.
func New(e entity.Entity, tag string, width, height int, offset point.Point) *Hitbox {
	return &Hitbox{
		I:      0,
		E:      e,
		Tag:    tag,
		Width:  width,
		Height: height,
		Offset: offset,
	}
}

func (h *Hitbox) SetID(i uint)          { h.I = i }
func (h *Hitbox) ID() uint              { return h.I }
func (h *Hitbox) Type() component.Type  { return Type }
func (h *Hitbox) Entity() entity.Entity { return h.E }
