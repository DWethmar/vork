package chunk

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/grid"
	"github.com/dwethmar/vork/point"
)

const (
	Type        = component.Type("chunk")
	ChunkWidth  = 16
	ChunkHeight = 16
)

// Chunk is a component thats part of the terrain.
type Chunk struct {
	component.BaseComponent
	Index point.Point
	Tiles grid.Grid
}

// New creates a new chunk.
func New(e entity.Entity, index point.Point) *Chunk {
	return &Chunk{
		BaseComponent: component.BaseComponent{
			ID:     0,
			Entity: e,
		},
		Index: index,
		Tiles: grid.New(ChunkWidth, ChunkHeight, 0),
	}
}

func (c *Chunk) ID() uint              { return c.BaseComponent.ID }
func (c *Chunk) SetID(i uint)          { c.BaseComponent.ID = i }
func (c *Chunk) Type() component.Type  { return Type }
func (c *Chunk) Entity() entity.Entity { return c.BaseComponent.Entity }
