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

func (a *arena) movePlayer() error {
	if err := a.player.move(); err != nil {
		return err
	}
	return nil
}
