package controller

import (
	"fmt"
	"log"
	"rover/utils"
	"strings"

	"github.com/pkg/errors"
)

type Rover struct {
	ID        int    `json:"id"`
	Coords    *Coord `json:"coords"`
	Direction string `json:"direction"`
}

var (
	ErrInvalidDirection = errors.New("invalid direction; must be N, S, W or E")
	ErrInvalidRotation  = errors.New("invalid rotation; must be L or R")
	ErrInvalidMovement  = errors.New("invalid rotation; must be L, R, F, B")

	compass = map[string]int{
		"N": 0,
		"E": 1,
		"S": 2,
		"W": 3,
	}
	compassIndex = map[int]string{
		0: "N",
		1: "E",
		2: "S",
		3: "W",
	}
	rotations  = map[string]struct{}{"L": {}, "R": {}}
	movements  = map[string]struct{}{"L": {}, "R": {}, "F": {}, "B": {}}
	roverCount = 0
)

func Rotate(cardinalDirection, rotation string) (string, error) {
	index, ok := compass[cardinalDirection]
	if !ok {
		return "", ErrInvalidDirection
	}

	if _, ok := rotations[rotation]; !ok {
		return "", ErrInvalidRotation
	}

	if rotation == "R" {
		index++
	}
	if index > 3 {
		index = 0
	}

	if rotation == "L" {
		index--
	}
	if index < 0 {
		index = 3
	}

	return compassIndex[index], nil
}

func NewRover(x, y int, d string) (*Rover, error) {
	if d != "N" && d != "S" && d != "W" && d != "E" {
		return nil, ErrInvalidDirection
	}

	defer func() {
		roverCount++
	}()

	return &Rover{Coords: &Coord{x, y}, Direction: d, ID: roverCount}, nil
}

func (r *Rover) Move(movement string) error {
	if _, ok := movements[movement]; !ok {
		return ErrInvalidMovement
	}

	if _, ok := rotations[movement]; ok {
		direction, err := Rotate(r.Direction, movement)
		if err != nil {
			return errors.Wrapf(err, "rotate direction: %s movement %s", r.Direction, movement)
		}
		log.Printf("rotating rover %d %s", r.ID, direction)
		r.Direction = direction
		return nil
	}

	Plateau.RoverBlocker.Lock()
	defer Plateau.RoverBlocker.Unlock()

	coord := &Coord{}
	var err error
	defer func() {
		if err == nil {
			log.Printf("moving rover %d %s %s", r.ID, pointDirection(r.Direction), movement)
		}
	}()
	switch true {
	case r.Direction == "N" && movement == "F":
		coord = r.Coords.MoveUpward()
	case r.Direction == "E" && movement == "F":
		coord = r.Coords.MoveRight()
	case r.Direction == "S" && movement == "F":
		coord = r.Coords.MoveDownward()
	case r.Direction == "W" && movement == "F":
		coord = r.Coords.MoveLeft()
	case r.Direction == "N" && movement == "B":
		coord = r.Coords.MoveDownward()
	case r.Direction == "E" && movement == "B":
		coord = r.Coords.MoveLeft()
	case r.Direction == "S" && movement == "B":
		coord = r.Coords.MoveUpward()
	case r.Direction == "W" && movement == "B":
		coord = r.Coords.MoveRight()
	}

	if err = Plateau.isCoordAllowed(coord); err != nil {
		return err
	}

	r.Coords.Reset(coord.X, coord.Y)

	return nil
}

func (r Rover) LogTxt() string {
	return fmt.Sprintf(`
	ID --------- %d
	Coords ----- [%d,%d]
	Direction -- %s
	`, r.ID, r.Coords.X, r.Coords.Y, r.Direction)
}

func (r Rover) Position() (string, error) {
	if Plateau == nil {
		return "", ErrUnavailablePlateau
	}
	if !Plateau.available {
		return "", ErrUnavailablePlateau
	}

	row := []string{}
	colLabels := []string{utils.LeftPad(Plateau.Height, "X", " ")}
	for i := Plateau.Height; i >= 0; i-- {
		cols := []string{utils.LeftPad(Plateau.Height, fmt.Sprint(i), " ")}
		for j := 0; j <= Plateau.Width; j++ {
			if len(colLabels) <= Plateau.Width+1 {
				colLabels = append(colLabels, utils.LeftPad(Plateau.Width, fmt.Sprint(j), " "))
			}

			if r.Coords.X == j && r.Coords.Y == i {
				cols = append(cols, utils.LeftPad(Plateau.Width-1, pointDirection(r.Direction), " "))
			} else {
				cols = append(cols, utils.LeftPad(Plateau.Width-1, "  ", " "))
			}
		}
		row = append(row, fmt.Sprintf("|%s|", strings.Join(cols, "|")))
	}
	colLabel := fmt.Sprintf("|%s|", strings.Join(colLabels, "|"))
	return strings.Join(append(row, colLabel), "\n"), nil
}

func pointDirection(input string) string {
	switch input {
	case "N":
		return `/\`
	case "E":
		return `->`
	case "S":
		return `\/`
	case "W":
		return `<-`
	default:
		return "NA"
	}
}
