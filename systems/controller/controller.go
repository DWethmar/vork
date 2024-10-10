package controller

import (
	"log/slog"

	"github.com/dwethmar/vork/ecsys"
	"github.com/hajimehoshi/ebiten/v2"
)

// System is a controller system.
type System struct {
	logger *slog.Logger
	ecs    *ecsys.ECS
}

// New creates a new controller system. It moves all controllable entities in the direction of the direction keys.
func New(logger *slog.Logger, ecs *ecsys.ECS) *System {
	return &System{
		logger: logger,
		ecs:    ecs,
	}
}

// Close closes the system.
func (s *System) Close() error {
	return nil
}

func (s *System) Update() error {
	x, y := direction()
	if x == 0 && y == 0 {
		return nil
	}
	for _, c := range s.ecs.ListControllables() {
		p, err := s.ecs.GetPosition(c.Entity())
		if err != nil {
			return err
		}
		p.X += x
		p.Y += y
		if err = s.ecs.UpdatePositionComponent(p); err != nil {
			return err
		}
	}
	return nil
}

func (s *System) Draw(_ *ebiten.Image) error {
	return nil
}
