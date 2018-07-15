package rogue

type enemy struct {
	body      coord
	health    int
	character rune
}

func spawnEnemy(b coord) *enemy {
	return &enemy{
		body:      b,
		health:    1,
		character: 'O',
	}
}

func (e *enemy) die() {
	e = nil
}
