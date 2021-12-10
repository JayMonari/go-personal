package blackjack

import (
	"blkjk/deck"
	"strings"
)

// TODO: check for Ace and Face card to say BlackJack
type Hand []deck.Card

func (h Hand) String() string {
	ss := make([]string, len(h))
	for i := range h {
		ss[i] = h[i].String()
	}
	return strings.Join(ss, ", ")
}

func (h Hand) DealerString() string { return h[0].String() + ", **HIDDEN**" }

