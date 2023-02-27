package controller

import (
	"sync"

	"github.com/pkg/errors"
)

const (
	// as defined in the description
	roverLimit = 4
)

var (
	ErrUnavailableCoord   = errors.New("unavailable coords")
	ErrReachedRoverLimit  = errors.New("reached rover limit")
	ErrUnavailablePlateau = errors.New("no plateau was setted yet")
	Plateau               *plateau
)

type plateau struct {
	Width        int
	Height       int
	Rovers       map[*Coord]*Rover
	RoverBlocker sync.Mutex
	available    bool
}

func SetPlateau(x, y int) *plateau {
	Plateau = &plateau{Width: x, Height: y, available: true, Rovers: map[*Coord]*Rover{}, RoverBlocker: sync.Mutex{}}
	return Plateau
}

func (p *plateau) AddRover(r *Rover) error {
	if !p.available {
		return ErrUnavailablePlateau
	}

	if err := p.isCoordAllowed(r.Coords); err != nil {
		return ErrUnavailableCoord
	}

	if len(p.Rovers) >= roverLimit {
		return ErrReachedRoverLimit
	}

	p.Rovers[r.Coords] = r
	return nil
}

func (p *plateau) isCoordAllowed(c *Coord) error {
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

func (p *plateau) roversCoords() map[Coord]struct{} {
	output := map[Coord]struct{}{}
	for coord := range p.Rovers {
		output[*coord] = struct{}{}
	}
	return output
}
