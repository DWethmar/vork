package keyinput

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/hajimehoshi/ebiten/v2"
)

// System is a controller system.
type System struct {
	logger              *slog.Logger
	velocityStore       *velocity.Store
	controllableStore   *controllable.Store
	velocityScaleFactor int
}

// Options is the options for the system.
type Options struct {
	Logger              *slog.Logger
	VelocityStore       *velocity.Store
	ControllableStore   *controllable.Store
	VelocityScaleFactor int
}

// New creates a new keyinput system. It moves all controllable entities in the direction of the direction keys.
func New(opts Options) *System {
	return &System{
		logger:              opts.Logger.With("system", "keyinput"),
		velocityStore:       opts.VelocityStore,
		controllableStore:   opts.ControllableStore,
		velocityScaleFactor: opts.VelocityScaleFactor,
	}
}

// Init initializes the system.
func (s *System) Init() error {
	if s.logger == nil {
		return errors.New("logger is nil")
	}
	if s.velocityStore == nil {
		return errors.New("velocity store is nil")
	}
	if s.controllableStore == nil {
		return errors.New("controllable store is nil")
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
	for _, c := range s.controllableStore.All() {
		v, err := s.velocityStore.Get(c.Entity())
		if err != nil {
			return fmt.Errorf("failed to get velocity component: %w", err)
		}

		v.X = x * s.velocityScaleFactor
		v.Y = y * s.velocityScaleFactor

		if err = s.velocityStore.Update(*v); err != nil {
			return fmt.Errorf("failed to update velocity component: %w", err)
		}
	}
	return nil
}

func (s *System) Draw(_ *ebiten.Image) error {
	return nil
}
