package rogue

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	altBgColour  = termbox.ColorBlack
	playerColour = termbox.ColorWhite
)

func (g *Game) render() error {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h   = termbox.Size()
		midY   = h / 2
		left   = (w - g.arena.width) / 2
		right  = (w + g.arena.width) / 2
		top    = midY - (g.arena.height / 2)
		bottom = midY + (g.arena.height / 2) + 1
	)
	renderArena(g.arena, top, bottom, left)
	renderPlayer(left, bottom, g.arena.player)
	renderEnemies(left, bottom, g.arena.enemies)
	renderHealth(left, bottom, g.arena.player)
	renderQuitMessage(right, bottom)

	return termbox.Flush()
}

func renderEnemies(left, bottom int, es []enemy) {
	for _, e := range es {
		termbox.SetCell(
			left+e.body.x,
			bottom-e.body.y,
			e.character,
			playerColour,
			altBgColour,
		)
	}
}

func renderPlayer(left, bottom int, p *player) {
	termbox.SetCell(
		left+p.body.x,
		bottom-p.body.y,
		p.character,
		playerColour,
		altBgColour,
	)
}

func renderArena(a *arena, top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '║', defaultColor, altBgColour)
		termbox.SetCell(left+a.width, i, '║', defaultColor, altBgColour)
		for j := left; j < left+a.width; j++ {
			termbox.SetCell(j, i, '·', defaultColor, bgColor)
		}
	}

	termbox.SetCell(left-1, top, '╔', defaultColor, altBgColour)
	termbox.SetCell(left-1, bottom, '╚', defaultColor, altBgColour)
	termbox.SetCell(left+a.width, top, '╗', defaultColor, altBgColour)
	termbox.SetCell(left+a.width, bottom, '╝', defaultColor, altBgColour)

	fill(left, top, a.width, 1, termbox.Cell{Ch: '═', Bg: altBgColour})
	fill(left, bottom, a.width, 1, termbox.Cell{Ch: '═', Bg: altBgColour})
}

func renderHealth(left, bottom int, p *player) {
	hp := fmt.Sprintf("Health: %v/%v", p.current_health, p.max_health)
	tbprint(left+1, bottom+1, defaultColor, defaultColor, hp)
}

func renderQuitMessage(right, bottom int) {
	m := "Press 'q' to quit"
	tbprint(right-15, bottom+2, defaultColor, defaultColor, m)
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
