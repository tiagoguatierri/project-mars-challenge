package command

import (
	"slices"

	"github.com/tiagoguatierri/project-mars-challenge/internal/domain"
)

type TurnLeftCmd struct{}

func NewTurnLeftCmd() *TurnLeftCmd {
	return &TurnLeftCmd{}
}

func (c *TurnLeftCmd) Execute(rover *domain.Rover) error {
	index := slices.Index(domain.Directions, rover.Direction)
	next := (index - 1 + len(domain.Directions)) % len(domain.Directions)
	rover.Direction = domain.Directions[next]

	return nil
}
