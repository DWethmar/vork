package controller

import (
	"log/slog"

	"github.com/dwethmar/vork/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type System struct {
	logger *slog.Logger
	ecs    *systems.ECS
}

func New(logger *slog.Logger, ecs *systems.ECS) *System {
	return &System{
		logger: logger,
		ecs:    ecs,
	}
}

func (s *System) Update() error {
	x, y := direction()
	for _, c := range s.ecs.Controllables() {
		p, err := s.ecs.Position(c.Entity())
		if err != nil {
			return err
		}
		p.X += int64(x)
		p.Y += int64(y)
		if err := s.ecs.UpdatePosition(p); err != nil {
			return err
		}
	}
	return nil
}

func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}
