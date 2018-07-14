package rogue

import "errors"

// player movement directions
const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type direction int

type player struct {
	body      []coord
	direction direction
	length    int
}

func newPlayer(d direction, b []coord) *player {
	return &player{
		length:    len(b),
		body:      b,
		direction: d,
	}
}

func (p *player) changeDirection(d direction) {
	p.direction = d
}

func (p *player) head() coord {
	return p.body[len(p.body)-1]
}

func (p *player) die() error {
	return errors.New("Died")
}

func (p *player) move() error {
	h := p.head()
	c := coord{x: h.x, y: h.y}

	switch p.direction {
	case RIGHT:
		c.x++
	case LEFT:
		c.x--
	case UP:
		c.y++
	case DOWN:
		c.y--
	}
	if p.isOnPosition(c) {
		return p.die()
	}
	return nil
}

func (p *player) isOnPosition(c coord) bool {
	for _, b := range p.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}
