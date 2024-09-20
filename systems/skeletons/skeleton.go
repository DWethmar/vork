package skeletons

import (
	"log/slog"

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
	return &System{
		logger:   logger,
		ecs:      ecs,
		eventBus: eventBus,
	}
}

func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}

func (s *System) Update() error {
	return nil
}
