package store_test

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/store"
	"github.com/dwethmar/vork/entity"
	"github.com/google/go-cmp/cmp"
)

type TestComponent struct {
	I   uint          // ID
	E   entity.Entity // Entity
	Tag string        // Tag used to identify the sprite
}

func (t *TestComponent) ID() uint              { return t.I }
func (t *TestComponent) SetID(i uint)          { t.I = i }
func (t *TestComponent) Type() component.Type  { return "test" }
func (t *TestComponent) Entity() entity.Entity { return t.E }

func TestNewMemStore(t *testing.T) {
	s := store.NewMemStore[*TestComponent](false)
	if s == nil {
		t.Error("NewMemStore() should not return nil")
	}
}

func TestMemStoreAdd(t *testing.T) {
	t.Run("should add component", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		id, err := s.Add(c)
		if err != nil {
			t.Error("Add() should not return an error")
		}

		if id != 1 {
			t.Errorf("Add() should return 1, got %d", id)
		}
	})

	t.Run("should fail to add component with same ID", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, err := s.Add(c)
		if err != nil {
			t.Error("Add() should not return an error")
		}

		_, err = s.Add(c)
		if err == nil {
			t.Error("Add() should return an error")
		}
	})

	t.Run("should fail if store only 1 component per entity", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](true)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, err := s.Add(c)
		if err != nil {
			t.Error("Add() should not return an error")
		}
		c = &TestComponent{
			I:   2,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, err = s.Add(c)
		if !errors.Is(err, store.ErrUniqueComponentViolation) {
			t.Errorf("Add() should return ErrUniquePerEntity, got %v", err)
		}
	})
}

func TestMemStoreGet(t *testing.T) {
	t.Run("should return component", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, err := s.Add(c)
		if err != nil {
			t.Error("Add() should not return an error")
		}

		got, err := s.Get(1)
		if err != nil {
			t.Error("Get() should not return an error")
		}

		expect := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}

		if diff := cmp.Diff(expect, got); diff != "" {
			t.Errorf("Get() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("should return error if component not found", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		_, err := s.Get(1)
		if !errors.Is(err, store.ErrComponentNotFound) {
			t.Errorf("Get() should return ErrComponentNotFound, got %v", err)
		}
	})
}

func TestMemStoreUpdate(t *testing.T) {
	t.Run("should update component", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, err := s.Add(c)
		if err != nil {
			t.Error("Add() should not return an error")
		}

		c.Tag = "updated"
		err = s.Update(c)
		if err != nil {
			t.Error("Update() should not return an error")
		}

		got, err := s.First(entity.Entity(1))
		if err != nil {
			t.Error("Get() should not return an error")
		}

		expect := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "updated",
		}

		if diff := cmp.Diff(expect, got); diff != "" {
			t.Errorf("Get() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("should return error if component not found", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		err := s.Update(c)
		if !errors.Is(err, store.ErrComponentNotFound) {
			t.Errorf("Update() should return ErrComponentNotFound, got %v", err)
		}
	})

	t.Run("should return error if unique constraint is violated", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](true)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, err := s.Add(c)
		if err != nil {
			t.Error("Add() should not return an error")
		}

		c = &TestComponent{
			I:   2,
			E:   entity.Entity(1),
			Tag: "test",
		}
		err = s.Update(c)
		if !errors.Is(err, store.ErrUniqueComponentViolation) {
			t.Errorf("Add() should return ErrUniquePerEntity, got %v", err)
		}
	})
}

func TestMemStoreDelete(t *testing.T) {
	t.Run("should delete component", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, err := s.Add(c)
		if err != nil {
			t.Error("Add() should not return an error")
		}

		c2 := &TestComponent{
			I:   2,
			E:   entity.Entity(1),
			Tag: "test",
		}
		if _, err = s.Add(c2); err != nil {
			t.Error("Add() should not return an error")
		}

		err = s.Delete(1)
		if err != nil {
			t.Error("Delete() should not return an error")
		}

		_, err = s.Get(1)
		if !errors.Is(err, store.ErrComponentNotFound) {
			t.Errorf("Get() should return ErrComponentNotFound, got %v", err)
		}

		_, err = s.Get(2)
		if err != nil {
			t.Error("Get() should not return an error")
		}
	})

	t.Run("should return error if component not found", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		err := s.Delete(1)
		if !errors.Is(err, store.ErrComponentNotFound) {
			t.Errorf("Delete() should return ErrComponentNotFound, got %v", err)
		}
	})
}

func TestMemStoreList(t *testing.T) {
	t.Run("should return all components", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c1 := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		c2 := &TestComponent{
			I:   2,
			E:   entity.Entity(2),
			Tag: "test",
		}
		_, _ = s.Add(c1)
		_, _ = s.Add(c2)

		got := s.List()
		expect := []*TestComponent{c1, c2}

		if diff := cmp.Diff(expect, got); diff != "" {
			t.Errorf("List() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMemStoreFirstByEntity(t *testing.T) {
	t.Run("should return first component by entity", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c1 := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		c2 := &TestComponent{
			I:   2,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, _ = s.Add(c1)
		_, _ = s.Add(c2)

		got, err := s.First(entity.Entity(1))
		if err != nil {
			t.Error("First() should not return an error")
		}

		expect := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}

		if diff := cmp.Diff(expect, got); diff != "" {
			t.Errorf("FirstByEntity() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("should return error if component not found", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		_, err := s.First(entity.Entity(1))
		if !errors.Is(err, store.ErrEntityNotFound) {
			t.Errorf("FirstByEntity() should return ErrComponentNotFound, got %v", err)
		}
	})
}

func TestMemStoreListByEntity(t *testing.T) {
	t.Run("should return all components by entity", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c1 := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		c2 := &TestComponent{
			I:   2,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, _ = s.Add(c1)
		_, _ = s.Add(c2)

		got := s.ListByEntity(entity.Entity(1))
		expect := []*TestComponent{c1, c2}

		if diff := cmp.Diff(expect, got); diff != "" {
			t.Errorf("ListByEntity() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("should return nil if no components found", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		got := s.ListByEntity(entity.Entity(1))
		if got != nil {
			t.Errorf("ListByEntity() should return nil, got %v", got)
		}
	})
}

func TestMemStoreDeleteByEntity(t *testing.T) {
	t.Run("should delete all components by entity", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		c1 := &TestComponent{
			I:   1,
			E:   entity.Entity(1),
			Tag: "test",
		}
		c2 := &TestComponent{
			I:   2,
			E:   entity.Entity(1),
			Tag: "test",
		}
		_, _ = s.Add(c1)
		_, _ = s.Add(c2)

		err := s.DeleteByEntity(entity.Entity(1))
		if err != nil {
			t.Error("DeleteByEntity() should not return an error")
		}

		got := s.ListByEntity(entity.Entity(1))
		if got != nil {
			t.Errorf("ListByEntity() should return nil, got %v", got)
		}
	})

	t.Run("should return error if entity not found", func(t *testing.T) {
		s := store.NewMemStore[*TestComponent](false)
		err := s.DeleteByEntity(entity.Entity(1))
		if !errors.Is(err, store.ErrEntityNotFound) {
			t.Errorf("DeleteByEntity() should return ErrEntityNotFound, got %v", err)
		}
	})
}
