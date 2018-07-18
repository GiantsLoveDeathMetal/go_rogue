package rogue

import "time"
import "math/rand"

type coord struct {
	x, y int
}

func (c *coord) move(d direction) error {
	switch d {
	case RIGHT:
		c.x += 1
	case LEFT:
		c.x -= 1
	case UP:
		c.y += 1
	case DOWN:
		c.y -= 1
	case NO_MOVE:
		return nil
	}
	return nil
}

func (c *coord) onBorder(o coord) bool {
	return c.x >= o.x || c.y > o.y || c.x <= -1 || c.y <= 0
}

func (c *coord) randomCoord() coord {
	// Create seed
	s := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(s)

	// Generate new x and y
	return coord{
		x: rnd.Intn(c.x),
		y: rnd.Intn(c.y),
	}

}
