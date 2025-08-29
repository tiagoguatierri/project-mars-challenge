package domain

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrPlaceAlreadyOccupied       = fmt.Errorf("place already occupied")
	ErrOutOfBounds                = fmt.Errorf("out of bounds")
	ErrPlateauInstructionArgsSize = fmt.Errorf("plateau instructions should have two args")
	ErrPlateauInstructionArgsType = fmt.Errorf("plateau args should be integer")
)

type Plateau struct {
	minX     int
	minY     int
	maxX     int
	maxY     int
	occupied map[string]*Rover
}

type PlateauFactory struct{}

func NewPlateau(maxX, maxY int) *Plateau {
	return &Plateau{
		minX:     0,
		minY:     0,
		maxX:     maxX,
		maxY:     maxY,
		occupied: make(map[string]*Rover),
	}
}

func (p *Plateau) IsOccupied(x, y int) bool {
	key := p.key(x, y)
	_, ok := p.occupied[key]
	return ok
}

func (p *Plateau) IsOutOfBound(x, y int) bool {
	return x < p.minX || x > p.maxX || y < p.minY || y > p.maxY
}

func (p *Plateau) Place(rover *Rover) error {
	rX := rover.X
	rY := rover.Y
	if p.IsOutOfBound(rX, rY) {
		return ErrOutOfBounds
	}

	if p.IsOccupied(rX, rY) {
		return ErrPlaceAlreadyOccupied
	}

	key := p.key(rover.X, rover.Y)
	p.occupied[key] = rover
	return nil
}

func (p *Plateau) Move(rover *Rover, newX, newY int) error {
	if p.IsOutOfBound(newX, newY) {
		return ErrOutOfBounds
	}
	if p.IsOccupied(newX, newY) {
		return ErrPlaceAlreadyOccupied
	}

	delete(p.occupied, p.key(rover.X, rover.Y))

	rover.X, rover.Y = newX, newY

	key := p.key(rover.X, rover.Y)
	p.occupied[key] = rover

	return nil
}

func (p *Plateau) key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func NewPlateauFactory() *PlateauFactory {
	return &PlateauFactory{}
}

func (f *PlateauFactory) Create(instructions []string) (*Plateau, error) {
	args, err := f.parseArgs(instructions)
	if err != nil {
		return nil, err
	}

	return NewPlateau(args[0], args[1]), nil
}

func (f *PlateauFactory) parseArgs(args []string) ([]int, error) {
	platArgs := strings.Fields(args[0])
	if len(platArgs) < 2 {
		return nil, ErrPlateauInstructionArgsSize
	}

	var parsedArgs []int
	for _, arg := range platArgs {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return nil, ErrPlateauInstructionArgsType
		}
		parsedArgs = append(parsedArgs, val)
	}

	return parsedArgs, nil
}
