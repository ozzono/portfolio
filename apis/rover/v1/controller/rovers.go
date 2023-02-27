package controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type Rover struct {
	Label     string
	Coords    *Coord
	Direction string
}

var (
	ErrInvalidDirection = errors.New("invalid direction; must be N, S, W or E")
	ErrInvalidRotation  = errors.New("invalid rotation; must be L or R")
	ErrInvalidMovement  = errors.New("invalid rotation; must be L, R, F, B")
	ErrInvalidLabel     = errors.New("invalid rover label")

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
	rotations = map[string]struct{}{"L": {}, "R": {}}
	movements = map[string]struct{}{"L": {}, "R": {}, "F": {}, "B": {}}
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

func NewRover(l, d string, x, y int) (*Rover, error) {
	if d != "N" && d != "S" && d != "W" && d != "E" {
		return nil, ErrInvalidDirection
	}

	if len(l) == 0 {
		return nil, errors.Wrap(ErrInvalidLabel, "adding a label shouldn't be required, but makes things so much better")
	}

	return &Rover{Coords: &Coord{x, y}, Direction: d, Label: l}, nil
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
		log.Printf("rotating rover %s %s", r.Label, direction)
		r.Direction = direction
		return nil
	}

	Plateau.RoverBlocker.Lock()
	defer Plateau.RoverBlocker.Unlock()

	coord := &Coord{}
	var err error
	defer func() {
		if err == nil {
			log.Printf("moving %s %s %s", r.Label, pointDirection(r.Direction), movement)
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
	Label ------ %s
	Coords ----- %d %d
	Direction -- %s
	`, r.Label, r.Coords.X, r.Coords.Y, r.Direction)
}

func (r Rover) ShowPosition() string {
	row := []string{}
	colLabels := []string{" X "}
	for i := Plateau.Height; i >= 0; i-- {
		cols := []string{" " + fmt.Sprint(i) + " "}
		colLabels = append(colLabels, fmt.Sprint(Plateau.Height-i)+" ")
		for j := 0; j <= Plateau.Width; j++ {
			if r.Coords.X == j && r.Coords.Y == i {
				cols = append(cols, pointDirection(r.Direction))
			} else {
				cols = append(cols, "  ")
			}
		}
		row = append(row, fmt.Sprintf("|%s|", strings.Join(cols, "|")))
	}
	colLabel := fmt.Sprintf("|%s|", strings.Join(colLabels, "|"))
	return strings.Join(append(row, colLabel), "\n")
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
