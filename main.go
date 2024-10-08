package main

import (
	"fmt"
	"log"
	"log/slog"

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
		log.Fatal(err)
	}

	db, err := bbolt.Open("vork.db", 0600, nil)
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
		log.Fatal(err)
	}

	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}

	// Add scenes
	if err := g.AddScene(gameplayScene); err != nil {
		log.Fatal(err)
	}

	// Switch to the gameplay scene
	if err := g.SwitchScene(gameplayScene.Name()); err != nil {
		log.Fatal(err)
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
