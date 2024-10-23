package skeletons

import (
	"errors"
	"fmt"
	"image/color"
	"log/slog"

	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/direction"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/event/mouse"
	"github.com/dwethmar/vork/point"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	walkAnimationPerFrames = 4
	walkAnimationSteps     = 8 * walkAnimationPerFrames // frames every 8 steps (3 frames per step)
)

// System is a system that manages skeletons in the game.
type System struct {
	logger        *slog.Logger
	ecs           *ecsys.ECS
	eventBus      *event.Bus
	subscriptions []int
}

// New creates a new skeleton system. It listens to skeleton events and adds the necessary components to the entity to make it a skeleton.
func New(logger *slog.Logger, ecs *ecsys.ECS, eventBus *event.Bus) *System {
	s := &System{
		logger:        logger.With("system", "skeletons"),
		ecs:           ecs,
		eventBus:      eventBus,
		subscriptions: []int{},
	}

	// Subscribe to the skeleton events
	s.subscriptions = append(s.subscriptions, s.eventBus.Subscribe(
		event.MatchAny(skeleton.UpdatedEventType, skeleton.CreatedEventType, skeleton.DeletedEventType),
		s.skeletonCreatedHandler,
	))

	s.subscriptions = append(s.subscriptions, s.eventBus.Subscribe(
		event.MatchAny(mouse.LeftMouseClickedEventType),
		s.skeletonCreatedHandler,
	))

	return s
}

// Init initializes the system.
func (s *System) Init() error {
	if s.ecs == nil {
		return errors.New("ecs is nil")
	}
	if s.eventBus == nil {
		return errors.New("eventBus is nil")
	}
	// Setup existing skeletons
	for _, sk := range s.ecs.ListSkeletons() {
		if err := s.setupSkeleton(sk); err != nil {
			return fmt.Errorf("could not setup skeleton (%v): %w", sk.Entity(), err)
		}
	}
	return nil
}

// Close closes the system.
func (s *System) Close() error {
	for _, sub := range s.subscriptions {
		s.eventBus.Unsubscribe(sub)
	}
	return nil
}

func (s *System) skeletonCreatedHandler(e event.Event) error {
	switch e := e.(type) {
	case *skeleton.CreatedEvent:
		s.logger.Debug("skeleton created", "skeleton", e.Skeleton)
		if err := s.setupSkeleton(*e.Skeleton()); err != nil {
			return err
		}
	case *skeleton.UpdatedEvent:
		s.logger.Debug("skeleton updated", "skeleton", e.Skeleton)
	case *skeleton.DeletedEvent:
		s.logger.Debug("skeleton deleted", "skeleton", e.Skeleton)
	case *mouse.LeftClickedEvent:
		s.logger.Info("clicked", "x", e.X, "y", e.Y)
	default:
		return fmt.Errorf("unhandled event type %T", e)
	}
	return nil
}

// setupSkeleton adds the necessary components to the entity to make it a skeleton.
func (s *System) setupSkeleton(sk skeleton.Skeleton) error {
	e := sk.Entity()
	rect := shape.NewRectangle(e, 10, 10, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff})
	if _, err := s.ecs.AddRectangle(*rect); err != nil {
		return fmt.Errorf("could not add rectangle component to entity %v: %w", e, err)
	}
	if _, err := s.ecs.AddSprite(*sprite.New(e, "skeleton", sprite.SkeletonMoveDown1)); err != nil {
		return fmt.Errorf("could not add sprite component to entity %v: %w", e, err)
	}

	// ensure velocity component is present
	if _, err := s.ecs.GetVelocity(e); err != nil {
		if errors.Is(err, ecsys.ErrEntityNotFound) || errors.Is(err, ecsys.ErrComponentNotFound) {
			vel := velocity.New(e, point.Zero())
			if _, err = s.ecs.AddVelocity(*vel); err != nil {
				return fmt.Errorf("could not add velocity component to entity %v: %w", e, err)
			}
		} else {
			return fmt.Errorf("could not get velocity component for entity %v: %w", e, err)
		}
	}

	return nil
}

func (s *System) Draw(_ *ebiten.Image) error {
	return nil
}

