package memory_test

import (
	"testing"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/scene/memory"
	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		m := memory.New()
		if m == nil {
			t.Error("Memory is nil")
		}
	})
}

func TestMemory_CreateEntity(t *testing.T) {
	t.Run("CreateEntity", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		if e == 0 {
			t.Error("Entity is 0")
		}
	})
}

func TestMemory_DeleteEntity(t *testing.T) {
	t.Run("deleting a non-existing entity should return an error", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		if m.DeleteEntity(e) == nil {
			t.Errorf("DeleteEntity should return an error")
		}
	})

	t.Run("deleting an existing entity should not return an error", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		m.AddComponent(&component.BaseComponent{
			E: e,
			T: "test",
		})
		if err := m.DeleteEntity(e); err != nil {
			t.Errorf("DeleteEntity should not return an error, got: %v", err)
		}
	})
}

func TestMemory_Component(t *testing.T) {
	t.Run("component should return a component", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		c := &component.BaseComponent{
			I: 1,
			E: e,
			T: "test",
		}
		m.AddComponent(c)
		_, ok := m.Component(e, c.Type())
		if !ok {
			t.Error("Component not found")
		}
	})
}

func TestMemory_Components(t *testing.T) {
	t.Run("Components should return a list of components", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		c1 := &component.BaseComponent{
			I: 1,
			E: e,
			T: "test",
		}
		m.AddComponent(c1)
		c2 := &component.BaseComponent{
			I: 2,
			E: e,
			T: "test",
		}
		m.AddComponent(c2)
		c3 := &component.BaseComponent{
			I: 3,
			E: e,
			T: "test-2", // different type
		}
		m.AddComponent(c3)
		r, err := m.Components(e)
		if err != nil {
			t.Error("Components failed")
		}
		if len(r) == 0 {
			t.Error("Components is empty")
		}
		if diff := cmp.Diff(r, []component.Component{c1, c2, c3}); diff != "" {
			t.Errorf("Components mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMemory_ComponentsByType(t *testing.T) {
	t.Run("ComponentsByType should return a list of components", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		c1 := &component.BaseComponent{
			I: 1,
			E: e,
			T: "test",
		}
		m.AddComponent(c1)
		c2 := &component.BaseComponent{
			I: 2,
			E: e,
			T: "test",
		}
		m.AddComponent(c2)
		c3 := &component.BaseComponent{
			I: 3,
			E: e,
			T: "test-2", // different type and should not be returned
		}
		m.AddComponent(c3)
		r := m.ComponentsByType(c1.Type())
		if len(r) == 0 {
			t.Error("ComponentsByType is empty")
		}
		if diff := cmp.Diff(r, []component.Component{c1, c2}); diff != "" {
			t.Errorf("ComponentsByType mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMemory_AddComponent(t *testing.T) {
	t.Run("AddComponent should add a component", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		c := &component.BaseComponent{
			I: 1,
			E: e,
			T: "test",
		}
		if m.AddComponent(c) == 0 {
			t.Error("ID is 0")
		}
		r, ok := m.Component(e, c.Type())
		if !ok {
			t.Error("Component not found")
		}
		expected := &component.BaseComponent{
			I: 1,
			E: e,
			T: "test",
		}
		if diff := cmp.Diff(r, expected); diff != "" {
			t.Errorf("AddComponent mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMemory_DeleteComponents(t *testing.T) {
	t.Run("DeleteComponents should delete a component", func(t *testing.T) {
		m := memory.New()
		e := m.CreateEntity()
		c := &component.BaseComponent{
			I: 1,
			E: e,
			T: "test",
		}
		m.AddComponent(c)
		if err := m.DeleteComponents(e, 1); err != nil {
			t.Errorf("DeleteComponents should not return an error, got: %v", err)
		}
		_, ok := m.Component(e, c.Type())
		if ok {
			t.Error("Component not deleted")
		}
	})
}
