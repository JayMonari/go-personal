package list

import "fmt"

type Bullet struct {
	bullet rune
}

func NewBullet(b rune) *Bullet { return &Bullet{bullet: b} }

func (b *Bullet) Print(todos []string) {
	for _, td := range todos {
		fmt.Println("\t", string(b.bullet), td)
	}
}
