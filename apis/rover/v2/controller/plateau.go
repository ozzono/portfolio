package controller

import (
	"fmt"
	"rover/utils"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrUnavailableCoord   = errors.New("unavailable coords")
	ErrUnavailablePlateau = errors.New("no plateau was setted yet")
	ErrRoverNotFound      = errors.New("rover not found")
	Plateau               *plateau
)

type plateau struct {
	Width        int
	Height       int
	Rovers       map[*Coord]*Rover
	RoverByID    map[int]*Rover
	RoverBlocker sync.Mutex
	available    bool
}

func SetPlateau(x, y int) *plateau {
	Plateau = &plateau{
		Width:        x,
		Height:       y,
		available:    true,
		Rovers:       map[*Coord]*Rover{},
		RoverByID:    map[int]*Rover{},
		RoverBlocker: sync.Mutex{},
	}
	return Plateau
}

func (p *plateau) AddRover(r *Rover) error {
	if p == nil {
		return ErrUnavailablePlateau
	}
	if !p.available {
		return ErrUnavailablePlateau
	}

	if err := p.isCoordAllowed(r.Coords); err != nil {
		return ErrUnavailableCoord
	}

	p.Rovers[r.Coords] = r
	p.RoverByID[r.ID] = r
	return nil
}

func (p *plateau) isCoordAllowed(c *Coord) error {
	if p == nil {
		return ErrUnavailablePlateau
	}
	if !p.available {
		return ErrUnavailablePlateau
	}

	coords := p.roversCoords()
	if _, ok := coords[*c]; ok {
		return errors.Wrap(ErrUnavailableCoord, "coord occupied")
	}

	if c.X < 0 || c.X > p.Width {
		return errors.Wrapf(ErrUnavailableCoord, "x out of bounds: %d", c.X)
	}

	if c.Y < 0 || c.Y > p.Height {
		return errors.Wrapf(ErrUnavailableCoord, "y out of bounds: %d", c.Y)
	}

	return nil
}

func (p *plateau) GerRoverByID(id int) (*Rover, error) {
	rover, ok := p.RoverByID[id]
	if !ok {
		return nil, ErrRoverNotFound
	}
	return rover, nil
}

func (p *plateau) roversCoords() map[Coord]string {
	output := map[Coord]string{}
	for coord := range p.Rovers {
		output[*coord] = p.Rovers[coord].Direction
	}
	return output
}

func (p *plateau) Show() (string, error) {
	if p == nil {
		return "", ErrUnavailablePlateau
	}
	if !p.available {
		return "", ErrUnavailablePlateau
	}

	row := []string{}
	colLabels := []string{utils.LeftPad(p.Height, "X", " ")}
	coords := p.roversCoords()
	for i := Plateau.Height; i >= 0; i-- {
		cols := []string{utils.LeftPad(p.Height, fmt.Sprint(i), " ")}
		for j := 0; j <= Plateau.Width; j++ {
			if len(colLabels) <= p.Width+1 {
				colLabels = append(colLabels, utils.LeftPad(p.Width, fmt.Sprint(j), " "))
			}

			if direction, ok := coords[Coord{j, i}]; ok {
				cols = append(cols, utils.LeftPad(p.Width, pointDirection(direction), " "))
			} else {
				cols = append(cols, utils.LeftPad(p.Width, "  ", " "))
			}
		}
		row = append(row, fmt.Sprintf("|%s|", strings.Join(cols, "|")))
	}
	colLabel := fmt.Sprintf("|%s|", strings.Join(colLabels, "|"))
	return strings.Join(append(row, colLabel), "\n"), nil
}
