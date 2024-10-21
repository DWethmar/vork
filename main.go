package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/dwethmar/vork/game"
	"github.com/dwethmar/vork/game/scenes/gameplay"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	logger := slog.Default()
	spriteSheet, err := spritesheet.New()
	if err != nil {
		return fmt.Errorf("failed to create sprite sheet: %w", err)
	}

	// ensure .tmp directory exists
	if err = ensureFolderExists(".tmp"); err != nil {
		return fmt.Errorf("failed to ensure folder exists: %w", err)
	}

	gameplayScene, err := gameplay.New(logger, "my-first-game", spriteSheet)
	if err != nil {
		return fmt.Errorf("failed to create gameplay scene: %w", err)
	}

	g, err := game.New()
	if err != nil {
		return fmt.Errorf("failed to create new game: %w", err)
	}

	// Add scenes
	if err = g.AddScene(gameplayScene); err != nil {
		return fmt.Errorf("failed to add scene %s: %w", gameplayScene.Name(), err)
	}

	// Switch to the gameplay scene
	if err = g.SwitchScene(gameplayScene.Name()); err != nil {
		return fmt.Errorf("failed to switch to scene %s: %w", gameplayScene.Name(), err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("vorK")
	if err = ebiten.RunGame(&EbitenGame{
		g: g,
	}); err != nil {
		return fmt.Errorf("failed to run game: %w", err)
	}

	return nil
}

func ensureFolderExists(path string) error {
	// Check if the directory already exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		fmt.Printf("Directory %s created.\n", path)
	} else {
		fmt.Printf("Directory %s already exists.\n", path)
	}
	return nil
}
