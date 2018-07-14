package rogue

import "time"
import "github.com/nsf/termbox-go"

var keyboardEventsChan = make(chan keyboardEvent)

// Game instance
type Game struct {
	arena  *arena
	isOver bool
}

func initialPlayer() *player {
	return newPlayer(RIGHT, []coord{
		coord{x: 1, y: 1},
	})
}

func (g *Game) end() {
	g.isOver = true
}

func initialArena() *arena {
	return newArena(initialPlayer(), 20, 50)
}

func (g *Game) moveInterval() time.Duration {
	ms := 100 - (10)
	return time.Duration(ms) * time.Millisecond
}

func NewGame() *Game {
	return &Game{arena: initialArena()}
}

func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.arena.player.changeDirection(d)
				g.arena.movePlayer()
			case RETRY:
				continue
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.arena.movePlayer(); err != nil {
					g.end()
				}
			}
			if err := g.render(); err != nil {
				panic(err)
			}

			time.Sleep(g.moveInterval())
		}
	}
}
