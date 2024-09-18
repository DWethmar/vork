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
	g, err := game.New()
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
