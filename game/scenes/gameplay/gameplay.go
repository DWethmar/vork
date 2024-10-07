package gameplay

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/persistence"
	"github.com/dwethmar/vork/sprites"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems"
	"github.com/dwethmar/vork/systems/controller"
	"github.com/dwethmar/vork/systems/render"
	"github.com/dwethmar/vork/systems/skeletons"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go.etcd.io/bbolt"
)

var (
	sceneKey     = []byte("gameplay")
	startedAtKey = []byte("started_at")
)

type Scene struct {
	logger      *slog.Logger
	db          *bbolt.DB
	systems     []systems.System
	persistence *persistence.Persistance
}

func NewScene(logger *slog.Logger, db *bbolt.DB, s *spritesheet.Spritesheet) *Scene {
	eventBus := event.NewBus()
	ecs := ecsys.New(eventBus)

	systems := []systems.System{
		controller.New(logger, ecs),
		render.New(logger, sprites.Sprites(s), ecs),
		skeletons.New(logger, ecs, eventBus),
	}

	persistence := persistence.New(eventBus, ecs)

	// check if it is an existing save
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(sceneKey)
		if err != nil {
			return err
		}
		startedAt := bucket.Get(startedAtKey)
		if startedAt == nil {
			// create a new game
			logger.Info("creating a new game")
			addPlayer(ecs, 10, 10)
			addEnemy(ecs, 100, 100)
			if err := bucket.Put(startedAtKey, []byte(time.Now().String())); err != nil {
				return err
			}
			if err := persistence.Save(tx); err != nil {
				return fmt.Errorf("failed to save game: %w", err)
			}
		} else {
			// load the game
			logger.Info("loading game")
			if err := persistence.Load(tx); err != nil {
				return fmt.Errorf("failed to load game: %w", err)
			}
		}
		return nil
	})

	if err != nil {
		logger.Error("failed to load game", slog.String("error", err.Error()))
	}

	return &Scene{
		logger:      logger,
		db:          db,
		systems:     systems,
		persistence: persistence,
	}
}

func (s *Scene) Name() string { return "gameplay" }

func (s *Scene) Draw(screen *ebiten.Image) error {
	for _, sys := range s.systems {
		if err := sys.Draw(screen); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scene) Update() error {
	// check if F5 is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		s.logger.Info("saving game")
		s.db.Update(func(tx *bbolt.Tx) error {
			if err := s.persistence.Save(tx); err != nil {
				return fmt.Errorf("failed to save game: %w", err)
			}
			return nil
		})
		s.logger.Info("game saved")
	}

	for _, sys := range s.systems {
		if err := sys.Update(); err != nil {
			return err
		}
	}

	return nil
}
