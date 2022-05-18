package main

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	w, h := termbox.Size()
	s := []string{"Get at ya boi!", "This is the second line."}
	tbprint(w/2, h/2, termbox.ColorBlack, termbox.ColorWhite, s)
	termbox.Flush()
	offset := 0
	for {
		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyEsc {
			break
		}
		r, _ := utf8.DecodeRuneInString(s[0][offset:])
		if ev.Key == termbox.KeySpace {
			ev.Ch = ' '
		}
		if ev.Ch == r {
			termbox.SetCell(w/2-len(s[0])/2+offset, h/2, r, termbox.ColorWhite, termbox.ColorBlack)
			termbox.Flush()
			offset++
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
