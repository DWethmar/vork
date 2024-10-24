package grid_test

import (
	"errors"
	"testing"

	"github.com/dwethmar/vork/grid"
	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	g := grid.New(10, 10, 0)
	if len(g) != 10 {
		t.Errorf("expected 10, got %d", len(g))
	}
}

func TestGrid_Get(t *testing.T) {
	g := grid.New(10, 10, 0)
	if g.Get(0, 0, 0) != 0 {
		t.Errorf("expected 0, got %d", g.Get(0, 0, 0))
	}
	if g.Get(10, 10, 0) != 0 {
		t.Errorf("expected 0, got %d", g.Get(10, 10, 0))
	}
	if g.Get(-1, -1, 0) != 0 {
		t.Errorf("expected 0, got %d", g.Get(-1, -1, 0))
	}
}

func TestGrid_Set(t *testing.T) {
	g := grid.New(10, 10, 0)
	if err := g.Set(0, 0, 1); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if err := g.Set(10, 0, 1); !errors.Is(err, grid.ErrXOutOfBounds) {
		t.Errorf("expected %v, got %v", grid.ErrXOutOfBounds, err)
	}
	if err := g.Set(0, 10, 1); !errors.Is(err, grid.ErrYOutOfBounds) {
		t.Errorf("expected %v, got %v", grid.ErrYOutOfBounds, err)
	}
}

func TestGrid_Iterator(t *testing.T) {
	g := grid.New(2, 2, 9)
	expect := []grid.Item{
		grid.NewItem(0, 0, 9),
		grid.NewItem(1, 0, 9),
		grid.NewItem(0, 1, 9),
		grid.NewItem(1, 1, 9),
	}
	got := []grid.Item{}
	for p := range g.Iterator() {
		got = append(got, p)
	}
	if diff := cmp.Diff(expect, got); diff != "" {
		t.Errorf("unexpected iterator result (-want +got):\n%s", diff)
	}
}
