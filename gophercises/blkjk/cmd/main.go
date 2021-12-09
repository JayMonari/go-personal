package main

import (
	"blkjk/deck"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	stand = "s"
	hit   = "h"
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3),
		deck.Shuffle(*rand.New(rand.NewSource(time.Now().UnixNano()))))
	return ret
}

func Deal(gs GameState) GameState {
	gsCln := clone(gs)
	gsCln.Player = make(Hand, 0, 5)
	gsCln.Dealer = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, gsCln.Deck = gsCln.Deck[0], gsCln.Deck[1:]
		gsCln.Player = append(gsCln.Player, card)
		card, gsCln.Deck = gsCln.Deck[0], gsCln.Deck[1:]
		gsCln.Dealer = append(gsCln.Dealer, card)
	}
	gsCln.State = StatePlayerTurn
	return gsCln
}

func Hit(gs GameState) GameState {
	gsCln := clone(gs)
	hand := gsCln.CurrentPlayer()
	var card deck.Card
	card, gsCln.Deck = gsCln.Deck[0], gsCln.Deck[1:]
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(gsCln)
	}
	return gsCln
}

func Stand(gs GameState) GameState {
	gsCln := clone(gs)
	gsCln.State++
	return gsCln
}

func EndHand(gs GameState) GameState {
	gsCln := clone(gs)
	pScore, dScore := gsCln.Player.Score(), gsCln.Dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", gsCln.Player, "\nScore: ", pScore)
	fmt.Println("Dealer:", gsCln.Dealer, "\nScore: ", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw game")
	}
	fmt.Println()
	gsCln.Player = nil
	gsCln.Dealer = nil
	return gsCln
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("player:", gs.Player)
			fmt.Println("dealer:", gs.Dealer.DealerString())
			fmt.Println("What do you want to do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case hit:
				gs = Hit(gs)
			case stand:
				gs = Stand(gs)
			default:
				fmt.Println("Invalid option:", input)
			}
		}

		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}
		gs = EndHand(gs)
	}
}

type State uint8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck []deck.Card
	State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it isn't any player's turn.")
	}
}

func clone(gs GameState) GameState {
	cln := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(cln.Deck, gs.Deck)
	copy(cln.Player, gs.Player)
	copy(cln.Dealer, gs.Dealer)
	return cln
}
