//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

func New(opts ...func([]Card) []Card) []Card {
	cc := make([]Card, 0, len(suits)*int(maxRank))
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			cc = append(cc, Card{Suit: s, Rank: r})
		}
	}
	for _, opt := range opts {
		cc = opt(cc)
	}
	return cc
}

func DefaultSort(cc []Card) []Card {
	sort.Slice(cc, Less(cc))
	return cc
}

func Sort(less func(cc []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cc []Card) []Card {
		sort.Slice(cc, less(cc))
		return cc
	}
}

func Less(cc []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cc[i]) < absRank(cc[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Shuffle(r rand.Rand) func([]Card) []Card {
	return func(cc []Card) []Card {
		shuf := make([]Card, len(cc))
		for i, j := range r.Perm(len(cc)) {
			shuf[i] = cc[j]
		}
		return shuf
	}
}

func Jokers(n int) func([]Card) []Card {
	return func(cc []Card) []Card {
		for i := 0; i < n; i++ {
			cc = append(cc, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cc
	}
}

func Filter(f func(c Card) bool) func([]Card) []Card {
	return func(cc []Card) []Card {
		var fil []Card
		for _, c := range cc {
			if !f(c) {
				fil = append(fil, c)
			}
		}
		return fil
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cc []Card) []Card {
		var dup []Card
		for i := 0; i < n; i++ {
			dup = append(dup, cc...)
		}
		return dup
	}
}
