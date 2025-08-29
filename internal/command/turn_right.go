package command

import (
	"slices"

	"github.com/tiagoguatierri/project-mars-challenge/internal/domain"
)

type TurnRightCmd struct{}

func NewTurnRightCmd() *TurnRightCmd {
	return &TurnRightCmd{}
}

func (c *TurnRightCmd) Execute(rover *domain.Rover) error {
	index := slices.Index(domain.Directions, rover.Direction)
	next := (index + 1 + len(domain.Directions)) % len(domain.Directions)
	rover.Direction = domain.Directions[next]

	return nil
}
