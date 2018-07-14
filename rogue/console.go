package rogue

import "github.com/nsf/termbox-go"
import "github.com/mattn/go-runewidth"

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
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
	renderQuitMessage(right, bottom)

	return termbox.Flush()
}

func renderPlayer(left, bottom int, p *player) {
	for _, b := range p.body {
		termbox.SetCell(left+b.x, bottom-b.y, ' ', playerColour, playerColour)
	}
}

func renderArena(a *arena, top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+a.width, i, '│', defaultColor, bgColor)
	}

	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+a.width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+a.width, bottom, '┘', defaultColor, bgColor)

	fill(left, top, a.width, 1, termbox.Cell{Ch: '─'})
	fill(left, bottom, a.width, 1, termbox.Cell{Ch: '─'})
}

func renderQuitMessage(right, bottom int) {
	m := "Press ESC to quit"
	tbprint(right-17, bottom+1, defaultColor, defaultColor, m)
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
