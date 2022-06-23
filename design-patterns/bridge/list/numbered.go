package list

import "fmt"

type Numbered struct {
	n int
}

func NewNumbered(i int) *Numbered { return &Numbered{n: i} }

func (b *Numbered) Print(todos []string) {
	for i, td := range todos {
		fmt.Printf("\t%d. %s\n", i+b.n, td)
	}
}
