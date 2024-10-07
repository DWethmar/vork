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
	screenWidth  = 320
	screenHeight = 240
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
			fmt.Errorf("error closing db: %v", err)
		}
	}()

	gameplayScene := gameplay.NewScene(logger, db, spriteSheet)

	g, err := game.New(map[string]game.Scene{
		gameplayScene.Name(): gameplayScene,
	}, gameplayScene.Name())

	if err != nil {
		log.Fatal(err)
	}
	e := &EbitenGame{
		g: g,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("vorK")
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}
