package keyinput

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/dwethmar/vork/ecsys"
	"github.com/hajimehoshi/ebiten/v2"
)

// System is a controller system.
type System struct {
	logger              *slog.Logger
	ecs                 *ecsys.ECS
	velocityScaleFactor int
}

// Options is the options for the system.
type Options struct {
	Logger              *slog.Logger
	ECS                 *ecsys.ECS
	VelocityScaleFactor int
}

// New creates a new keyinput system. It moves all controllable entities in the direction of the direction keys.
func New(opts Options) *System {
	return &System{
		logger:              opts.Logger.With("system", "keyinput"),
		ecs:                 opts.ECS,
		velocityScaleFactor: opts.VelocityScaleFactor,
	}
}

// Init initializes the system.
func (s *System) Init() error {
	if s.ecs == nil {
		return errors.New("ecs is nil")
	}
	return nil
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
	for _, c := range s.ecs.AllControllables() {
		v, err := s.ecs.GetVelocity(c.Entity())
		if err != nil {
			return fmt.Errorf("failed to get velocity component: %w", err)
		}

		v.X = x * s.velocityScaleFactor
		v.Y = y * s.velocityScaleFactor

		if err = s.ecs.UpdateVelocityComponent(v); err != nil {
			return fmt.Errorf("failed to update velocity component: %w", err)
		}
	}
	return nil
}

func (s *System) Draw(_ *ebiten.Image) error {
	return nil
}
