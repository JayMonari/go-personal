package structs

import "fmt"

type Puzzle struct {
	Title      string
	Price      float32
	Difficulty Difficulty
}

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

func (d Difficulty) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

func (p Puzzle) Print() { fmt.Printf("%-15s: %.2f -- Difficulty %+v", p.Title, p.Price, p.Difficulty) }
