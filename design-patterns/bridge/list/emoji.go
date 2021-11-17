package list

import (
	"fmt"
	"math/rand"
	"time"
)

type Emoji struct{}

func NewEmoji() *Emoji { return &Emoji{} }

func (b *Emoji) Print(todos []string) {
	rand.Seed(time.Now().UnixNano())
	for _, td := range todos {
		emoji := rand.Intn(37) + 0x1f604
		fmt.Printf("\t %c %s\n", emoji, td)
	}
}
