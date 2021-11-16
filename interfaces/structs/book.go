package structs

import "fmt"

type Book struct {
	Title string
	Price float32
}

func (b Book) Print() { fmt.Printf("%-15s: %.2f\n", b.Title, b.Price) }
