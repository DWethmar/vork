package gameplay

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
	"github.com/dwethmar/vork/point"
)

func addPlayer(parent entity.Entity, ecs *ecsys.ECS, p point.Point) error {
	e, err := ecs.CreateEntity(parent, p)
	if err != nil {
		return fmt.Errorf("could not create entity: %w", err)
	}
	if _, err = ecs.AddSkeletonComponent(*skeleton.New(e)); err != nil {
		return fmt.Errorf("could not add skeleton: %w", err)
	}
	if _, err = ecs.AddControllableComponent(*controllable.New(e)); err != nil {
		return fmt.Errorf("could not add controllable: %w", err)
	}
	return nil
}

func addEnemy(parent entity.Entity, ecs *ecsys.ECS, p point.Point) error {
	e, err := ecs.CreateEntity(parent, p)
	if err != nil {
		fmt.Printf("could not create entity: %v\n", err)
	}
	if _, err = ecs.AddSkeletonComponent(*skeleton.New(e)); err != nil {
		return fmt.Errorf("could not add skeleton: %w", err)
	}
	return nil
}
