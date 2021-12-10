package main

import (
	"blkjk/blackjack"
	"blkjk/deck"
	"fmt"
	"math/rand"
	"time"
)

const (
	stand = "s"
	hit   = "h"
)

func Shuffle(gs blackjack.GameState) blackjack.GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3),
		deck.Shuffle(*rand.New(rand.NewSource(time.Now().UnixNano()))))
	return ret
}

func Deal(gs blackjack.GameState) blackjack.GameState {
	gsCln := clone(gs)
	gsCln.Player = make(blackjack.Hand, 0, 5)
	gsCln.Dealer = make(blackjack.Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, gsCln.Deck = gsCln.Deck[0], gsCln.Deck[1:]
		gsCln.Player = append(gsCln.Player, card)
		card, gsCln.Deck = gsCln.Deck[0], gsCln.Deck[1:]
		gsCln.Dealer = append(gsCln.Dealer, card)
	}
	gsCln.State = blackjack.StatePlayerTurn
	return gsCln
}

func Hit(gs blackjack.GameState) blackjack.GameState {
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

func Stand(gs blackjack.GameState) blackjack.GameState {
	gsCln := clone(gs)
	gsCln.State++
	return gsCln
}

func EndHand(gs blackjack.GameState) blackjack.GameState {
	gsCln := clone(gs)
	pScore, dScore := gsCln.Player.Score(), gsCln.Dealer.Score()
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
	var gs blackjack.GameState
	gs = Shuffle(gs)
	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		var input string
		for gs.State == blackjack.StatePlayerTurn {
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

		for gs.State == blackjack.StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}
		gs = EndHand(gs)
	}
}

func clone(gs blackjack.GameState) blackjack.GameState {
	cln := blackjack.GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(blackjack.Hand, len(gs.Player)),
		Dealer: make(blackjack.Hand, len(gs.Dealer)),
	}
	copy(cln.Deck, gs.Deck)
	copy(cln.Player, gs.Player)
	copy(cln.Dealer, gs.Dealer)
	return cln
}
