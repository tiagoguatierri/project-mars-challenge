package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tiagoguatierri/project-mars-challenge/internal/command"
	"github.com/tiagoguatierri/project-mars-challenge/internal/domain"
)

const (
	instructionsFile = "input.txt"
)

var (
	ErrEmptyInstructions = fmt.Errorf("empty instructions")
)

func main() {
	instructionsRaw, err := loadInstructions()
	if err != nil {
		panic(err)
	}

	instructions, err := parseInstructions(instructionsRaw)
	if err != nil {
		panic(err)
	}

	plateauFactory := domain.NewPlateauFactory()
	plateau, err := plateauFactory.Create(instructions)
	if err != nil {
		panic(err)
	}

	roverFactory := domain.NewRoverFactory()
	instructions = instructions[1:]
	dispatcher := command.NewCommandDispatcher(plateau)
	result := []string{}

	for i := range instructions {
		if i%2 == 0 {
			rover, err := roverFactory.Create(instructions[i])

			if err != nil {
				log.Fatalf("error creating rover: %v", err)
				continue
			}

			if err := plateau.Place(rover); err != nil {
				log.Fatalf("error placing rover: %v", err)
				continue
			}

			dispatcher.Invoke(instructions[i+1], rover)

			result = append(
				result, fmt.Sprintf(
					"%d %d %s %s",
					rover.X,
					rover.Y,
					rover.Direction,
					rover.Name,
				),
			)
		}

	}

	outputContent := strings.Join(result, "\n")
	if err := os.WriteFile("output.txt", []byte(outputContent), 0644); err != nil {
		panic(err)
	}
}

func loadInstructions() (string, error) {
	instructions, err := os.ReadFile(instructionsFile)
	if err != nil {
		return "", fmt.Errorf("something was wrong when try opening file: %w", err)
	}

	return string(instructions), nil
}

func parseInstructions(raw string) ([]string, error) {
	parsed := strings.TrimSpace(raw)
	instructions := strings.Split(parsed, "\n")

	if len(instructions) == 0 {
		return nil, ErrEmptyInstructions
	}

	return instructions, nil
}
