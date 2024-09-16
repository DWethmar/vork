package spritesheet

import (
	"embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	skeletonDeathHeight = 64
	skeletonDeathWidth  = 64
)

//go:embed images
var imagesFS embed.FS

type Sprites struct {
	SkeletonDeath1 *ebiten.Image
	SkeletonDeath2 *ebiten.Image
	SkeletonDeath3 *ebiten.Image
	SkeletonDeath4 *ebiten.Image
	SkeletonDeath5 *ebiten.Image
	SkeletonDeath6 *ebiten.Image
}

func loadSkeletonDeath(s *Sprites) error {
	img, err := loadPng(imagesFS, "images/skeleton_death.png")
	if err != nil {
		return fmt.Errorf("failed to load images/skeleton_death.png: %w", err)
	}

	SkeletonKillImg := ebiten.NewImageFromImage(img)
	cells := CreateRectangleGrid(6, 1, skeletonDeathWidth, skeletonDeathHeight)

	s.SkeletonDeath1 = SkeletonKillImg.SubImage(cells[0][0]).(*ebiten.Image)
	s.SkeletonDeath2 = SkeletonKillImg.SubImage(cells[1][0]).(*ebiten.Image)
	s.SkeletonDeath3 = SkeletonKillImg.SubImage(cells[2][0]).(*ebiten.Image)
	s.SkeletonDeath4 = SkeletonKillImg.SubImage(cells[3][0]).(*ebiten.Image)
	s.SkeletonDeath5 = SkeletonKillImg.SubImage(cells[4][0]).(*ebiten.Image)
	s.SkeletonDeath6 = SkeletonKillImg.SubImage(cells[5][0]).(*ebiten.Image)

	return nil
}

func New() (*Sprites, error) {
	s := &Sprites{}
	if err := loadSkeletonDeath(s); err != nil {
		return nil, fmt.Errorf("failed to load sprites: %w", err)
	}
	return s, nil
}
