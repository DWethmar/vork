package game

import (
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

func calculateOffsetX(img *ebiten.Image) int {
	return -(img.Bounds().Dx() / 2)
}

func calculateOffsetY(img *ebiten.Image) int {
	return -(img.Bounds().Dy())
}

func Sprites(s *spritesheet.Spritesheet) []render.Sprite {
	return []render.Sprite{
		// Skeleton Death Sprites
		{Graphic: sprite.SkeletonDeath1, Img: s.SkeletonDeath1, OffsetX: calculateOffsetX(s.SkeletonDeath1), OffsetY: calculateOffsetY(s.SkeletonDeath1)},
		{Graphic: sprite.SkeletonDeath2, Img: s.SkeletonDeath2, OffsetX: calculateOffsetX(s.SkeletonDeath2), OffsetY: calculateOffsetY(s.SkeletonDeath2)},
		{Graphic: sprite.SkeletonDeath3, Img: s.SkeletonDeath3, OffsetX: calculateOffsetX(s.SkeletonDeath3), OffsetY: calculateOffsetY(s.SkeletonDeath3)},
		{Graphic: sprite.SkeletonDeath4, Img: s.SkeletonDeath4, OffsetX: calculateOffsetX(s.SkeletonDeath4), OffsetY: calculateOffsetY(s.SkeletonDeath4)},
		{Graphic: sprite.SkeletonDeath5, Img: s.SkeletonDeath5, OffsetX: calculateOffsetX(s.SkeletonDeath5), OffsetY: calculateOffsetY(s.SkeletonDeath5)},
		{Graphic: sprite.SkeletonDeath6, Img: s.SkeletonDeath6, OffsetX: calculateOffsetX(s.SkeletonDeath6), OffsetY: calculateOffsetY(s.SkeletonDeath6)},
		// Skeleton Move Up Sprites
		{Graphic: sprite.SkeletonMoveUp1, Img: s.SkeletonMoveUp1, OffsetX: calculateOffsetX(s.SkeletonMoveUp1), OffsetY: calculateOffsetY(s.SkeletonMoveUp1)},
		{Graphic: sprite.SkeletonMoveUp2, Img: s.SkeletonMoveUp2, OffsetX: calculateOffsetX(s.SkeletonMoveUp2), OffsetY: calculateOffsetY(s.SkeletonMoveUp2)},
		{Graphic: sprite.SkeletonMoveUp3, Img: s.SkeletonMoveUp3, OffsetX: calculateOffsetX(s.SkeletonMoveUp3), OffsetY: calculateOffsetY(s.SkeletonMoveUp3)},
		{Graphic: sprite.SkeletonMoveUp4, Img: s.SkeletonMoveUp4, OffsetX: calculateOffsetX(s.SkeletonMoveUp4), OffsetY: calculateOffsetY(s.SkeletonMoveUp4)},
		{Graphic: sprite.SkeletonMoveUp5, Img: s.SkeletonMoveUp5, OffsetX: calculateOffsetX(s.SkeletonMoveUp5), OffsetY: calculateOffsetY(s.SkeletonMoveUp5)},
		{Graphic: sprite.SkeletonMoveUp6, Img: s.SkeletonMoveUp6, OffsetX: calculateOffsetX(s.SkeletonMoveUp6), OffsetY: calculateOffsetY(s.SkeletonMoveUp6)},
		{Graphic: sprite.SkeletonMoveUp7, Img: s.SkeletonMoveUp7, OffsetX: calculateOffsetX(s.SkeletonMoveUp7), OffsetY: calculateOffsetY(s.SkeletonMoveUp7)},
		{Graphic: sprite.SkeletonMoveUp8, Img: s.SkeletonMoveUp8, OffsetX: calculateOffsetX(s.SkeletonMoveUp8), OffsetY: calculateOffsetY(s.SkeletonMoveUp8)},
		{Graphic: sprite.SkeletonMoveUp9, Img: s.SkeletonMoveUp9, OffsetX: calculateOffsetX(s.SkeletonMoveUp9), OffsetY: calculateOffsetY(s.SkeletonMoveUp9)},
		// Skeleton Move Left Sprites
		{Graphic: sprite.SkeletonMoveLeft1, Img: s.SkeletonMoveLeft1, OffsetX: calculateOffsetX(s.SkeletonMoveLeft1), OffsetY: calculateOffsetY(s.SkeletonMoveLeft1)},
		{Graphic: sprite.SkeletonMoveLeft2, Img: s.SkeletonMoveLeft2, OffsetX: calculateOffsetX(s.SkeletonMoveLeft2), OffsetY: calculateOffsetY(s.SkeletonMoveLeft2)},
		{Graphic: sprite.SkeletonMoveLeft3, Img: s.SkeletonMoveLeft3, OffsetX: calculateOffsetX(s.SkeletonMoveLeft3), OffsetY: calculateOffsetY(s.SkeletonMoveLeft3)},
		{Graphic: sprite.SkeletonMoveLeft4, Img: s.SkeletonMoveLeft4, OffsetX: calculateOffsetX(s.SkeletonMoveLeft4), OffsetY: calculateOffsetY(s.SkeletonMoveLeft4)},
		{Graphic: sprite.SkeletonMoveLeft5, Img: s.SkeletonMoveLeft5, OffsetX: calculateOffsetX(s.SkeletonMoveLeft5), OffsetY: calculateOffsetY(s.SkeletonMoveLeft5)},
		{Graphic: sprite.SkeletonMoveLeft6, Img: s.SkeletonMoveLeft6, OffsetX: calculateOffsetX(s.SkeletonMoveLeft6), OffsetY: calculateOffsetY(s.SkeletonMoveLeft6)},
		{Graphic: sprite.SkeletonMoveLeft7, Img: s.SkeletonMoveLeft7, OffsetX: calculateOffsetX(s.SkeletonMoveLeft7), OffsetY: calculateOffsetY(s.SkeletonMoveLeft7)},
		{Graphic: sprite.SkeletonMoveLeft8, Img: s.SkeletonMoveLeft8, OffsetX: calculateOffsetX(s.SkeletonMoveLeft8), OffsetY: calculateOffsetY(s.SkeletonMoveLeft8)},
		{Graphic: sprite.SkeletonMoveLeft9, Img: s.SkeletonMoveLeft9, OffsetX: calculateOffsetX(s.SkeletonMoveLeft9), OffsetY: calculateOffsetY(s.SkeletonMoveLeft9)},
		// Skeleton Move Down Sprites
		{Graphic: sprite.SkeletonMoveDown1, Img: s.SkeletonMoveDown1, OffsetX: calculateOffsetX(s.SkeletonMoveDown1), OffsetY: calculateOffsetY(s.SkeletonMoveDown1)},
		{Graphic: sprite.SkeletonMoveDown2, Img: s.SkeletonMoveDown2, OffsetX: calculateOffsetX(s.SkeletonMoveDown2), OffsetY: calculateOffsetY(s.SkeletonMoveDown2)},
		{Graphic: sprite.SkeletonMoveDown3, Img: s.SkeletonMoveDown3, OffsetX: calculateOffsetX(s.SkeletonMoveDown3), OffsetY: calculateOffsetY(s.SkeletonMoveDown3)},
		{Graphic: sprite.SkeletonMoveDown4, Img: s.SkeletonMoveDown4, OffsetX: calculateOffsetX(s.SkeletonMoveDown4), OffsetY: calculateOffsetY(s.SkeletonMoveDown4)},
		{Graphic: sprite.SkeletonMoveDown5, Img: s.SkeletonMoveDown5, OffsetX: calculateOffsetX(s.SkeletonMoveDown5), OffsetY: calculateOffsetY(s.SkeletonMoveDown5)},
		{Graphic: sprite.SkeletonMoveDown6, Img: s.SkeletonMoveDown6, OffsetX: calculateOffsetX(s.SkeletonMoveDown6), OffsetY: calculateOffsetY(s.SkeletonMoveDown6)},
		{Graphic: sprite.SkeletonMoveDown7, Img: s.SkeletonMoveDown7, OffsetX: calculateOffsetX(s.SkeletonMoveDown7), OffsetY: calculateOffsetY(s.SkeletonMoveDown7)},
		{Graphic: sprite.SkeletonMoveDown8, Img: s.SkeletonMoveDown8, OffsetX: calculateOffsetX(s.SkeletonMoveDown8), OffsetY: calculateOffsetY(s.SkeletonMoveDown8)},
		{Graphic: sprite.SkeletonMoveDown9, Img: s.SkeletonMoveDown9, OffsetX: calculateOffsetX(s.SkeletonMoveDown9), OffsetY: calculateOffsetY(s.SkeletonMoveDown9)},
		// Skeleton Move Right Sprites
		{Graphic: sprite.SkeletonMoveRight1, Img: s.SkeletonMoveRight1, OffsetX: calculateOffsetX(s.SkeletonMoveRight1), OffsetY: calculateOffsetY(s.SkeletonMoveRight1)},
		{Graphic: sprite.SkeletonMoveRight2, Img: s.SkeletonMoveRight2, OffsetX: calculateOffsetX(s.SkeletonMoveRight2), OffsetY: calculateOffsetY(s.SkeletonMoveRight2)},
		{Graphic: sprite.SkeletonMoveRight3, Img: s.SkeletonMoveRight3, OffsetX: calculateOffsetX(s.SkeletonMoveRight3), OffsetY: calculateOffsetY(s.SkeletonMoveRight3)},
		{Graphic: sprite.SkeletonMoveRight4, Img: s.SkeletonMoveRight4, OffsetX: calculateOffsetX(s.SkeletonMoveRight4), OffsetY: calculateOffsetY(s.SkeletonMoveRight4)},
		{Graphic: sprite.SkeletonMoveRight5, Img: s.SkeletonMoveRight5, OffsetX: calculateOffsetX(s.SkeletonMoveRight5), OffsetY: calculateOffsetY(s.SkeletonMoveRight5)},
		{Graphic: sprite.SkeletonMoveRight6, Img: s.SkeletonMoveRight6, OffsetX: calculateOffsetX(s.SkeletonMoveRight6), OffsetY: calculateOffsetY(s.SkeletonMoveRight6)},
		{Graphic: sprite.SkeletonMoveRight7, Img: s.SkeletonMoveRight7, OffsetX: calculateOffsetX(s.SkeletonMoveRight7), OffsetY: calculateOffsetY(s.SkeletonMoveRight7)},
		{Graphic: sprite.SkeletonMoveRight8, Img: s.SkeletonMoveRight8, OffsetX: calculateOffsetX(s.SkeletonMoveRight8), OffsetY: calculateOffsetY(s.SkeletonMoveRight8)},
		{Graphic: sprite.SkeletonMoveRight9, Img: s.SkeletonMoveRight9, OffsetX: calculateOffsetX(s.SkeletonMoveRight9), OffsetY: calculateOffsetY(s.SkeletonMoveRight9)},
	}
}
