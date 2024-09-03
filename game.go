package main

import (
	"github.com/dwethmar/vork/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	g *game.Game
}

func (g *Game) Update() error {
	return g.g.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	if err := g.g.Draw(screen); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
