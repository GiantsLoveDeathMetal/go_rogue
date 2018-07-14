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
