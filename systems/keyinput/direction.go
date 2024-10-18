package keyinput

import "github.com/hajimehoshi/ebiten/v2"

var left = []ebiten.Key{
	ebiten.KeyA,
	ebiten.KeyLeft,
}

var right = []ebiten.Key{
	ebiten.KeyD,
	ebiten.KeyRight,
}

var up = []ebiten.Key{
	ebiten.KeyW,
	ebiten.KeyUp,
}

var down = []ebiten.Key{
	ebiten.KeyS,
	ebiten.KeyDown,
}

// direction returns the x and y direction of the keys pressed.
func direction() (int, int) {
	var x, y int
	for _, k := range left {
		if ebiten.IsKeyPressed(k) {
			x--
			break
		}
	}

	for _, k := range right {
		if ebiten.IsKeyPressed(k) {
			x++
			break
		}
	}

	for _, k := range up {
		if ebiten.IsKeyPressed(k) {
			y--
			break
		}
	}

	for _, k := range down {
		if ebiten.IsKeyPressed(k) {
			y++
			break
		}
	}

	return x, y
}
