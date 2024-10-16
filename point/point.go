package point

// Point is a simple 2D point with integer coordinates.
type Point struct {
	X int
	Y int
}

func (p Point) Add(x, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

// New creates a new Point.
func New(x, y int) Point {
	return Point{X: x, Y: y}
}

// Zero creates a new Point at (0, 0).
func Zero() Point {
	return Point{X: 0, Y: 0}
}
