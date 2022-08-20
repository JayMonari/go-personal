package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	game := Game{
		screen: screen,
		ball:   Ball{x: 1, y: 1, speedX: 1, speedY: 1},
	}
	go game.Run()

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}
