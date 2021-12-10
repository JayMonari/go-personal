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

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}

	for _, c := range h {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, and we are changing it to be worth 11
			// 11 - 1 = 10
			return minScore + 10
		}
	}

	return minScore
}
