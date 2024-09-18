package spritesheet

import (
	"embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	skeletonDeathHeight = 64
	skeletonDeathWidth  = 64
	skeletonMoveHeight  = 64
	skeletonMoveWidth   = 64
)

//go:embed images
var imagesFS embed.FS

// Sprites is a collection of sprites.
type Sprites struct {
	SkeletonDeath1 *ebiten.Image
	SkeletonDeath2 *ebiten.Image
	SkeletonDeath3 *ebiten.Image
	SkeletonDeath4 *ebiten.Image
	SkeletonDeath5 *ebiten.Image
	SkeletonDeath6 *ebiten.Image
	// Skeleton moving up
	SkeletonMoveUp1 *ebiten.Image // idle
	SkeletonMoveUp2 *ebiten.Image
	SkeletonMoveUp3 *ebiten.Image
	SkeletonMoveUp4 *ebiten.Image
	SkeletonMoveUp5 *ebiten.Image
	SkeletonMoveUp6 *ebiten.Image
	SkeletonMoveUp7 *ebiten.Image
	SkeletonMoveUp8 *ebiten.Image
	SkeletonMoveUp9 *ebiten.Image
	// Skeleton moving left
	SkeletonMoveLeft1 *ebiten.Image // idle
	SkeletonMoveLeft2 *ebiten.Image
	SkeletonMoveLeft3 *ebiten.Image
	SkeletonMoveLeft4 *ebiten.Image
	SkeletonMoveLeft5 *ebiten.Image
	SkeletonMoveLeft6 *ebiten.Image
	SkeletonMoveLeft7 *ebiten.Image
	SkeletonMoveLeft8 *ebiten.Image
	SkeletonMoveLeft9 *ebiten.Image
	// Skeleton moving down
	SkeletonMoveDown1 *ebiten.Image // idle
	SkeletonMoveDown2 *ebiten.Image
	SkeletonMoveDown3 *ebiten.Image
	SkeletonMoveDown4 *ebiten.Image
	SkeletonMoveDown5 *ebiten.Image
	SkeletonMoveDown6 *ebiten.Image
	SkeletonMoveDown7 *ebiten.Image
	SkeletonMoveDown8 *ebiten.Image
	SkeletonMoveDown9 *ebiten.Image
	// Skeleton moving right
	SkeletonMoveRight1 *ebiten.Image // idle
	SkeletonMoveRight2 *ebiten.Image
	SkeletonMoveRight3 *ebiten.Image
	SkeletonMoveRight4 *ebiten.Image
	SkeletonMoveRight5 *ebiten.Image
	SkeletonMoveRight6 *ebiten.Image
	SkeletonMoveRight7 *ebiten.Image
	SkeletonMoveRight8 *ebiten.Image
	SkeletonMoveRight9 *ebiten.Image
}

func loadSkeletonDeath(s *Sprites) error {
	img, err := loadPng(imagesFS, "images/skeleton_death.png")
	if err != nil {
		return fmt.Errorf("failed to load images/skeleton_death.png: %w", err)
	}
	cells := CreateRectangleGrid(6, 1, skeletonDeathWidth, skeletonDeathHeight)
	skeletonKillImg := ebiten.NewImageFromImage(img)
	s.SkeletonDeath1 = skeletonKillImg.SubImage(cells[0][0]).(*ebiten.Image)
	s.SkeletonDeath2 = skeletonKillImg.SubImage(cells[1][0]).(*ebiten.Image)
	s.SkeletonDeath3 = skeletonKillImg.SubImage(cells[2][0]).(*ebiten.Image)
	s.SkeletonDeath4 = skeletonKillImg.SubImage(cells[3][0]).(*ebiten.Image)
	s.SkeletonDeath5 = skeletonKillImg.SubImage(cells[4][0]).(*ebiten.Image)
	s.SkeletonDeath6 = skeletonKillImg.SubImage(cells[5][0]).(*ebiten.Image)
	return nil
}

