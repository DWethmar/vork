package game

import (
	"image/color"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/ecsys"
)

func addPlayer(ecs *ecsys.ECS, x, y int64) {
	e := ecs.CreateEntity()
	ecs.AddRectangle(*shape.NewRectangle(e, 10, 10, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}))
	ecs.AddPosition(*position.New(e, x, y))
	ecs.AddControllable(*controllable.New(e))
	ecs.AddSprite(*sprite.New(e, sprite.SkeletonMoveDown1))
}

func addEnemy(ecs *ecsys.ECS, x, y int64) {
	e := ecs.CreateEntity()
	ecs.AddRectangle(*shape.NewRectangle(e, 10, 10, color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}))
	ecs.AddPosition(*position.New(e, x, y))
	ecs.AddSprite(*sprite.New(e, sprite.SkeletonMoveUp1))
}
