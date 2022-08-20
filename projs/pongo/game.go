package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen tcell.Screen
	ball   Ball
}

func (g Game) Run() {
	sdefault := tcell.StyleDefault.Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	for x := 0; ; x++ {
		g.screen.Clear()

		w, h := g.screen.Size()
		g.ball.Bounce(w, h)
		g.ball.Update()
		g.screen.SetContent(g.ball.x, g.ball.y, ballRune, nil, sdefault)
		time.Sleep(30 * time.Millisecond)

		g.screen.Show()
	}
}
