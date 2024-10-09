package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	Draw(screen *ebiten.Image) error
	Update() error
	Close() error
}
