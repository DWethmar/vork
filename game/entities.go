package game

import (
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/scene"
)

func addPlayer(s scene.Scene, x, y int64) {
	e := s.CreateEntity()
	s.AddComponent(shape.NewRectangle(e, 10, 10))
	s.AddComponent(position.New(e, x, y))
	s.AddComponent(controllable.New(e))
	s.AddComponent(sprite.New(e, "player"))
}

func addEnemy(s scene.Scene, x, y int64) {
	e := s.CreateEntity()
	s.AddComponent(shape.NewRectangle(e, 10, 10))
	s.AddComponent(position.New(e, x, y))
}