func loadSkeletonMove(s *Sprites) error {
	img, err := loadPng(imagesFS, "images/skeleton_move.png")
	if err != nil {
		return fmt.Errorf("failed to load images/skeleton_move.png: %w", err)
	}
	cells := CreateRectangleGrid(9, 4, skeletonMoveWidth, skeletonMoveHeight)
	skeletonMoveImg := ebiten.NewImageFromImage(img)
	// Skeleton moving up
	s.SkeletonMoveUp1 = skeletonMoveImg.SubImage(cells[0][0]).(*ebiten.Image)
	s.SkeletonMoveUp2 = skeletonMoveImg.SubImage(cells[1][0]).(*ebiten.Image)
	s.SkeletonMoveUp3 = skeletonMoveImg.SubImage(cells[2][0]).(*ebiten.Image)
	s.SkeletonMoveUp4 = skeletonMoveImg.SubImage(cells[3][0]).(*ebiten.Image)
	s.SkeletonMoveUp5 = skeletonMoveImg.SubImage(cells[4][0]).(*ebiten.Image)
	s.SkeletonMoveUp6 = skeletonMoveImg.SubImage(cells[5][0]).(*ebiten.Image)
	s.SkeletonMoveUp7 = skeletonMoveImg.SubImage(cells[6][0]).(*ebiten.Image)
	s.SkeletonMoveUp8 = skeletonMoveImg.SubImage(cells[7][0]).(*ebiten.Image)
	s.SkeletonMoveUp9 = skeletonMoveImg.SubImage(cells[8][0]).(*ebiten.Image)
	// Skeleton moving left
	s.SkeletonMoveLeft1 = skeletonMoveImg.SubImage(cells[0][1]).(*ebiten.Image)
	s.SkeletonMoveLeft2 = skeletonMoveImg.SubImage(cells[1][1]).(*ebiten.Image)
	s.SkeletonMoveLeft3 = skeletonMoveImg.SubImage(cells[2][1]).(*ebiten.Image)
	s.SkeletonMoveLeft4 = skeletonMoveImg.SubImage(cells[3][1]).(*ebiten.Image)
	s.SkeletonMoveLeft5 = skeletonMoveImg.SubImage(cells[4][1]).(*ebiten.Image)
	s.SkeletonMoveLeft6 = skeletonMoveImg.SubImage(cells[5][1]).(*ebiten.Image)
	s.SkeletonMoveLeft7 = skeletonMoveImg.SubImage(cells[6][1]).(*ebiten.Image)
	s.SkeletonMoveLeft8 = skeletonMoveImg.SubImage(cells[7][1]).(*ebiten.Image)
	s.SkeletonMoveLeft9 = skeletonMoveImg.SubImage(cells[8][1]).(*ebiten.Image)
	// Skeleton moving down
	s.SkeletonMoveDown1 = skeletonMoveImg.SubImage(cells[0][2]).(*ebiten.Image)
	s.SkeletonMoveDown2 = skeletonMoveImg.SubImage(cells[1][2]).(*ebiten.Image)
	s.SkeletonMoveDown3 = skeletonMoveImg.SubImage(cells[2][2]).(*ebiten.Image)
	s.SkeletonMoveDown4 = skeletonMoveImg.SubImage(cells[3][2]).(*ebiten.Image)
	s.SkeletonMoveDown5 = skeletonMoveImg.SubImage(cells[4][2]).(*ebiten.Image)
	s.SkeletonMoveDown6 = skeletonMoveImg.SubImage(cells[5][2]).(*ebiten.Image)
	s.SkeletonMoveDown7 = skeletonMoveImg.SubImage(cells[6][2]).(*ebiten.Image)
	s.SkeletonMoveDown8 = skeletonMoveImg.SubImage(cells[7][2]).(*ebiten.Image)
	s.SkeletonMoveDown9 = skeletonMoveImg.SubImage(cells[8][2]).(*ebiten.Image)
	// Skeleton moving right
	s.SkeletonMoveRight1 = skeletonMoveImg.SubImage(cells[0][3]).(*ebiten.Image)
	s.SkeletonMoveRight2 = skeletonMoveImg.SubImage(cells[1][3]).(*ebiten.Image)
	s.SkeletonMoveRight3 = skeletonMoveImg.SubImage(cells[2][3]).(*ebiten.Image)
	s.SkeletonMoveRight4 = skeletonMoveImg.SubImage(cells[3][3]).(*ebiten.Image)
	s.SkeletonMoveRight5 = skeletonMoveImg.SubImage(cells[4][3]).(*ebiten.Image)
	s.SkeletonMoveRight6 = skeletonMoveImg.SubImage(cells[5][3]).(*ebiten.Image)
	s.SkeletonMoveRight7 = skeletonMoveImg.SubImage(cells[6][3]).(*ebiten.Image)
	s.SkeletonMoveRight8 = skeletonMoveImg.SubImage(cells[7][3]).(*ebiten.Image)
	s.SkeletonMoveRight9 = skeletonMoveImg.SubImage(cells[8][3]).(*ebiten.Image)
	return nil
}

func New() (*Sprites, error) {
	s := &Sprites{}
	if err := loadSkeletonDeath(s); err != nil {
		return nil, fmt.Errorf("failed to skeleton death sprites: %w", err)
	}
	if err := loadSkeletonMove(s); err != nil {
		return nil, fmt.Errorf("failed to load skeleton move sprites: %w", err)
	}
	return s, nil
}
