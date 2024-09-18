package scene

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type mockSystem struct {
	updateFunc func() error
	drawFunc   func(screen *ebiten.Image) error
}

func (m *mockSystem) Update() error {
	if m.updateFunc != nil {
		return m.updateFunc()
	}
	return nil // Default behavior if no function is provided
}

func (m *mockSystem) Draw(screen *ebiten.Image) error {
	if m.drawFunc != nil {
		return m.drawFunc(screen)
	}
	return nil // Default behavior if no function is provided
}

func TestScene_Draw(t *testing.T) {
	// Create a dummy screen
	screen := ebiten.NewImage(640, 480)

	tests := []struct {
		name          string
		systems       []systems.System
		expectedError error
	}{
		{
			name: "All systems draw successfully",
			systems: []systems.System{
				&mockSystem{
					drawFunc: func(s *ebiten.Image) error {
						if s != screen {
							t.Error("First system received incorrect screen parameter")
						}
						return nil
					},
				},
				&mockSystem{
					drawFunc: func(s *ebiten.Image) error {
						if s != screen {
							t.Error("Second system received incorrect screen parameter")
						}
						return nil
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "First system returns an error during draw",
			systems: []systems.System{
				&mockSystem{
					drawFunc: func(s *ebiten.Image) error {
						return errors.New("draw error")
					},
				},
				&mockSystem{
					drawFunc: func(s *ebiten.Image) error {
						t.Error("Second system's Draw() should not be called")
						return nil
					},
				},
			},
			expectedError: errors.New("draw error"),
		},
		{
			name:          "No systems to draw",
			systems:       []systems.System{},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a scene with the provided systems
			s := &Scene{
				systems: tt.systems,
			}

			// Call Draw
			err := s.Draw(screen)

			// Check expected error
			if tt.expectedError != nil {
				if err == nil {
					t.Fatalf("Expected error '%v', got nil", tt.expectedError)
				}
				if err.Error() != tt.expectedError.Error() {
					t.Fatalf("Expected error '%v', got '%v'", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestScene_Update(t *testing.T) {
	tests := []struct {
		name          string
		systems       []systems.System
		expectedError error
	}{
		{
			name: "All systems update successfully",
			systems: []systems.System{
				&mockSystem{
					updateFunc: func() error {
						return nil
					},
				},
				&mockSystem{
					updateFunc: func() error {
						return nil
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "First system returns an error",
			systems: []systems.System{
				&mockSystem{
					updateFunc: func() error {
						return errors.New("update error")
					},
				},
				&mockSystem{
					updateFunc: func() error {
						t.Error("Second system's Update() should not be called")
						return nil
					},
				},
			},
			expectedError: errors.New("update error"),
		},
		{
			name:          "No systems to update",
			systems:       []systems.System{},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a scene with the provided systems
			s := &Scene{
				systems: tt.systems,
			}

			// Call Update
			err := s.Update()

			// Check expected error
			if tt.expectedError != nil {
				if err == nil {
					t.Fatalf("Expected error '%v', got nil", tt.expectedError)
				}
				if err.Error() != tt.expectedError.Error() {
					t.Fatalf("Expected error '%v', got '%v'", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Run("Create new scene with systems", func(t *testing.T) {
		// Create mock systems
		system1 := &mockSystem{}
		system2 := &mockSystem{}

		// Create a scene using the New function
		s := New([]systems.System{system1, system2})

		// Assertions
		if s == nil {
			t.Fatal("Expected Scene to be non-nil")
		}
		if len(s.systems) != 2 {
			t.Fatalf("Expected 2 systems, got %d", len(s.systems))
		}
		if s.systems[0] != system1 {
			t.Error("First system in Scene is not system1")
		}
		if s.systems[1] != system2 {
			t.Error("Second system in Scene is not system2")
		}
	})

	t.Run("Create new scene with no systems", func(t *testing.T) {
		s := New([]systems.System{})

		if s == nil {
			t.Fatal("Expected Scene to be non-nil")
		}
		if len(s.systems) != 0 {
			t.Fatalf("Expected 0 systems, got %d", len(s.systems))
		}
	})
}
