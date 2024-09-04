package main

import (
	"log"

	"github.com/dwethmar/vork/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	game := &Game{
		g: game.New(),
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("vorK")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
