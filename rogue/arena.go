package rogue

type arena struct {
	player  *player
	enemies []enemy
	height  int
	width   int
}

func newArena(p *player, w, h int) *arena {
	spawn_area := coord{x: w, y: h}

	e1 := *spawnEnemyTypeA(spawn_area)
	e2 := *spawnEnemyTypeB(spawn_area)
	e3 := *spawnEnemyTypeA(spawn_area)
	e4 := *spawnEnemyTypeB(spawn_area)
	e5 := *spawnEnemyTypeC(spawn_area)
	e6 := *spawnEnemyTypeA(spawn_area)
	e7 := *spawnEnemyTypeB(spawn_area)
	e8 := *spawnEnemyTypeA(spawn_area)
	e9 := *spawnEnemyTypeB(spawn_area)
	e10 := *spawnEnemyTypeC(spawn_area)
	a := &arena{
		player: p,
		enemies: []enemy{
			e1,
			e2,
			e3,
			e4,
			e5,
			e6,
			e7,
			e8,
			e9,
			e10,
		},
		height: h,
		width:  w,
	}
	return a
}

func (a *arena) move_player(d direction) error {
	h := a.player.body
	// Check new co-ordinate
	nc := coord{x: h.x, y: h.y}
	max_coord := coord{x: a.width, y: a.height}

	nc.move(d)
	if nc.onBorder(max_coord) {
		return nil
	}

	if a.player.current_health <= 1 {
		return a.player.die()
	} else if a.notEmpty(nc) {
		return nil
	} else {
		a.player.body.move(d)
		return nil
	}
	return nil
}

// Check that coord is not empty when moving player
func (a *arena) notEmpty(c coord) bool {
	for i := 0; i < len(a.enemies); i++ {
		e := &a.enemies[i]
		if e.body == c {
			e.health -= 1
			if e.health < 1 {
				a.enemies[i] = a.enemies[len(a.enemies)-1]
				a.enemies = a.enemies[:len(a.enemies)-1]
			}
			return true
		}
	}
	return false
}

// Check if a coord is in an slice of coords
func (c coord) isIn(a []coord) bool {
	for _, i := range a {
		if i == c {
			return true
		}
	}
	return false
}

// Make enemies follow their designated patterns
func (a *arena) move_enemies() error {
	var d direction
	var blocked []coord

	for i := 0; i < len(a.enemies); i++ {
		enemy := &a.enemies[i]
		d = enemy.next_move()

		nc := coord{x: enemy.body.x, y: enemy.body.y}
		max_coord := coord{x: a.width, y: a.height}
		nc.move(d)

		if a.player.isOnPosition(nc) {
			a.player.current_health -= 1
		} else if nc.onBorder(max_coord) {
			enemy.reversePattern()
		} else if nc.isIn(blocked) {
			return nil
		} else {
			enemy.body.move(d)
			blocked = append(blocked, nc)
		}
	}
	return nil
}
