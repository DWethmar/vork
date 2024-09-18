package game

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/systems"
)

func addPlayer(ecs *systems.ECS, x, y int64) {
	e := ecs.CreateEntity()
	ecs.AddRectangle(*shape.NewRectangle(e, 10, 10))
	ecs.AddPosition(*position.New(e, x, y))
	ecs.AddControllable(*controllable.New(e))
	ecs.AddSprite(*sprite.New(e, "player"))
}

func addEnemy(ecs *systems.ECS, x, y int64) {
	e := ecs.CreateEntity()
	ecs.AddRectangle(*shape.NewRectangle(e, 10, 10))
	ecs.AddPosition(*position.New(e, x, y))
	ecs.AddSprite(*sprite.New(e, "enemy"))
}
