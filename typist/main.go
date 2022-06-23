package main

import (
	"strings"
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

type document []string

const (
	startPos = 2
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	doc := strings.Split(`tbprint(2, 2, termbox.ColorBlack, termbox.ColorWhite, doc)
	termbox.Flush()
	col := 0
	row := 0
	for {
		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyEsc {
			break
		}
		if col == len(doc[row]) {
			for i, r := range doc[row] {
				termbox.SetCell(w-len(doc[row])/2+i, h+row, r, termbox.ColorBlack, termbox.ColorGreen)
				termbox.Flush()
			}
			col = 0
			row++
		}
		r, size := utf8.DecodeRuneInString(doc[row][col:])
		if ev.Key == termbox.KeySpace {
			ev.Ch = ' '
		}
		if ev.Ch == r {
			termbox.SetCell(w-len(doc[row])/2+col, h+row, r, termbox.ColorWhite, termbox.ColorBlack)
			termbox.Flush()
			col += size
		}
	}`, "\n")
	tbprint(startPos, startPos, termbox.ColorBlack, termbox.ColorWhite, doc)
	termbox.Flush()
	col := 0
	row := 0
	for {
		// TODO(jay): Remove need to hit Tab characters.
		// TODO(jay): Wrap lines if they go off screen.
		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyEsc {
			break
		}
		if col == len(doc[row]) {
			for i, r := range doc[row] {
				termbox.SetCell(startPos+i, startPos+row, r,
					termbox.ColorBlack, termbox.ColorGreen)
				termbox.Flush()
			}
			col = 0
			row++
		}
		r, size := utf8.DecodeRuneInString(doc[row][col:])
		if ev.Key == termbox.KeySpace {
			ev.Ch = ' '
		} else if ev.Key == termbox.KeyTab {
			ev.Ch = '	'
		}
		if ev.Ch == r {
			termbox.SetCell(startPos+col, startPos+row, r,
				termbox.ColorWhite, termbox.ColorBlack)
			termbox.Flush()
			col += size
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, ss []string) {
	for _, s := range ss {
		x := x
		for _, r := range s {
			termbox.SetCell(x, y, r, fg, bg)
			x += utf8.RuneLen(r)
		}
		y++
	}
}