// Update updates the skeletons in the ECS.
func (s *System) Update() error {
	skeletons := s.ecs.ListSkeletons()
	for i := range skeletons {
		e := &skeletons[i]
		if err := s.updateSkeleton(e); err != nil {
			return err
		}

		// Update the skeleton component in the ECS
		if err := s.ecs.UpdateSkeletonComponent(*e); err != nil {
			return fmt.Errorf("could not update skeleton: %w", err)
		}

		// Move the sprite updating code to a separate function
		if err := s.updateSprite(e); err != nil {
			return err
		}
	}

	return nil
}

// updateSkeleton applies skeleton behavior to the entity.
func (s *System) updateSkeleton(e *skeleton.Skeleton) error {
	pos, err := s.ecs.GetPosition(e.Entity())
	if err != nil {
		return fmt.Errorf("could not get position: %w", err)
	}

	isMoving := false
	x, y := pos.Cords()
	if e.PrefX != x || e.PrefY != y {
		isMoving = true
		// Calculate facing direction before updating e.PrefX and e.PrefY
		facing := direction.Get(e.PrefX, e.PrefY, x, y)
		if e.Facing != facing {
			e.AnimationStep = 0
		}
		e.Facing = facing
	}

	if isMoving {
		e.State = skeleton.Moving
	} else {
		e.State = skeleton.Idle
	}

	// Update e.PrefX and e.PrefY after calculating facing
	e.PrefX, e.PrefY = x, y
	if e.State == skeleton.Moving {
		e.AnimationStep++
		if e.AnimationStep >= walkAnimationSteps {
			e.AnimationStep = 0
		}
	}

	return nil
}

// updateSprite updates the sprite associated with the skeleton based on its state and facing direction.
func (s *System) updateSprite(e *skeleton.Skeleton) error {
	// Retrieve the sprite component associated with the skeleton
	var spr *sprite.Sprite
	sprites := s.ecs.ListSpritesByEntity(e.Entity())
	for i := range sprites {
		sp := &sprites[i]
		if sp.Tag == "skeleton" {
			spr = sp
			break // Break early once found
		}
	}
	if spr == nil {
		s.logger.Error("could not find sprite", "entity", e.Entity())
		return fmt.Errorf("could not find sprite for entity %v", e.Entity())
	}

	// Determine the appropriate graphic based on state and facing direction
	step := int(e.AnimationStep / walkAnimationPerFrames) // 3 frames per step
	var graphic sprite.Graphic
	if e.State == skeleton.Moving {
		switch e.Facing {
		case direction.None:
			graphic = sprite.SkeletonMoveDownFrames()[step]
		case direction.North, direction.NorthEast, direction.NorthWest:
			graphic = sprite.SkeletonMoveUpFrames()[step]
		case direction.South, direction.SouthEast, direction.SouthWest:
			graphic = sprite.SkeletonMoveDownFrames()[step]
		case direction.East:
			graphic = sprite.SkeletonMoveRightFrames()[step]
		case direction.West:
			graphic = sprite.SkeletonMoveLeftFrames()[step]
		default:
			s.logger.Warn("Unhandled facing direction", "direction", e.Facing)
			graphic = spr.Graphic // Keep current graphic if direction is unhandled
		}
	} else { // Idle state
		switch e.Facing {
		case direction.None:
			graphic = sprite.SkeletonMoveDown1
		case direction.North, direction.NorthEast, direction.NorthWest:
			graphic = sprite.SkeletonMoveUp1
		case direction.South, direction.SouthEast, direction.SouthWest:
			graphic = sprite.SkeletonMoveDown1
		case direction.East:
			graphic = sprite.SkeletonMoveRight1
		case direction.West:
			graphic = sprite.SkeletonMoveLeft1
		default:
			s.logger.Warn("Unhandled facing direction", "direction", e.Facing)
			graphic = spr.Graphic // Keep current graphic if direction is unhandled
		}
	}

	// Update the sprite's graphic if it has changed
	if spr.Graphic != graphic {
		spr.Graphic = graphic
		if err := s.ecs.UpdateSpriteComponent(*spr); err != nil {
			return fmt.Errorf("could not update skeleton sprite: %w", err)
		}
	}

	return nil
}
