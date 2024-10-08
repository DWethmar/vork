package gameplay

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/game"
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
	_              game.Scene = &GamePlay{}
	sceneKey                  = []byte("gameplay")
	initializedKey            = []byte("initialized")
)

// GamePlay is a scene where the game is played.
type GamePlay struct {
	logger      *slog.Logger
	db          *bbolt.DB
	systems     []systems.System
	persistence *persistence.Persistance
}

func New(logger *slog.Logger, save string, db *bbolt.DB, s *spritesheet.Spritesheet) (*GamePlay, error) {
	eventBus := event.NewBus()
	ecs := ecsys.New(eventBus)

	systems := []systems.System{
		controller.New(logger, ecs),
		render.New(logger, sprites.Sprites(s), ecs),
		skeletons.New(logger, ecs, eventBus),
	}

	persistence := persistence.New(eventBus, ecs)

	// check if it is an existing save
	ok, err := initializedGame(db, save)
	if err != nil {
		return nil, fmt.Errorf("failed to check if game is initialized: %w", err)
	}
	if ok {
		// create a new game
		logger.Info("creating a new game")
		addPlayer(ecs, 10, 10)
		addEnemy(ecs, 100, 100)
		setInitialized(db, save)
		if err := persistence.Save(db); err != nil {
			return nil, fmt.Errorf("failed to save new game: %w", err)
		}
	} else {
		// load the game
		logger.Info("loading game")
		if err := persistence.Load(db); err != nil {
			return nil, fmt.Errorf("failed to load game: %w", err)
		}
	}

	if err != nil {
		logger.Error("failed to load game", slog.String("error", err.Error()))
	}

	return &GamePlay{
		logger:      logger,
		db:          db,
		systems:     systems,
		persistence: persistence,
	}, nil
}

func (s *GamePlay) Name() string { return "gameplay" }

func (s *GamePlay) Draw(screen *ebiten.Image) error {
	for _, sys := range s.systems {
		if err := sys.Draw(screen); err != nil {
			return err
		}
	}
	return nil
}

func (s *GamePlay) Update() error {
	// check if F5 is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		s.logger.Info("saving game")
		if err := s.persistence.Save(s.db); err != nil {
			return fmt.Errorf("failed to save game: %w", err)
		}
		s.logger.Info("game saved")
		return nil
	}

	for _, sys := range s.systems {
		if err := sys.Update(); err != nil {
			return err
		}
	}

	return nil
}

func initializedGame(db *bbolt.DB, name string) (bool, error) {
	exists := false
	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(sceneKey)
		if bucket == nil {
			return nil
		}
		exists = bucket.Get(initializedKey) != nil
		return nil
	})
	return exists, err
}

func setInitialized(db *bbolt.DB, name string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(sceneKey)
		if err != nil {
			return err
		}
		return bucket.Put(initializedKey, []byte(""))
	})
}
