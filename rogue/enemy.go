package rogue

type enemy struct {
	body      coord
	health    int
	pattern   []direction
	character rune
}

func spawnEnemyTypeA(spawn_area coord) *enemy {
	b := spawn_area.randomCoord()
	return &enemy{
		body:   b,
		health: 2,
		pattern: []direction{
			LEFT,
			LEFT,
			LEFT,
			LEFT,
			LEFT,
			NO_MOVE,
		},
		character: '=',
	}
}

func spawnEnemyTypeB(spawn_area coord) *enemy {
	b := spawn_area.randomCoord()
	return &enemy{
		body:   b,
		health: 2,
		pattern: []direction{
			LEFT,
			UP,
			RIGHT,
			DOWN,
			NO_MOVE,
		},
		character: 'X',
	}
}

func spawnEnemyTypeC(spawn_area coord) *enemy {
	b := spawn_area.randomCoord()
	return &enemy{
		body:   b,
		health: 2,
		pattern: []direction{
			UP,
			UP,
			UP,
			UP,
			UP,
		},
		character: '|',
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

func (e *enemy) reversePattern() {
	m := make(map[direction]direction)

	// Set key pairs
	m[LEFT] = RIGHT
	m[RIGHT] = LEFT
	m[UP] = DOWN
	m[DOWN] = UP
	m[NO_MOVE] = NO_MOVE
	for i := 0; i < len(e.pattern); i++ {
		e.pattern[i] = m[e.pattern[i]]
	}

}

func (e *enemy) next_move() direction {
	d := nextMove(&e.pattern)
	return d
}

func (e *enemy) die() {
	e = nil
}
