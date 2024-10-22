package gameplay

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/point"
)

func addPlayer(parent entity.Entity, ecs *ecsys.ECS, p point.Point) (entity.Entity, error) {
	e, err := ecs.CreateEntity(parent, p)
	if err != nil {
		return e, fmt.Errorf("could not create entity: %w", err)
	}
	if _, err = ecs.AddSkeleton(*skeleton.New(e)); err != nil {
		return e, fmt.Errorf("could not add skeleton: %w", err)
	}
	if _, err = ecs.AddControllable(*controllable.New(e)); err != nil {
		return e, fmt.Errorf("could not add controllable: %w", err)
	}
	return e, nil
}

func addEnemy(parent entity.Entity, ecs *ecsys.ECS, p point.Point) (entity.Entity, error) {
	e, err := ecs.CreateEntity(parent, p)
	if err != nil {
		return e, fmt.Errorf("could not create entity: %w", err)
	}
	if _, err = ecs.AddSkeleton(*skeleton.New(e)); err != nil {
		return e, fmt.Errorf("could not add skeleton: %w", err)
	}
	return e, nil
}
