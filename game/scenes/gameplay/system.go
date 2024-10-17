package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// System is the interface that wraps the basic methods of a gameplay system.
type System interface {
	Init() error
	Draw(screen *ebiten.Image) error
	Update() error
	Close() error
}
