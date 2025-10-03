package types

type Direction byte

const (
	North Direction = iota
	East
	South
	West
	Undefined Direction = 255
)

var directionNames = map[Direction]string{
	North:     "North",
	East:      "East",
	South:     "South",
	West:      "West",
	Undefined: "Undefined",
}

func (d Direction) String() string {
	return directionNames[d]
}
