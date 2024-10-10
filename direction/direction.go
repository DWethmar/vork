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

	// Adjust the angle by 22.5 degrees to align with direction sectors
	angle += 22.5
	if angle >= 360 {
		angle -= 360
	}

	// Map the angle to an index in the directions array
	index := int(angle / 45.0)
	directions := []Direction{
		East,
		SouthEast,
		South,
		SouthWest,
		West,
		NorthWest,
		North,
		NorthEast,
	}

	return directions[index%8]
}
