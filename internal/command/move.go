package command

import (
	"github.com/tiagoguatierri/project-mars-challenge/internal/domain"
)

type MoveCmd struct {
	plateau *domain.Plateau
}

func NewMoveCmd(plateau *domain.Plateau) *MoveCmd {
	return &MoveCmd{
		plateau: plateau,
	}
}

func (c *MoveCmd) Execute(rover *domain.Rover) error {
	delta := domain.Delta[rover.Direction]
	nX := rover.X + delta.X
	nY := rover.Y + delta.Y

	return c.plateau.Move(rover, nX, nY)
}
