package domain

const (
	North = Direction("N")
	East  = Direction("E")
	South = Direction("S")
	West  = Direction("W")
)

var (
	Directions = []Direction{"N", "E", "S", "W"}

	Short = map[string]Direction{
		"N": North,
		"E": East,
		"S": South,
		"W": West,
	}

	Delta = map[Direction]Coord{
		North: {X: 0, Y: 1},
		East:  {X: 1, Y: 0},
		South: {X: 0, Y: -1},
		West:  {X: -1, Y: 0},
	}
)

type Direction string
