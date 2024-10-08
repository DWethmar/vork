// package ecsys contains the Entity-Component-System architecture.
package ecsys_test

import (
	"testing"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/event"
)

func TestECS_CreateEntity(t *testing.T) {
	t.Run("CreateEntity", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus())
		entity, err := ecs.CreateEntity(0, 0)
		if err != nil {
			t.Errorf("CreateEntity() error = %v", err)
		}
		if entity == 0 {
			t.Errorf("CreateEntity() entity = %v", entity)
		}
	})
}

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		ecs := ecsys.New(event.NewBus())
		if ecs == nil {
			t.Errorf("New() ecs = %v", ecs)
		}
	})
}
