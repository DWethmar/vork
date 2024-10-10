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

// newRenderSprite creates a new render sprite.
func newRenderSprite(graphic sprite.Graphic, img *ebiten.Image) render.Sprite {
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
		newRenderSprite(sprite.SkeletonDeath1, s.SkeletonDeath1),
		newRenderSprite(sprite.SkeletonDeath2, s.SkeletonDeath2),
		newRenderSprite(sprite.SkeletonDeath3, s.SkeletonDeath3),
		newRenderSprite(sprite.SkeletonDeath4, s.SkeletonDeath4),
		newRenderSprite(sprite.SkeletonDeath5, s.SkeletonDeath5),
		newRenderSprite(sprite.SkeletonDeath6, s.SkeletonDeath6),
		// Skeleton Move Up Sprites
		newRenderSprite(sprite.SkeletonMoveUp1, s.SkeletonMoveUp1),
		newRenderSprite(sprite.SkeletonMoveUp2, s.SkeletonMoveUp2),
		newRenderSprite(sprite.SkeletonMoveUp3, s.SkeletonMoveUp3),
		newRenderSprite(sprite.SkeletonMoveUp4, s.SkeletonMoveUp4),
		newRenderSprite(sprite.SkeletonMoveUp5, s.SkeletonMoveUp5),
		newRenderSprite(sprite.SkeletonMoveUp6, s.SkeletonMoveUp6),
		newRenderSprite(sprite.SkeletonMoveUp7, s.SkeletonMoveUp7),
		newRenderSprite(sprite.SkeletonMoveUp8, s.SkeletonMoveUp8),
		newRenderSprite(sprite.SkeletonMoveUp9, s.SkeletonMoveUp9),
		// Skeleton Move Left Sprites
		newRenderSprite(sprite.SkeletonMoveLeft1, s.SkeletonMoveLeft1),
		newRenderSprite(sprite.SkeletonMoveLeft2, s.SkeletonMoveLeft2),
		newRenderSprite(sprite.SkeletonMoveLeft3, s.SkeletonMoveLeft3),
		newRenderSprite(sprite.SkeletonMoveLeft4, s.SkeletonMoveLeft4),
		newRenderSprite(sprite.SkeletonMoveLeft5, s.SkeletonMoveLeft5),
		newRenderSprite(sprite.SkeletonMoveLeft6, s.SkeletonMoveLeft6),
		newRenderSprite(sprite.SkeletonMoveLeft7, s.SkeletonMoveLeft7),
		newRenderSprite(sprite.SkeletonMoveLeft8, s.SkeletonMoveLeft8),
		newRenderSprite(sprite.SkeletonMoveLeft9, s.SkeletonMoveLeft9),
		// Skeleton Move Down Sprites
		newRenderSprite(sprite.SkeletonMoveDown1, s.SkeletonMoveDown1),
		newRenderSprite(sprite.SkeletonMoveDown2, s.SkeletonMoveDown2),
		newRenderSprite(sprite.SkeletonMoveDown3, s.SkeletonMoveDown3),
		newRenderSprite(sprite.SkeletonMoveDown4, s.SkeletonMoveDown4),
		newRenderSprite(sprite.SkeletonMoveDown5, s.SkeletonMoveDown5),
		newRenderSprite(sprite.SkeletonMoveDown6, s.SkeletonMoveDown6),
		newRenderSprite(sprite.SkeletonMoveDown7, s.SkeletonMoveDown7),
		newRenderSprite(sprite.SkeletonMoveDown8, s.SkeletonMoveDown8),
		newRenderSprite(sprite.SkeletonMoveDown9, s.SkeletonMoveDown9),
		// Skeleton Move Right Sprites
		newRenderSprite(sprite.SkeletonMoveRight1, s.SkeletonMoveRight1),
		newRenderSprite(sprite.SkeletonMoveRight2, s.SkeletonMoveRight2),
		newRenderSprite(sprite.SkeletonMoveRight3, s.SkeletonMoveRight3),
		newRenderSprite(sprite.SkeletonMoveRight4, s.SkeletonMoveRight4),
		newRenderSprite(sprite.SkeletonMoveRight5, s.SkeletonMoveRight5),
		newRenderSprite(sprite.SkeletonMoveRight6, s.SkeletonMoveRight6),
		newRenderSprite(sprite.SkeletonMoveRight7, s.SkeletonMoveRight7),
		newRenderSprite(sprite.SkeletonMoveRight8, s.SkeletonMoveRight8),
		newRenderSprite(sprite.SkeletonMoveRight9, s.SkeletonMoveRight9),
	}
}
