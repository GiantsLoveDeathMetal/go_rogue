package rogue

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
	}
	return nil
}

func (c *coord) onBorder(o coord) bool {
	return c.x >= o.x || c.y > o.y || c.x <= -1 || c.y <= 0
}
