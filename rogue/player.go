package rogue

import "errors"

type player struct {
	body           coord
	max_health     int
	current_health int
	character      rune
}

func newPlayer(b coord) *player {
	return &player{
		// Position of the player
		body:           b,
		max_health:     3,
		current_health: 3,
		character:      '@',
	}
}

func (p *player) die() error {
	return errors.New("You Died!")
}

func (p *player) isOnPosition(c coord) bool {
	h := p.body
	if h.x == c.x && h.y == c.y {
		return true
	}
	return false
}
