package direction

import "math"

type Direction string

const (
	None      Direction = "none"
	North     Direction = "north"
	South     Direction = "south"
	West      Direction = "west"
	East      Direction = "east"
	NorthWest Direction = "north_west"
	NorthEast Direction = "north_east"
	SouthWest Direction = "south_west"
	SouthEast Direction = "south_east"
)

func Get(sX, sY, dX, dY int) Direction {
	deltaX := float64(dX - sX)
	deltaY := float64(dY - sY)

	if deltaX == 0 && deltaY == 0 {
		return None
	}

	angle := math.Atan2(deltaY, deltaX) * (180 / math.Pi)
	if angle < 0 {
		angle += 360
	}

	switch {
	case angle >= 337.5 || angle < 22.5:
		return East
	case angle >= 22.5 && angle < 67.5:
		return SouthEast
	case angle >= 67.5 && angle < 112.5:
		return South
	case angle >= 112.5 && angle < 157.5:
		return SouthWest
	case angle >= 157.5 && angle < 202.5:
		return West
	case angle >= 202.5 && angle < 247.5:
		return NorthWest
	case angle >= 247.5 && angle < 292.5:
		return North
	case angle >= 292.5 && angle < 337.5:
		return NorthEast
	default:
		return None
	}
}
