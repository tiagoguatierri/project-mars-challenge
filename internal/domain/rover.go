package domain

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var (
	ErrRoverInstructionArgsType = fmt.Errorf("rover first tow args should be integer")
)

type Rover struct {
	Name      string
	X         int
	Y         int
	Direction Direction
}

type RoverFactory struct{}

func NewRover(name string, x, y int, direction Direction) *Rover {
	return &Rover{
		Name:      name,
		X:         x,
		Y:         y,
		Direction: direction,
	}
}

func NewRoverFactory() *RoverFactory {
	return &RoverFactory{}
}

func (f *RoverFactory) Create(instructions string) (*Rover, error) {
	x, y, d, name, err := f.parseArgs(instructions)
	if err != nil {
		return nil, err
	}

	return NewRover(name, x, y, d), nil
}

func (f *RoverFactory) parseArgs(instructions string) (
	x int,
	y int,
	d Direction,
	name string,
	err error,
) {
	cmds := strings.Fields(instructions)
	argsSize := len(cmds)
	if argsSize < 4 {
		err = fmt.Errorf("plateau instructions should have four args. Given: %d", argsSize)
		return
	}

	x, y, err = f.parseCords(cmds[:2])
	if err != nil {
		return
	}

	d = Direction(cmds[2])
	if !slices.Contains(Directions, d) {
		err = fmt.Errorf("third arg should be a valid direction: Given: %s", d)
		return
	}

	name = cmds[3]

	return
}

func (f *RoverFactory) parseCords(args []string) (int, int, error) {
	parsed := make([]int, len(args))
	for i, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return -1, -1, ErrRoverInstructionArgsType
		}
		parsed[i] = val
	}
	return parsed[0], parsed[1], nil
}
