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

// Spritesheet is a collection of sprites.
type Spritesheet struct {
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

func loadSkeletonDeath(s *Spritesheet) error {
	img, err := LoadPng(imagesFS, "images/skeleton_death.png")
	if err != nil {
		return fmt.Errorf("failed to load images/skeleton_death.png: %w", err)
	}
	cells := CreateRectangleGrid(6, 1, skeletonDeathWidth, skeletonDeathHeight)
	skeletonKillImg := ebiten.NewImageFromImage(img)
	s.SkeletonDeath1 = ebiten.NewImageFromImage(skeletonKillImg.SubImage(cells[0][0]))
	s.SkeletonDeath2 = ebiten.NewImageFromImage(skeletonKillImg.SubImage(cells[1][0]))
	s.SkeletonDeath3 = ebiten.NewImageFromImage(skeletonKillImg.SubImage(cells[2][0]))
	s.SkeletonDeath4 = ebiten.NewImageFromImage(skeletonKillImg.SubImage(cells[3][0]))
	s.SkeletonDeath5 = ebiten.NewImageFromImage(skeletonKillImg.SubImage(cells[4][0]))
	s.SkeletonDeath6 = ebiten.NewImageFromImage(skeletonKillImg.SubImage(cells[5][0]))
	return nil
}

func loadSkeletonMove(s *Spritesheet) error {
	img, err := LoadPng(imagesFS, "images/skeleton_move.png")
	if err != nil {
		return fmt.Errorf("failed to load images/skeleton_move.png: %w", err)
	}
	cells := CreateRectangleGrid(9, 4, skeletonMoveWidth, skeletonMoveHeight)
	skeletonMoveImg := ebiten.NewImageFromImage(img)
	// Skeleton moving up
	s.SkeletonMoveUp1 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[0][0]))
	s.SkeletonMoveUp2 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[1][0]))
	s.SkeletonMoveUp3 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[2][0]))
	s.SkeletonMoveUp4 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[3][0]))
	s.SkeletonMoveUp5 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[4][0]))
	s.SkeletonMoveUp6 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[5][0]))
	s.SkeletonMoveUp7 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[6][0]))
	s.SkeletonMoveUp8 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[7][0]))
	s.SkeletonMoveUp9 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[8][0]))
	// Skeleton moving left
	s.SkeletonMoveLeft1 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[0][1]))
	s.SkeletonMoveLeft2 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[1][1]))
	s.SkeletonMoveLeft3 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[2][1]))
	s.SkeletonMoveLeft4 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[3][1]))
	s.SkeletonMoveLeft5 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[4][1]))
	s.SkeletonMoveLeft6 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[5][1]))
	s.SkeletonMoveLeft7 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[6][1]))
	s.SkeletonMoveLeft8 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[7][1]))
	s.SkeletonMoveLeft9 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[8][1]))
	// Skeleton moving down
	s.SkeletonMoveDown1 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[0][2]))
	s.SkeletonMoveDown2 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[1][2]))
	s.SkeletonMoveDown3 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[2][2]))
	s.SkeletonMoveDown4 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[3][2]))
	s.SkeletonMoveDown5 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[4][2]))
	s.SkeletonMoveDown6 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[5][2]))
	s.SkeletonMoveDown7 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[6][2]))
	s.SkeletonMoveDown8 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[7][2]))
	s.SkeletonMoveDown9 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[8][2]))
	// Skeleton moving right
	s.SkeletonMoveRight1 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[0][3]))
	s.SkeletonMoveRight2 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[1][3]))
	s.SkeletonMoveRight3 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[2][3]))
	s.SkeletonMoveRight4 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[3][3]))
	s.SkeletonMoveRight5 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[4][3]))
	s.SkeletonMoveRight6 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[5][3]))
	s.SkeletonMoveRight7 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[6][3]))
	s.SkeletonMoveRight8 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[7][3]))
	s.SkeletonMoveRight9 = ebiten.NewImageFromImage(skeletonMoveImg.SubImage(cells[8][3]))
	return nil
}

func New() (*Spritesheet, error) {
	s := &Spritesheet{}
	if err := loadSkeletonDeath(s); err != nil {
		return nil, fmt.Errorf("failed to skeleton death sprites: %w", err)
	}
	if err := loadSkeletonMove(s); err != nil {
		return nil, fmt.Errorf("failed to load skeleton move sprites: %w", err)
	}
	return s, nil
}
