package game

import (
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

func calcOffsetX(img *ebiten.Image) int {
	return -(img.Bounds().Dx() / 2)
}

func calcOffsetY(img *ebiten.Image) int {
	return -(img.Bounds().Dy())
}

func Sprites(s *spritesheet.Spritesheet) []render.Sprite {
	return []render.Sprite{
		// Skeleton Death Sprites
		{Graphic: sprite.SkeletonDeath1, Img: s.SkeletonDeath1, OffsetX: calcOffsetX(s.SkeletonDeath1), OffsetY: calcOffsetY(s.SkeletonDeath1)},
		{Graphic: sprite.SkeletonDeath2, Img: s.SkeletonDeath2, OffsetX: calcOffsetX(s.SkeletonDeath2), OffsetY: calcOffsetY(s.SkeletonDeath2)},
		{Graphic: sprite.SkeletonDeath3, Img: s.SkeletonDeath3, OffsetX: calcOffsetX(s.SkeletonDeath3), OffsetY: calcOffsetY(s.SkeletonDeath3)},
		{Graphic: sprite.SkeletonDeath4, Img: s.SkeletonDeath4, OffsetX: calcOffsetX(s.SkeletonDeath4), OffsetY: calcOffsetY(s.SkeletonDeath4)},
		{Graphic: sprite.SkeletonDeath5, Img: s.SkeletonDeath5, OffsetX: calcOffsetX(s.SkeletonDeath5), OffsetY: calcOffsetY(s.SkeletonDeath5)},
		{Graphic: sprite.SkeletonDeath6, Img: s.SkeletonDeath6, OffsetX: calcOffsetX(s.SkeletonDeath6), OffsetY: calcOffsetY(s.SkeletonDeath6)},
		// Skeleton Move Up Sprites
		{Graphic: sprite.SkeletonMoveUp1, Img: s.SkeletonMoveUp1, OffsetX: calcOffsetX(s.SkeletonMoveUp1), OffsetY: calcOffsetY(s.SkeletonMoveUp1)},
		{Graphic: sprite.SkeletonMoveUp2, Img: s.SkeletonMoveUp2, OffsetX: calcOffsetX(s.SkeletonMoveUp2), OffsetY: calcOffsetY(s.SkeletonMoveUp2)},
		{Graphic: sprite.SkeletonMoveUp3, Img: s.SkeletonMoveUp3, OffsetX: calcOffsetX(s.SkeletonMoveUp3), OffsetY: calcOffsetY(s.SkeletonMoveUp3)},
		{Graphic: sprite.SkeletonMoveUp4, Img: s.SkeletonMoveUp4, OffsetX: calcOffsetX(s.SkeletonMoveUp4), OffsetY: calcOffsetY(s.SkeletonMoveUp4)},
		{Graphic: sprite.SkeletonMoveUp5, Img: s.SkeletonMoveUp5, OffsetX: calcOffsetX(s.SkeletonMoveUp5), OffsetY: calcOffsetY(s.SkeletonMoveUp5)},
		{Graphic: sprite.SkeletonMoveUp6, Img: s.SkeletonMoveUp6, OffsetX: calcOffsetX(s.SkeletonMoveUp6), OffsetY: calcOffsetY(s.SkeletonMoveUp6)},
		{Graphic: sprite.SkeletonMoveUp7, Img: s.SkeletonMoveUp7, OffsetX: calcOffsetX(s.SkeletonMoveUp7), OffsetY: calcOffsetY(s.SkeletonMoveUp7)},
		{Graphic: sprite.SkeletonMoveUp8, Img: s.SkeletonMoveUp8, OffsetX: calcOffsetX(s.SkeletonMoveUp8), OffsetY: calcOffsetY(s.SkeletonMoveUp8)},
		{Graphic: sprite.SkeletonMoveUp9, Img: s.SkeletonMoveUp9, OffsetX: calcOffsetX(s.SkeletonMoveUp9), OffsetY: calcOffsetY(s.SkeletonMoveUp9)},
		// Skeleton Move Left Sprites
		{Graphic: sprite.SkeletonMoveLeft1, Img: s.SkeletonMoveLeft1, OffsetX: calcOffsetX(s.SkeletonMoveLeft1), OffsetY: calcOffsetY(s.SkeletonMoveLeft1)},
		{Graphic: sprite.SkeletonMoveLeft2, Img: s.SkeletonMoveLeft2, OffsetX: calcOffsetX(s.SkeletonMoveLeft2), OffsetY: calcOffsetY(s.SkeletonMoveLeft2)},
		{Graphic: sprite.SkeletonMoveLeft3, Img: s.SkeletonMoveLeft3, OffsetX: calcOffsetX(s.SkeletonMoveLeft3), OffsetY: calcOffsetY(s.SkeletonMoveLeft3)},
		{Graphic: sprite.SkeletonMoveLeft4, Img: s.SkeletonMoveLeft4, OffsetX: calcOffsetX(s.SkeletonMoveLeft4), OffsetY: calcOffsetY(s.SkeletonMoveLeft4)},
		{Graphic: sprite.SkeletonMoveLeft5, Img: s.SkeletonMoveLeft5, OffsetX: calcOffsetX(s.SkeletonMoveLeft5), OffsetY: calcOffsetY(s.SkeletonMoveLeft5)},
		{Graphic: sprite.SkeletonMoveLeft6, Img: s.SkeletonMoveLeft6, OffsetX: calcOffsetX(s.SkeletonMoveLeft6), OffsetY: calcOffsetY(s.SkeletonMoveLeft6)},
		{Graphic: sprite.SkeletonMoveLeft7, Img: s.SkeletonMoveLeft7, OffsetX: calcOffsetX(s.SkeletonMoveLeft7), OffsetY: calcOffsetY(s.SkeletonMoveLeft7)},
		{Graphic: sprite.SkeletonMoveLeft8, Img: s.SkeletonMoveLeft8, OffsetX: calcOffsetX(s.SkeletonMoveLeft8), OffsetY: calcOffsetY(s.SkeletonMoveLeft8)},
		{Graphic: sprite.SkeletonMoveLeft9, Img: s.SkeletonMoveLeft9, OffsetX: calcOffsetX(s.SkeletonMoveLeft9), OffsetY: calcOffsetY(s.SkeletonMoveLeft9)},
		// Skeleton Move Down Sprites
		{Graphic: sprite.SkeletonMoveDown1, Img: s.SkeletonMoveDown1, OffsetX: calcOffsetX(s.SkeletonMoveDown1), OffsetY: calcOffsetY(s.SkeletonMoveDown1)},
		{Graphic: sprite.SkeletonMoveDown2, Img: s.SkeletonMoveDown2, OffsetX: calcOffsetX(s.SkeletonMoveDown2), OffsetY: calcOffsetY(s.SkeletonMoveDown2)},
		{Graphic: sprite.SkeletonMoveDown3, Img: s.SkeletonMoveDown3, OffsetX: calcOffsetX(s.SkeletonMoveDown3), OffsetY: calcOffsetY(s.SkeletonMoveDown3)},
		{Graphic: sprite.SkeletonMoveDown4, Img: s.SkeletonMoveDown4, OffsetX: calcOffsetX(s.SkeletonMoveDown4), OffsetY: calcOffsetY(s.SkeletonMoveDown4)},
		{Graphic: sprite.SkeletonMoveDown5, Img: s.SkeletonMoveDown5, OffsetX: calcOffsetX(s.SkeletonMoveDown5), OffsetY: calcOffsetY(s.SkeletonMoveDown5)},
		{Graphic: sprite.SkeletonMoveDown6, Img: s.SkeletonMoveDown6, OffsetX: calcOffsetX(s.SkeletonMoveDown6), OffsetY: calcOffsetY(s.SkeletonMoveDown6)},
		{Graphic: sprite.SkeletonMoveDown7, Img: s.SkeletonMoveDown7, OffsetX: calcOffsetX(s.SkeletonMoveDown7), OffsetY: calcOffsetY(s.SkeletonMoveDown7)},
		{Graphic: sprite.SkeletonMoveDown8, Img: s.SkeletonMoveDown8, OffsetX: calcOffsetX(s.SkeletonMoveDown8), OffsetY: calcOffsetY(s.SkeletonMoveDown8)},
		{Graphic: sprite.SkeletonMoveDown9, Img: s.SkeletonMoveDown9, OffsetX: calcOffsetX(s.SkeletonMoveDown9), OffsetY: calcOffsetY(s.SkeletonMoveDown9)},
		// Skeleton Move Right Sprites
		{Graphic: sprite.SkeletonMoveRight1, Img: s.SkeletonMoveRight1, OffsetX: calcOffsetX(s.SkeletonMoveRight1), OffsetY: calcOffsetY(s.SkeletonMoveRight1)},
		{Graphic: sprite.SkeletonMoveRight2, Img: s.SkeletonMoveRight2, OffsetX: calcOffsetX(s.SkeletonMoveRight2), OffsetY: calcOffsetY(s.SkeletonMoveRight2)},
		{Graphic: sprite.SkeletonMoveRight3, Img: s.SkeletonMoveRight3, OffsetX: calcOffsetX(s.SkeletonMoveRight3), OffsetY: calcOffsetY(s.SkeletonMoveRight3)},
		{Graphic: sprite.SkeletonMoveRight4, Img: s.SkeletonMoveRight4, OffsetX: calcOffsetX(s.SkeletonMoveRight4), OffsetY: calcOffsetY(s.SkeletonMoveRight4)},
		{Graphic: sprite.SkeletonMoveRight5, Img: s.SkeletonMoveRight5, OffsetX: calcOffsetX(s.SkeletonMoveRight5), OffsetY: calcOffsetY(s.SkeletonMoveRight5)},
		{Graphic: sprite.SkeletonMoveRight6, Img: s.SkeletonMoveRight6, OffsetX: calcOffsetX(s.SkeletonMoveRight6), OffsetY: calcOffsetY(s.SkeletonMoveRight6)},
		{Graphic: sprite.SkeletonMoveRight7, Img: s.SkeletonMoveRight7, OffsetX: calcOffsetX(s.SkeletonMoveRight7), OffsetY: calcOffsetY(s.SkeletonMoveRight7)},
		{Graphic: sprite.SkeletonMoveRight8, Img: s.SkeletonMoveRight8, OffsetX: calcOffsetX(s.SkeletonMoveRight8), OffsetY: calcOffsetY(s.SkeletonMoveRight8)},
		{Graphic: sprite.SkeletonMoveRight9, Img: s.SkeletonMoveRight9, OffsetX: calcOffsetX(s.SkeletonMoveRight9), OffsetY: calcOffsetY(s.SkeletonMoveRight9)},
	}
}
