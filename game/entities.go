package game

import (
	"fmt"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/ecsys"
)

func addPlayer(ecs *ecsys.ECS, x, y int64) {
	e := ecs.CreateEntity()
	fmt.Printf("player entity: %v\n", e)
	ecs.AddPosition(*position.New(e, x, y))
	ecs.AddSkeleton(*skeleton.New(e))
	ecs.AddControllable(*controllable.New(e))
}

func addEnemy(ecs *ecsys.ECS, x, y int64) {
	e := ecs.CreateEntity()
	fmt.Printf("enemy entity: %v\n", e)
	ecs.AddPosition(*position.New(e, x, y))
	ecs.AddSkeleton(*skeleton.New(e))
}
