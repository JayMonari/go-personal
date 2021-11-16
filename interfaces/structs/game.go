package structs

import "fmt"

type Game struct {
	Title string
	Price float32
}

func (g Game) Print() {
	fmt.Printf("%-15s: %.2f\n", g.Title, g.Price)
}

func (g *Game) discount(ratio float32) {
	g.Price *= (1 - ratio)
}
