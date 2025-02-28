package gameplay

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/hitbox"
	"github.com/dwethmar/vork/component/persistence"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/velocity"
	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/event/mouse"
	"github.com/dwethmar/vork/game"
	"github.com/dwethmar/vork/point"
	"github.com/dwethmar/vork/sprites"
	"github.com/dwethmar/vork/spritesheet"
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

// onClickHandler creates a click handler that publishes a clicked event.
func onClickHandler(logger *slog.Logger, eventBus *event.Bus) func(x, y int) {
	return func(x, y int) {
		if err := eventBus.Publish(mouse.NewLeftClickedEvent(x, y)); err != nil {
			logger.Error("failed to publish clicked event", slog.String("error", err.Error()))
		}
	}
}

func onHoverHandler() func(x, y int) {
	return func(x, y int) {
		ebiten.SetWindowTitle(fmt.Sprintf("vork x: %d, y: %d", x, y))
	}
}

// GamePlay is a scene where the game is played.
type GamePlay struct {
	logger                  *slog.Logger
	db                      *bbolt.DB
	systems                 []System
	stores                  *Stores
	skeletonPersistence     *persistence.Persistence[*skeleton.Skeleton]
	hitboxPersistence       *persistence.Persistence[*hitbox.Hitbox]
	velocityPersistence     *persistence.Persistence[*velocity.Velocity]
	controllablePersistence *persistence.Persistence[*controllable.Controllable]
	positionPersistence     *persistence.Persistence[*position.Position]
}

// New creates a new game play scene.
func New(logger *slog.Logger, saveName string, s *spritesheet.Spritesheet) (*GamePlay, error) {
	logger = logger.With("scene", "gameplay")
	eventBus := event.NewBus()

	savesFolder, err := getDefaultSaveFolder()
	if err != nil {
		return nil, fmt.Errorf("failed to get default save folder: %w", err)
	}

	cfg, err := LoadOrCreateConfig(saveName, savesFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.New() {
		logger.Info("creating new game", slog.String("save_name", cfg.SaveName), slog.String("db_path", cfg.DBPath))
	}

	db, err := bbolt.Open(cfg.DBPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	stores := NewStores(eventBus)

	// check if it is an existing save
	if err = setupGame(logger, stores, db); err != nil {
		return nil, fmt.Errorf("failed to setup game: %w", err)
	}

	systems := []System{
		// keyinput.New(keyinput.Options{
		// 	Logger:              logger,
		// 	ECS:                 ecs,
		// 	VelocityScaleFactor: 5,
		// }),
		render.New(render.Options{
			Logger:            logger,
			Sprites:           sprites.Sprites(s),
			PositionsStore:    stores.Positions,
			RectanglesStore:   stores.Shapes,
			SpriteStore:       stores.Sprites,
			ControllableStore: stores.Controllable,
			ClickHandler:      onClickHandler(logger, eventBus),
			HoverHandler:      onHoverHandler(),
		}),
		skeletons.New(logger, stores.Positions, stores.Skeletons, stores.Shapes, stores.Sprites, stores.Hitboxes, stores.Velocities, eventBus),
		// collision.New(collision.Options{
		// 	Logger:              logger,
		// 	ECS:                 ecs,
		// 	EventBus:            eventBus,
		// 	VelocityScaleFactor: 5,
		// 	Friction:            4,
		// 	VelocityThreshold:   1,
		// }),
	}
	// init all systems after loading the game to make sure all components are loaded
	for _, sys := range systems {
		if err = sys.Init(); err != nil {
			return nil, fmt.Errorf("failed to init system %T: %w", sys, err)
		}
	}
	return &GamePlay{
		logger:              logger,
		db:                  db,
		systems:             systems,
		stores:              stores,
		skeletonPersistence: persistence.New(skeleton.Type, stores.Skeletons),
	}, nil
}

func setupGame(logger *slog.Logger, stores *Stores, db *bbolt.DB) error {
	// check if it is an existing save
	ok, err := gameInitialized(db)
	if err != nil {
		return fmt.Errorf("failed to check if game is initialized: %w", err)
	}
	if ok {
		// // load the game
		// logger.Info("loading existing game")
		// if err = persistence.Load(db); err != nil {
		// 	return fmt.Errorf("failed to load game: %w", err)
		// }
		// if err = ecs.BuildHierarchy(); err != nil {
		// 	return fmt.Errorf("failed to rebuild hierarchy: %w", err)
		// }
		// logger.Info("game loaded")
		// return nil
	}
	// create a new game
	logger.Info("creating a new game")
	if err = initializeGame(stores, db); err != nil {
		return fmt.Errorf("failed to load game: %w", err)
	}
	// if err = persistence.Save(db); err != nil {
	// 	return fmt.Errorf("failed to save new game: %w", err)
	// }
	logger.Info("game created")
	return nil
}

// Name returns the name of the scene.
func (s *GamePlay) Name() string { return "gameplay" }

// Draw draws the game.
func (s *GamePlay) Draw(screen *ebiten.Image) error {
	for _, sys := range s.systems {
		if err := sys.Draw(screen); err != nil {
			return fmt.Errorf("failed to draw system: %w", err)
		}
	}
	return nil
}

// Update updates the game.
func (s *GamePlay) Update() error {
	// check if F5 is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		started := time.Now()
		// if err := s.persistence.Save(s.db); err != nil {
		// 	return fmt.Errorf("failed to save game: %w", err)
		// }
		s.logger.Info("game saved", slog.Duration("duration", time.Since(started)))
		return nil
	}
	for _, sys := range s.systems {
		if err := sys.Update(); err != nil {
			return fmt.Errorf("failed to update system %T: %w", sys, err)
		}
	}
	return nil
}

// Close closes the game.
func (s *GamePlay) Close() error {
	for _, sys := range s.systems {
		if err := sys.Close(); err != nil {
			return fmt.Errorf("failed to close system: %w", err)
		}
	}
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("failed to close db: %w", err)
	}
	return nil
}

// initializeGame creates a new game.
func initializeGame(stores *Stores, db *bbolt.DB) error {
	_, err := addPlayer(0, stores, point.New(10, 10))
	if err != nil {
		return fmt.Errorf("failed to add player: %w", err)
	}

	e, err := addEnemy(0, stores, point.New(100, 100))
	if err != nil {
		return fmt.Errorf("failed to add enemy %v: %w", e, err)
	}

	return db.Update(func(tx *bbolt.Tx) error {
		bucket, nErr := tx.CreateBucketIfNotExists(sceneKey)
		if nErr != nil {
			return fmt.Errorf("failed to create bucket: %w", nErr)
		}
		return bucket.Put(initializedKey, []byte(""))
	})
}

func gameInitialized(db *bbolt.DB) (bool, error) {
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
