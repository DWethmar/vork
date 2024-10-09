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
	"go.etcd.io/bbolt"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

func main() {
	logger := slog.Default()
	spriteSheet, err := spritesheet.New()
	if err != nil {
		log.Fatal(fmt.Printf("failed to create sprite sheet: %v", err))
	}

	// ensure .tmp directory exists
	if err := ensureFolderExists(".tmp"); err != nil {
		log.Fatal(fmt.Printf("failed to ensure folder exists: %v", err))
	}

	db, err := bbolt.Open(".tmp/vork.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(fmt.Errorf("failed to close db: %w", err))
		}
	}()

	gameplayScene, err := gameplay.New(logger, "my-save", db, spriteSheet)
	if err != nil {
		log.Fatal(fmt.Printf("failed to create gameplay scene: %v", err))
	}

	g, err := game.New()
	if err != nil {
		log.Fatal(fmt.Printf("failed to create new game: %v", err))
	}

	// Add scenes
	if err := g.AddScene(gameplayScene); err != nil {
		log.Fatal(fmt.Printf("failed to add scene %s: %v", gameplayScene.Name(), err))
	}

	// Switch to the gameplay scene
	if err := g.SwitchScene(gameplayScene.Name()); err != nil {
		log.Fatal(fmt.Printf("failed to switch to scene %s: %v", gameplayScene.Name(), err))
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("vorK")
	if err := ebiten.RunGame(&EbitenGame{
		g: g,
	}); err != nil {
		log.Fatal(err)
	}
}

func ensureFolderExists(path string) error {
	// Check if the directory already exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		fmt.Printf("Directory %s created.\n", path)
	} else {
		fmt.Printf("Directory %s already exists.\n", path)
	}
	return nil
}
