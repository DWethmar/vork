package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/dwethmar/vork/game/skelebork"
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

	skelebork, err := skelebork.NewGame(logger)
	if err != nil {
		return fmt.Errorf("failed to create skelebork game: %w", err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("vorK")
	if err = ebiten.RunGame(skelebork); err != nil {
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
