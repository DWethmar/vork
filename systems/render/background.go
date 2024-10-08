package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func renderGrid(s *ebiten.Image, offsetX, offsetY int) {
	// Draw checkered background
	for x := 0; x < 800; x += 40 {
		for y := 0; y < 600; y += 40 {
			dx := float32(x - offsetX)
			dy := float32(y - offsetY)
			if (x/40+y/40)%2 == 0 {
				vector.DrawFilledRect(s, dx, dy, 40, 40, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, true)
			} else {
				vector.DrawFilledRect(s, dx, dy, 40, 40, color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}, true)
			}
		}
	}
}
