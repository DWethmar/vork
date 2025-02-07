package skelebork

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/vork/game"
	"github.com/dwethmar/vork/game/scenes/gameplay"
	"github.com/dwethmar/vork/spritesheet"
)

func NewGame(logger *slog.Logger) (*game.Game, error) {
	spriteSheet, err := spritesheet.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create sprite sheet: %w", err)
	}

	gameplayScene, err := gameplay.New(logger, "my-first-game", spriteSheet)
	if err != nil {
		return nil, fmt.Errorf("failed to create gameplay scene: %w", err)
	}

	g := game.New(logger)

	// Add scenes
	if err = g.AddScene(gameplayScene); err != nil {
		return nil, fmt.Errorf("failed to add scene %s: %w", gameplayScene.Name(), err)
	}

	// Switch to the gameplay scene
	if err = g.SwitchScene(gameplayScene.Name()); err != nil {
		return nil, fmt.Errorf("failed to switch to scene %s: %w", gameplayScene.Name(), err)
	}

	return g, nil
}
