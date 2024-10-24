package chunk

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/entity"
)

// Chunk is a component thats part of the terrain.
type Chunk struct {
	component.BaseComponent
}

func New(e entity.Entity) *Chunk {
	return &Chunk{
		BaseComponent: component.BaseComponent{
			ID:     0,
			Entity: e,
		},
	}
}
