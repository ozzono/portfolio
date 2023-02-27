package controller

import "log"

// Coordinates
type Coord struct {
	X int
	Y int
}

func (c Coord) MoveUpward() *Coord {
	c.Y++
	return &Coord{c.X, c.Y}
}

func (c Coord) MoveDownward() *Coord {
	c.Y--
	return &Coord{c.X, c.Y}
}

func (c Coord) MoveRight() *Coord {
	c.X++
	return &Coord{c.X, c.Y}
}

func (c Coord) MoveLeft() *Coord {
	c.X--
	return &Coord{c.X, c.Y}
}

func (c *Coord) Reset(x, y int) {
	log.Printf("moving from [%d,%d] to [%d,%d]", c.X, c.Y, x, y)
	c.X = x
	c.Y = y
}
