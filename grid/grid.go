package grid

import (
	"errors"
	"iter"

	"github.com/dwethmar/vork/point"
)

// Item is a point with a value.
type Item struct {
	point.Point
	V int
}

func NewItem(x, y, v int) Item {
	return Item{point.New(x, y), v}
}

var (
	ErrXOutOfBounds = errors.New("x out of bounds")
	ErrYOutOfBounds = errors.New("y out of bounds")
)

type Grid [][]int

func New(width, height, v int) Grid {
	g := make([][]int, height)
	for i := range g {
		g[i] = make([]int, width)
		for j := range g[i] {
			g[i][j] = v
		}
	}
	return g
}

// Get returns the value at the given x, y position or the default value if the position is out of bounds.
func (g Grid) Get(x, y, d int) int {
	if x < 0 || x >= len(g[0]) {
		return d
	}
	if y < 0 || y >= len(g) {
		return d
	}
	return g[y][x]
}

func (g Grid) Set(x, y, value int) error {
	if x < 0 || x >= len(g[0]) {
		return ErrXOutOfBounds
	}
	if y < 0 || y >= len(g) {
		return ErrYOutOfBounds
	}
	g[y][x] = value
	return nil
}

func (g Grid) Iterator() iter.Seq[Item] {
	return func(yield func(Item) bool) {
		for y, row := range g {
			for x := range row {
				if !yield(NewItem(x, y, g[y][x])) {
					return
				}
			}
		}
	}
}
