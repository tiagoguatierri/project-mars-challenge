package command

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"

	"github.com/tiagoguatierri/project-mars-challenge/internal/domain"
)

var (
	ErrEmptyCommands = fmt.Errorf("empty commands")
)

type Command interface {
	Execute(rover *domain.Rover) error
}

type CommandDispatcher struct {
	cmdMap map[string]Command
}

func NewCommandDispatcher(plateau *domain.Plateau) *CommandDispatcher {
	return &CommandDispatcher{
		cmdMap: map[string]Command{
			"L": NewTurnLeftCmd(),
			"R": NewTurnRightCmd(),
			"M": NewMoveCmd(plateau),
		},
	}
}
func (d *CommandDispatcher) Invoke(cmd string, rover *domain.Rover) error {
	validCmds := slices.Sorted(maps.Keys(d.cmdMap))
	cmds := strings.Split(cmd, "")

	if len(cmds) == 0 {
		return ErrEmptyCommands
	}

	for _, c := range cmds {
		if !slices.Contains(validCmds, c) {
			log.Fatalf("invalid command: %s", c)
			continue
		}
		command := d.cmdMap[c]
		if err := command.Execute(rover); err != nil {
			return err
		}
	}

	return nil
}
