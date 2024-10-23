package collision

import (
	"errors"
	"log/slog"
	"sync"

	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/hajimehoshi/ebiten/v2"
)

// System is a collision system.
type System struct {
	logger              *slog.Logger
	ecs                 *ecsys.ECS
	eventBus            *event.Bus
	velocityScaleFactor int // Scale factor for the velocity
	friction            int // Friction to apply to the velocity
	velocityThreshold   int // Threshold for velocity to stop movement
	subscriptions       []int
	mux                 sync.RWMutex
	moving              map[uint]*velocity.Velocity
}

// Options for the collision system.
type Options struct {
	Logger              *slog.Logger
	ECS                 *ecsys.ECS
	EventBus            *event.Bus
	VelocityScaleFactor int // Scale factor for the velocity
	Friction            int // Friction to apply to the velocity
	VelocityThreshold   int // Threshold for velocity to stop movement
}

// New creates a new collision system.
func New(opts Options) *System {
	return &System{
		logger:              opts.Logger.With("system", "collision"),
		ecs:                 opts.ECS,
		eventBus:            opts.EventBus,
		moving:              make(map[uint]*velocity.Velocity),
		velocityScaleFactor: opts.VelocityScaleFactor,
		friction:            opts.Friction,
		velocityThreshold:   opts.VelocityThreshold,
	}
}

// Init initializes the system.
func (s *System) Init() error {
	if s.logger == nil {
		return errors.New("logger is nil")
	}
	if s.ecs == nil {
		return errors.New("ecs is nil")
	}
	if s.eventBus == nil {
		return errors.New("event bus is nil")
	}

	posEventsMatcher := event.MatchAny(velocity.CreatedEventType, velocity.UpdatedEventType)
	s.subscriptions = []int{
		s.eventBus.Subscribe(posEventsMatcher, s.onVelocityEvent),
	}
	return nil
}

func (s *System) onVelocityEvent(event event.Event) error {
	pe, ok := event.(velocity.Event)
	if !ok {
		return errors.New("event is not a position event")
	}

	s.logger.Debug("Position event received", slog.Any("entityID", pe.Velocity().ID()))

	if ve := pe.Velocity(); ve.Zero() || pe.Deleted() {
		delete(s.moving, pe.Velocity().ID())
	} else {
		s.moving[pe.Velocity().ID()] = ve
	}
	return nil
}

// Close closes the system.
func (s *System) Close() error {
	for _, id := range s.subscriptions {
		s.eventBus.Unsubscribe(id)
	}
	return nil
}

// Update checks for collisions and updates the positions based on velocity.
func (s *System) Update() error {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, vel := range s.moving {
		// Get position of the entity associated with this velocity
		pos, err := s.ecs.GetPosition(vel.Entity())
		if err != nil {
			return err
		}

		// Apply friction to velocity and scale
		vel.X = (vel.X * s.friction) / s.velocityScaleFactor
		vel.Y = (vel.Y * s.friction) / s.velocityScaleFactor

		// If the velocity is too small, stop the movement
		if abs(vel.X) < s.velocityThreshold && abs(vel.Y) < s.velocityThreshold {
			vel.X = 0
			vel.Y = 0

			// Update the velocity component if it's effectively zero
			if err = s.ecs.UpdateVelocityComponent(*vel); err != nil {
				return err
			}
		}

		s.logger.Info("Updating position", slog.Any("entityID", pos.ID()), slog.Any("velocity", vel.Point))

		// Update position based on the velocity
		pos.X += vel.X
		pos.Y += vel.Y

		// Update the position component in ECS
		if err = s.ecs.UpdatePositionComponent(pos); err != nil {
			return err
		}
	}
	return nil
}

// Utility function for absolute value.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Draw draws the system.
func (s *System) Draw(_ *ebiten.Image) error {
	return nil
}
