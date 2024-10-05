package skeletons

import (
	"errors"
	"image/color"
	"log/slog"

	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/shape"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/hajimehoshi/ebiten/v2"
)

type System struct {
	logger   *slog.Logger
	ecs      *ecsys.ECS
	eventBus *event.Bus
}

func New(logger *slog.Logger, ecs *ecsys.ECS, eventBus *event.Bus) *System {
	s := &System{
		logger:   logger.With("system", "skeletons"),
		ecs:      ecs,
		eventBus: eventBus,
	}

	// Subscribe to the skeleton events
	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		return e.Event() == skeleton.CreatedEventType
	}), s.skeletonCreatedHandler)

	s.eventBus.Subscribe(event.MatcherFunc(func(e event.Event) bool {
		return e.Event() == skeleton.UpdatedEventType
	}), s.skeletonCreatedHandler)

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
	default:
		return errors.New("unknown event type")
	}
	return nil
}

// setupSkeleton adds the necessary components to the entity to make it a skeleton.
func (s *System) setupSkeleton(sk skeleton.Skeleton) error {
	e := sk.Entity()
	_, err := s.ecs.Position(e)
	if err != nil {
		if errors.Is(err, ecsys.ErrNotFound) {
			s.ecs.AddPosition(*position.New(e, 0, 0))
		}
		return err
	}
	s.ecs.AddRectangle(*shape.NewRectangle(e, 10, 10, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}))
	s.ecs.AddSprite(*sprite.New(e, sprite.SkeletonMoveDown1))
	return nil
}

func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}

func (s *System) Update() error {
	return nil
}