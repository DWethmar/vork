package gameplay

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/event"
)

type Stores struct {
	EntityStore  *entity.Store
	Controllable *controllable.Store
	Skeletons    *skeleton.Store
	Positions    *position.Store
	Shapes       *shape.RectangleStore
	Sprites      *sprite.Store
	Hitboxes     *hitbox.Store
	Velocities   *velocity.Store
}

func NewStores(eb *event.Bus) *Stores {
	return &Stores{
		EntityStore:  entity.NewStore(),
		Controllable: controllable.NewStore(eb),
		Skeletons:    skeleton.NewStore(eb),
		Positions:    position.NewStore(eb),
		Shapes:       shape.NewRectangleStore(),
		Sprites:      sprite.NewStore(),
		Hitboxes:     hitbox.NewStore(eb),
		Velocities:   velocity.NewStore(eb),
	}
}
