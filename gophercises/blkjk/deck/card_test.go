package deck_test

import (
	"deck"
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(deck.Card{Rank: deck.Ace, Suit: deck.Heart})
	fmt.Println(deck.Card{Rank: deck.Two, Suit: deck.Spade})
	fmt.Println(deck.Card{Suit: deck.Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Joker
}

func TestNew(t *testing.T) {
	if len(deck.New()) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cc := deck.New(deck.DefaultSort)
	want := deck.Card{Rank: deck.Ace, Suit: deck.Spade}
	if cc[0] != want {
		t.Error("Wanted deck.Ace of Spades, got: ", cc[0])
	}
}

func TestJokers(t *testing.T) {
	got := 0
	for _, c := range deck.New(deck.Jokers(3)) {
		if c.Suit == deck.Joker {
			got++
		}
	}
	if got != 3 {
		t.Error("Want 3 Jokers, got:", got)
	}
}

func TestFilter(t *testing.T) {
	filFn := func(c deck.Card) bool {
		return c.Rank == deck.Two || c.Rank == deck.Three
	}
	for _, c := range deck.New(deck.Filter(filFn)) {
		if c.Rank == deck.Two || c.Rank == deck.Three {
			t.Error("Got Twos or Threes")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := deck.New(deck.Deck(3))
	// 13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("Want %d cards, got %d cards.", 13*4*3, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	want := deck.New()
	got := deck.New(deck.Shuffle(*rand.New(rand.NewSource(0))))
	if want[40] != got[0] {
		t.Errorf("Wanted %s got %s.", want[40], got[0])
	}
}
