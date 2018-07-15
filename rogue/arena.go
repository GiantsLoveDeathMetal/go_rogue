package rogue

type arena struct {
	player *player
	height int
	width  int
}

func newArena(p *player, h, w int) *arena {
	a := &arena{
		player: p,
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
