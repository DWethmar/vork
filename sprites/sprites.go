package sprites

import (
	"github.com/dwethmar/vork/component/sprite"
	"github.com/dwethmar/vork/spritesheet"
	"github.com/dwethmar/vork/systems/render"
	"github.com/hajimehoshi/ebiten/v2"
)

// centerXOffset calculates the horizontal offset to center the image along the X-axis.
func centerXOffset(img *ebiten.Image) int {
	return -(img.Bounds().Dx() / 2)
}

// bottomYOffset calculates the vertical offset to place the image at the bottom of the screen.
func bottomYOffset(img *ebiten.Image) int {
	return -(img.Bounds().Dy())
}

// bottomCenteredAlignedSprite creates a sprite that is centered along the X-axis and aligned to the bottom.
func bottomCenteredAlignedSprite(graphic sprite.Graphic, img *ebiten.Image) render.Sprite {
	return render.Sprite{
		Graphic: graphic,
		Img:     img,
		OffsetX: centerXOffset(img),
		OffsetY: bottomYOffset(img),
	}
}

// Sprites returns all the sprites used in the game.
func Sprites(s *spritesheet.Spritesheet) []render.Sprite {
	return []render.Sprite{
		// Skeleton Death Sprites
		bottomCenteredAlignedSprite(sprite.SkeletonDeath1, s.SkeletonDeath1),
		bottomCenteredAlignedSprite(sprite.SkeletonDeath2, s.SkeletonDeath2),
		bottomCenteredAlignedSprite(sprite.SkeletonDeath3, s.SkeletonDeath3),
		bottomCenteredAlignedSprite(sprite.SkeletonDeath4, s.SkeletonDeath4),
		bottomCenteredAlignedSprite(sprite.SkeletonDeath5, s.SkeletonDeath5),
		bottomCenteredAlignedSprite(sprite.SkeletonDeath6, s.SkeletonDeath6),
		// Skeleton Move Up Sprites
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp1, s.SkeletonMoveUp1),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp2, s.SkeletonMoveUp2),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp3, s.SkeletonMoveUp3),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp4, s.SkeletonMoveUp4),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp5, s.SkeletonMoveUp5),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp6, s.SkeletonMoveUp6),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp7, s.SkeletonMoveUp7),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp8, s.SkeletonMoveUp8),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveUp9, s.SkeletonMoveUp9),
		// Skeleton Move Left Sprites
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft1, s.SkeletonMoveLeft1),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft2, s.SkeletonMoveLeft2),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft3, s.SkeletonMoveLeft3),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft4, s.SkeletonMoveLeft4),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft5, s.SkeletonMoveLeft5),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft6, s.SkeletonMoveLeft6),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft7, s.SkeletonMoveLeft7),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft8, s.SkeletonMoveLeft8),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveLeft9, s.SkeletonMoveLeft9),
		// Skeleton Move Down Sprites
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown1, s.SkeletonMoveDown1),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown2, s.SkeletonMoveDown2),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown3, s.SkeletonMoveDown3),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown4, s.SkeletonMoveDown4),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown5, s.SkeletonMoveDown5),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown6, s.SkeletonMoveDown6),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown7, s.SkeletonMoveDown7),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown8, s.SkeletonMoveDown8),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveDown9, s.SkeletonMoveDown9),
		// Skeleton Move Right Sprites
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight1, s.SkeletonMoveRight1),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight2, s.SkeletonMoveRight2),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight3, s.SkeletonMoveRight3),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight4, s.SkeletonMoveRight4),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight5, s.SkeletonMoveRight5),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight6, s.SkeletonMoveRight6),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight7, s.SkeletonMoveRight7),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight8, s.SkeletonMoveRight8),
		bottomCenteredAlignedSprite(sprite.SkeletonMoveRight9, s.SkeletonMoveRight9),
	}
}
