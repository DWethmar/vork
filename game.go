package main

import (
	"github.com/dwethmar/vork/game"
	"github.com/hajimehoshi/ebiten/v2"
)

// EbitenGame is the game struct for ebiten.
type EbitenGame struct {
	g *game.Game
}

func (e *EbitenGame) Update() error {
	return e.g.Update()
}

func (e *EbitenGame) Draw(screen *ebiten.Image) {
	if err := e.g.Draw(screen); err != nil {
		panic(err)
	}
}

func (e *EbitenGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
