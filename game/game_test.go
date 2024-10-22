package game_test

import (
	"testing"

	"github.com/dwethmar/vork/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneMock struct {
	nameFunc   func() string
	drawFunc   func(screen *ebiten.Image) error
	updateFunc func() error
	closeFunc  func() error
}

func (s *SceneMock) Name() string                    { return s.nameFunc() }
func (s *SceneMock) Draw(screen *ebiten.Image) error { return s.drawFunc(screen) }
func (s *SceneMock) Update() error                   { return s.updateFunc() }
func (s *SceneMock) Close() error                    { return s.closeFunc() }

func TestNew(t *testing.T) {
	got, err := game.New()
	if err != nil {
		t.Errorf("New() = %v, %v", got, err)
	}
	if got == nil {
		t.Errorf("New() = %v, want not nil", got)
	}
}

func TestGame_SwitchScene(t *testing.T) {
	t.Run("scene not found", func(t *testing.T) {
		g, _ := game.New()
		err := g.SwitchScene("not found")
		if err == nil {
			t.Errorf("Game.SwitchScene() error = %v, want not nil", err)
		}
	})
}

func TestGame_AddScene(t *testing.T) {
	t.Run("scene already exists", func(t *testing.T) {
		g, _ := game.New()
		scene := &SceneMock{
			nameFunc: func() string {
				return "test"
			},
		}
		if err := g.AddScene(scene); err != nil {
			t.Errorf("Game.AddScene() error = %v, want nil", err)
		}
		if err := g.AddScene(scene); err == nil {
			t.Errorf("Game.AddScene() error = %v, want not nil", err)
		}
	})
}

func TestGame_Draw(t *testing.T) {
	t.Run("scene is nil", func(t *testing.T) {
		g, _ := game.New()
		err := g.Draw(nil)
		if err != nil {
			t.Errorf("Game.Draw() error = %v, want nil", err)
		}
	})
	t.Run("scene is not nil", func(t *testing.T) {
		g, _ := game.New()
		scene := &SceneMock{
			nameFunc: func() string {
				return "test"
			},
			drawFunc: func(_ *ebiten.Image) error {
				return nil
			},
		}
		if err := g.AddScene(scene); err != nil {
			t.Errorf("Game.AddScene() error = %v, want nil", err)
		}
		if err := g.SwitchScene(scene.Name()); err != nil {
			t.Errorf("Game.SwitchScene() error = %v, want nil", err)
		}
		if err := g.Draw(nil); err != nil {
			t.Errorf("Game.Draw() error = %v, want nil", err)
		}
	})
}

func TestGame_Update(t *testing.T) {
	t.Run("scene is nil", func(t *testing.T) {
		g, _ := game.New()
		err := g.Update()
		if err != nil {
			t.Errorf("Game.Update() error = %v, want nil", err)
		}
	})
	t.Run("scene is not nil", func(t *testing.T) {
		g, _ := game.New()
		scene := &SceneMock{
			nameFunc: func() string {
				return "test"
			},
			updateFunc: func() error {
				return nil
			},
		}
		var err error
		if err = g.AddScene(scene); err != nil {
			t.Errorf("Game.AddScene() error = %v, want nil", err)
		}
		if err = g.SwitchScene(scene.Name()); err != nil {
			t.Errorf("Game.SwitchScene() error = %v, want nil", err)
		}
		if err = g.Update(); err != nil {
			t.Errorf("Game.Update() error = %v, want nil", err)
		}
	})
}
