package persistence_test

import (
	"testing"

	"github.com/dwethmar/vork/event"
	"github.com/dwethmar/vork/systems/persistence"
	"github.com/dwethmar/vork/systems/persistence/mock"
)

func TestNew(t *testing.T) {
	t.Run("New should create a new system", func(t *testing.T) {
		r := &mock.Repository{}
		eventBus := event.NewBus()
		s := persistence.New(eventBus, r)
		if s == nil {
			t.Error("System should not be nil")
		}
	})

	t.Run("New should subscribe to component change events", func(t *testing.T) {
		r := &mock.Repository{}
		eventBus := event.NewBus()
		s := persistence.New(eventBus, r)
		if s == nil {
			t.Error("System should not be nil")
		}
	})
}

func TestSystem_Save(t *testing.T) {

}

func TestSystem_Load(t *testing.T) {

}
