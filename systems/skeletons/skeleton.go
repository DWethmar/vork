package skeletons

import (
	"errors"
	"fmt"
	"image/color"
	"log/slog"

	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/direction"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	walkAnimationSteps = 8
)

type System struct {
	logger   *slog.Logger
	ecs      *ecsys.ECS
	eventBus *event.Bus
}

// New creates a new skeleton system. It listens to skeleton events and adds the necessary components to the entity to make it a skeleton.
func New(logger *slog.Logger, ecs *ecsys.ECS, eventBus *event.Bus) *System {
	s := &System{
		logger:   logger.With("system", "skeletons"),
		ecs:      ecs,
		eventBus: eventBus,
	}

	// Subscribe to the skeleton events
	s.eventBus.Subscribe(
		event.MatchAny(skeleton.UpdatedEventType, skeleton.CreatedEventType, skeleton.DeletedEventType),
		s.skeletonCreatedHandler,
	)

	return s
}

func (s *System) skeletonCreatedHandler(e event.Event) error {
	switch e := e.(type) {
	case *skeleton.CreatedEvent:
		s.logger.Debug("skeleton created", "skeleton", e.Skeleton)
		if err := s.setupSkeleton(*e.Skeleton()); err != nil {
			s.logger.Error("could not add skeleton sprite", "error", err)
		}
	case *skeleton.UpdatedEvent:
		s.logger.Debug("skeleton updated", "skeleton", e.Skeleton)
	case *skeleton.DeletedEvent:
		s.logger.Debug("skeleton deleted", "skeleton", e.Skeleton)
	default:
		return errors.New("unknown event type")
	}
	return nil
}

// setupSkeleton adds the necessary components to the entity to make it a skeleton.
func (s *System) setupSkeleton(sk skeleton.Skeleton) error {
	e := sk.Entity()
	if _, err := s.ecs.AddRectangleComponent(*shape.NewRectangle(e, 10, 10, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff})); err != nil {
		return err
	}
	if _, err := s.ecs.AddSpriteComponent(*sprite.New(e, "skeleton", sprite.SkeletonMoveDown1)); err != nil {
		return err
	}
	return nil
}

func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}

func (s *System) Update() error {
	skeletons := s.ecs.Skeletons()
	for i := range skeletons {
		e := &skeletons[i]

		pos, err := s.ecs.Position(e.Entity())
		if err != nil {
			return fmt.Errorf("could not get position: %w", err)
		}

		isMoving := false
		if e.PrefX != pos.X || e.PrefY != pos.Y {
			isMoving = true
			// Calculate facing direction before updating e.PrefX and e.PrefY
			facing := direction.Get(e.PrefX, e.PrefY, pos.X, pos.Y)
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
		e.PrefX = pos.X
		e.PrefY = pos.Y

		if e.State == skeleton.Moving {
			e.AnimationStep++
			if e.AnimationStep >= walkAnimationSteps {
				e.AnimationStep = 0
			}
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

// updateSprite updates the sprite associated with the skeleton based on its state and facing direction.
func (s *System) updateSprite(e *skeleton.Skeleton) error {
	// Retrieve the sprite component associated with the skeleton
	var spr *sprite.Sprite
	sprites := s.ecs.SpritesByEntity(e.Entity())
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
	var graphic sprite.Graphic
	if e.State == skeleton.Moving {
		// The animation step was incremented in Update; no need to increment it again here

		switch e.Facing {
		case direction.North, direction.NorthEast, direction.NorthWest:
			graphic = sprite.SkeletonMoveUpFrames()[e.AnimationStep]
		case direction.South, direction.SouthEast, direction.SouthWest:
			graphic = sprite.SkeletonMoveDownFrames()[e.AnimationStep]
		case direction.East:
			graphic = sprite.SkeletonMoveRightFrames()[e.AnimationStep]
		case direction.West:
			graphic = sprite.SkeletonMoveLeftFrames()[e.AnimationStep]
		default:
			s.logger.Warn("Unhandled facing direction", "direction", e.Facing)
			graphic = spr.Graphic // Keep current graphic if direction is unhandled
		}
	} else { // Idle state
		switch e.Facing {
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
