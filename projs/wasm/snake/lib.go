package main

import (
	"math/rand"

	"github.com/oleiade/lane/v2"
	"golang.org/x/exp/constraints"
)

//go:generate go run github.com/dmarkham/enumer -type=Direction
type (
	Direction uint8
	Position  struct{ x, y uint16 }
	Game      struct {
		width  uint16
		food   Position
		height uint16
		snake  *lane.List[Position] // head is the first item, tail is the last item
		dir    Direction
		done   bool
	}
)

const (
	Top Direction = iota + 1
	Right
	Bottom
	Left
)

func NewGame(width, height uint16) Game {
	l := lane.New[Position]()
	return Game{
		width:  width,
		height: height,
		snake:  l,
		dir:    Left,
		food:   Position{x: min(2, width-1), y: height / 2},
	}
}

func (g *Game) changeDir(d Direction) {
	switch {
	case (g.dir == Bottom && d != Top) ||
		(g.dir == Top && d != Bottom) ||
		(g.dir == Right && d != Left) ||
		(g.dir == Left && d != Right):
		g.dir = d
	}
}

func (g Game) isValid(pos Position) bool {
	return pos.x < g.width && pos.y < g.height
}

func (g *Game) tick() {
	if g.done || g.snake.Len() == 0 {
		return
	}

	newHead := g.snake.Front().Value
	switch g.dir {
	case Top:
		newHead.y--
	case Left:
		newHead.x--
	case Bottom:
		newHead.y++
	case Right:
		newHead.x++
	}
	if !g.isValid(newHead) || g.hitSelf(newHead) {
		g.done = true
		return
	}
	if newHead != g.food {
		g.snake.Remove(g.snake.Back())
		return
	}

	freePosns := make([]Position, 0, g.height*g.width)
	for i, x := uint16(0), make([]Position, g.width); i < g.height; i++ {
		for j := uint16(0); j < g.width; j++ {
			x[j] = Position{y: i, x: j}
		}
		freePosns = append(freePosns, x...)
	}
	freePosns = filter(freePosns, func(p Position) bool {
		for node := g.snake.Front(); node != nil; node = node.Next() {
			v := node.Value
			if v == p {
				return true
			}
		}
		return false
	})
	if len(freePosns) == 0 {
		g.done = true
	}
	p := Position{
		x: uint16(rand.Int31n(int32(g.width))),
		y: uint16(rand.Int31n(int32(g.height))),
	}
	g.food = p
	g.snake.PushFront(newHead)
}

func (g Game) hitSelf(pos Position) bool {
	for head := g.snake.Front(); head != g.snake.Back(); head = head.Next() {
		node := head.Value
		if node == pos {
			return true
		}
	}
	return false
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func filter[T any](a []T, keep func(T) bool) []T {
	i := 0
	for _, x := range a {
		if keep(x) {
			a[i] = x
			i++
		}
	}
	return a[:i]
}
