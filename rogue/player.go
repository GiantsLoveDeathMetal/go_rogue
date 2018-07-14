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
	body      coord
	character rune
}

func newPlayer(b coord) *player {
	return &player{
		// Position of the player
		body:      b,
		character: '@',
	}
}

func (p *player) die() error {
	return errors.New("Died")
}

func (p *player) isOnPosition(c coord) bool {
	h := p.body
	if h.x == c.x && h.y == c.y {
		return true
	}
	return false
}
