package blackjack

import (
	"blkjk/deck"
)

type hand struct {
	cards []deck.Card
	bet   int
}
