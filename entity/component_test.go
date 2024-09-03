package entity_test

import (
	"testing"

	"github.com/dwethmar/vork/entity"
	"github.com/google/go-cmp/cmp"
)

type MockComponent struct {
	I uint32
	T string
	E entity.Entity
}

func (m *MockComponent) ID() uint32            { return m.I }
func (m *MockComponent) SetID(i uint32)        { m.I = i }
func (m *MockComponent) Type() string          { return m.T }
func (m *MockComponent) Entity() entity.Entity { return m.E }

func TestNewComponentStore(t *testing.T) {
	t.Run("new should return a new ComponentStore", func(t *testing.T) {
		cm := entity.NewComponentStore(0)
		if cm == nil {
			t.Error("NewComponentStore() should not return nil")
		}
	})
}

func TestComponentStore_Add(t *testing.T) {
	t.Run("add should add a component to the manager", func(t *testing.T) {
		cm := entity.NewComponentStore(0)
		c := &MockComponent{
			T: "test",
			E: 1,
		}
		cm.Add(c)
		e := cm.Get(c.Entity(), c.Type())

		if diff := cmp.Diff(c, e); diff != "" {
			t.Errorf("Add() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestComponentStore_Remove(t *testing.T) {
	t.Run("remove should remove a component from the manager", func(t *testing.T) {
		cm := entity.NewComponentStore(0)
		c := &MockComponent{
			T: "test",
			E: 1,
		}
		cm.Add(c)
		cm.Remove(c.Entity(), c.Type())
		e := cm.Get(c.Entity(), c.Type())
		if e != nil {
			t.Error("Remove() should remove the component")
		}
	})
}

func TestComponentStore_Get(t *testing.T) {
	t.Run("get should return a component from the manager", func(t *testing.T) {
		cm := entity.NewComponentStore(0)
		c := &MockComponent{
			T: "test",
			E: 1,
		}
		cm.Add(c)
		e := cm.Get(c.Entity(), c.Type())
		if diff := cmp.Diff(c, e); diff != "" {
			t.Errorf("Get() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestComponentStore_List(t *testing.T) {
	t.Run("list should return all components of a certain type, sorted by component ID", func(t *testing.T) {
		cm := entity.NewComponentStore(0)
		for _, c := range []entity.Component{
			&MockComponent{
				I: 2, // id 2 is inserted before id 1
				T: "test",
				E: 2,
			},
			&MockComponent{
				I: 1,
				T: "test",
				E: 1,
			},
			&MockComponent{
				I: 3,
				T: "test",
				E: 3,
			},
			&MockComponent{
				I: 4,
				T: "test-2", // different type
				E: 1,
			},
		} {
			cm.Add(c)
		}
		l := cm.List("test")
		if len(l) != 3 {
			t.Errorf("List() should return 1 component, got %d", len(l))
		}
		expected := []entity.Component{ // sorted by component ID
			&MockComponent{
				I: 1,
				T: "test",
				E: 1,
			},
			&MockComponent{
				I: 2,
				T: "test",
				E: 2,
			},
			&MockComponent{
				I: 3,
				T: "test",
				E: 3,
			},
		}
		if diff := cmp.Diff(expected, l); diff != "" {
			t.Errorf("List() mismatch (-want +got):\n%s", diff)
		}
	})
}
