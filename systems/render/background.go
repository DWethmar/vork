package render

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	normalFontSize = 12
)

//go:embed mplus-1p-regular.ttf
var fontData []byte
var mplusFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fontData))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

// renderGrid draws a checkered grid background onto the provided image.
// It accounts for the camera's offset and the current zoom level to ensure
// that the grid aligns correctly with the game world.
func renderGrid(s *ebiten.Image, offsetX, offsetY int, zoom float64, debug bool) error {
	screenWidth := s.Bounds().Dx()
	screenHeight := s.Bounds().Dy()
	gridSize := 40

	offsetXf := float64(offsetX)
	offsetYf := float64(offsetY)

	numGridX := int(float64(screenWidth)/(float64(gridSize)*zoom)) + 2
	numGridY := int(float64(screenHeight)/(float64(gridSize)*zoom)) + 2

	startX := int(offsetXf / float64(gridSize))
	startY := int(offsetYf / float64(gridSize))

	// Adjust font size based on zoom
	fontSize := normalFontSize * zoom
	// Set minimum and maximum font sizes
	if fontSize < 8 {
		fontSize = 8 // Minimum font size
	} else if fontSize > 72 {
		fontSize = 72 // Maximum font size
	}

	ff := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   fontSize,
	}

	for x := startX - 1; x < startX+numGridX; x++ {
		for y := startY - 1; y < startY+numGridY; y++ {

			dx := (float64(x*gridSize) - offsetXf) * zoom
			dy := (float64(y*gridSize) - offsetYf) * zoom
			size := float64(gridSize) * zoom

			if dx+size < 0 || dx > float64(screenWidth) ||
				dy+size < 0 || dy > float64(screenHeight) {
				continue
			}

			var clrA color.Color
			var clrB color.Color
			if (x+y)%2 == 0 {
				clrA = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff} // Black
				clrB = color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff} // Dark Gray
			} else {
				clrA = color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff} // Dark Gray
				clrB = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff} // Black
			}

			vector.DrawFilledRect(s, float32(dx), float32(dy), float32(size), float32(size), clrA, true)

			// Draw the tile coordinates
			if zoom < 1 || !debug {
				continue
			}

			// Draw the text onto the image
			op := &text.DrawOptions{}
			op.GeoM.Translate(dx, dy)
			op.ColorScale.ScaleWithColor(clrB)
			text.Draw(s, fmt.Sprintf("%d,%d", x, y), ff, op)
		}
	}

	return nil
}
