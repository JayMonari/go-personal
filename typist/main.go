package main

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

type document []string

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	w, h := termbox.Size()
	w, h = w/2, h/2
	doc := document{"Get at ya boi!", "This is the second line.",
		"This is the third line.",
		"This is the fourth line.",
		"This is the fifth line.",
		"This is the sixth line.",
		"This is the seventh line.",
	}
	tbprint(w, h, termbox.ColorBlack, termbox.ColorWhite, doc)
	termbox.Flush()
	xOffset := 0
	yOffset := 0
	for {
		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyEsc {
			break
		}
		if xOffset == len(doc[yOffset]) {
			for i, r := range doc[yOffset] {
				termbox.SetCell(w-len(doc[yOffset])/2+i, h+yOffset, r, termbox.ColorBlack, termbox.ColorGreen)
				termbox.Flush()
			}
			xOffset = 0
			yOffset++
		}
		r, size := utf8.DecodeRuneInString(doc[yOffset][xOffset:])
		if ev.Key == termbox.KeySpace {
			ev.Ch = ' '
		}
		if ev.Ch == r {
			termbox.SetCell(w-len(doc[yOffset])/2+xOffset, h+yOffset, r, termbox.ColorWhite, termbox.ColorBlack)
			termbox.Flush()
			xOffset += size
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, ss []string) {
	for _, s := range ss {
		x := x - len(s)/2
		for _, r := range s {
			termbox.SetCell(x, y, r, fg, bg)
			x += utf8.RuneLen(r)
		}
		y++
	}
}
