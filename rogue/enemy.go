package rogue

type enemy struct {
	body      coord
	health    int
	pattern   []direction
	character rune
}

func spawnEnemy(b coord) *enemy {
	return &enemy{
		body:      b,
		health:    1,
		pattern:   []direction{LEFT, LEFT, NO_MOVE, RIGHT, RIGHT, NO_MOVE},
		character: 'O',
	}
}

func nextMove(p *[]direction) direction {
	// Returns the first element in enemy pattern
	// after popping it and appending it to the end.
	var move direction
	move, *p = (*p)[0], (*p)[1:]
	*p = append(*p, move)
	return move
}

func (e *enemy) next_move() direction {
	d := nextMove(&e.pattern)
	return d
}

func (e *enemy) die() {
	e = nil
}
