package collision

import (
	"errors"
	"log/slog"
	"math"
	"sync"

	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
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
			continue
		}

		// Convert to float for normalization
		x := float64(vel.X)
		y := float64(vel.Y)

		// Calculate the magnitude of the velocity vector
		magnitude := math.Sqrt(x*x + y*y)

		// If the magnitude is greater than zero, normalize the velocity
		if magnitude > 0 {
			x /= magnitude
			y /= magnitude
		}

		// Update position based on the normalized velocity (using rounding)
		if err = s.collide(pos, int(math.Round(x)), int(math.Round(y))); err != nil {
			return err
		}
	}
	return nil
}

func (s *System) collide(pos position.Position, velX, velY int) error {
	// Get the hitbox of the moving entity
	hbList := s.ecs.ListHitboxesByEntity(pos.Entity())
	if len(hbList) == 0 {
		return errors.New("no hitbox found for entity")
	}
	hb := &hbList[0]

	// Get all hitboxes
	hbs := s.ecs.ListHitboxes()

	// Store original position
	origPos := pos

	// Initialize collision flags
	var collisionX, collisionY bool

	// Check collision along X-axis
	collisionX, err := s.checkCollision(pos, hb, hbs, velX, 0)
	if err != nil {
		return err
	}
	if collisionX {
		pos.X = origPos.X // Rollback X movement
	} else {
		pos.X += velX // Apply X movement
	}

	// Check collision along Y-axis
	collisionY, err = s.checkCollision(pos, hb, hbs, 0, velY)
	if err != nil {
		return err
	}
	if collisionY {
		pos.Y = origPos.Y // Rollback Y movement
	} else {
		pos.Y += velY // Apply Y movement
	}

	// Update velocity after collision
	if collisionX || collisionY {
		if err = s.updateVelocityAfterCollision(pos.Entity(), collisionX, collisionY); err != nil {
			return err
		}
	}

	// Update position component
	return s.ecs.UpdatePositionComponent(pos)
}

func (s *System) checkCollision(
	pos position.Position,
	hb *hitbox.Hitbox,
	hbs []hitbox.Hitbox,
	deltaX, deltaY int,
) (bool, error) {
	// Move position by delta values
	pos.X += deltaX
	pos.Y += deltaY

	// Get moving entity's bounding box
	movingBox := getBoundingBox(pos, hb)

	// Check for collisions
	for _, otherHb := range hbs {
		if otherHb.Entity() == pos.Entity() {
			continue
		}

		otherPos, err := s.ecs.GetPosition(otherHb.Entity())
		if err != nil {
			if errors.Is(err, ecsys.ErrComponentNotFound) {
				continue
			}
			return false, err
		}

		otherBox := getBoundingBox(otherPos, &otherHb)

		if boxesOverlap(movingBox, otherBox) {
			return true, nil
		}
	}

	return false, nil
}

func (s *System) updateVelocityAfterCollision(
	entityID entity.Entity,
	collisionX, collisionY bool,
) error {
	vel, err := s.ecs.GetVelocity(entityID)
	if err != nil {
		return err
	}

	if collisionX {
		vel.X = 0
	}
	if collisionY {
		vel.Y = 0
	}

	// Update the velocity component
	return s.ecs.UpdateVelocityComponent(vel)
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
