package rogue

type arena struct {
	player  *player
	enemies []enemy
	height  int
	width   int
}

func newArena(p *player, h, w int) *arena {
	e1 := *spawnEnemy(coord{x: 1, y: 1})
	e2 := *spawnEnemy(coord{x: 12, y: 10})
	a := &arena{
		player:  p,
		enemies: []enemy{e1, e2},
		height:  h,
		width:   w,
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
		// a.player.current_health -= 1
		return nil
	}

	if a.player.current_health <= 1 {
		return a.player.die()
	} else {
		a.player.body.move(d)
		return nil
	}
	return nil
}
