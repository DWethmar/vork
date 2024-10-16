package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	Init() error
	Draw(screen *ebiten.Image) error
	Update() error
	Close() error
}
