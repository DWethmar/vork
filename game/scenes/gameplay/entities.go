package gameplay

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
)

func addPlayer(ecs *ecsys.ECS, x, y int64) error {
	e, err := ecs.CreateEntity(x, y)
	if err != nil {
		return fmt.Errorf("could not create entity: %v", err)
	}
	if _, err = ecs.AddSkeletonComponent(*skeleton.New(e)); err != nil {
		return fmt.Errorf("could not add skeleton: %v", err)
	}
	if _, err = ecs.AddControllableComponent(*controllable.New(e)); err != nil {
		return fmt.Errorf("could not add controllable: %v", err)
	}
	return nil
}

func addEnemy(ecs *ecsys.ECS, x, y int64) error {
	e, err := ecs.CreateEntity(x, y)
	if err != nil {
		fmt.Printf("could not create entity: %v\n", err)
	}
	if _, err = ecs.AddSkeletonComponent(*skeleton.New(e)); err != nil {
		return fmt.Errorf("could not add skeleton: %v", err)
	}
	return nil
}
