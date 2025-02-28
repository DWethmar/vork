package gameplay

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/point"
)

func addPlayer(parent entity.Entity, stores *Stores, p point.Point) (entity.Entity, error) {
	e := stores.EntityStore.CreateEntity()
	if _, err := stores.Positions.Add(*position.New(0, e, p)); err != nil {
		return e, fmt.Errorf("could not add position component to entity %v: %w", e, err)
	}
	if _, err := stores.Skeletons.Add(skeleton.New(e)); err != nil {
		return e, fmt.Errorf("could not add skeleton: %w", err)
	}
	if _, err := stores.Controllable.Add(*controllable.New(e)); err != nil {
		return e, fmt.Errorf("could not add controllable: %w", err)
	}
	if _, err := stores.Velocities.Add(*velocity.New(e, point.Zero())); err != nil {
		return e, fmt.Errorf("could not add velocity component to entity %v: %w", e, err)
	}
	return e, nil
}

func addEnemy(parent entity.Entity, stores *Stores, p point.Point) (entity.Entity, error) {
	e := stores.EntityStore.CreateEntity()
	if _, err := stores.Positions.Add(*position.New(0, e, p)); err != nil {
		return e, fmt.Errorf("could not add position component to entity %v: %w", e, err)
	}
	if _, err := stores.Skeletons.Add(skeleton.New(e)); err != nil {
		return e, fmt.Errorf("could not add skeleton: %w", err)
	}
	if _, err := stores.Velocities.Add(*velocity.New(e, point.Zero())); err != nil {
		return e, fmt.Errorf("could not add velocity component to entity %v: %w", e, err)
	}
	return e, nil
}
