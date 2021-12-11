package blackjack

import (
	"blkjk/deck"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	stateBet state = iota
	statePlayerTurn
	stateDealerTurn
	stateHandOver
)

type state uint8

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

type Game struct {
	// unexported fields
	nDecks          int
	nHands          int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []hand
	handIdx   int
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI AI
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
	}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 10
	}
	if opts.BlackjackPayout == 0.0 {
		opts.BlackjackPayout = 1.5
	}
	g.nDecks = opts.Decks
	g.nHands = opts.Hands
	g.blackjackPayout = opts.BlackjackPayout
	return g
}

func (g *Game) Play(ai AI) int {
	g.deck = nil
	min := 52 * g.nDecks / 3
	for i := 0; i < g.nHands; i++ {
		var shuffled bool
		if len(g.deck) < min {
			g.deck = deck.New(deck.Deck(g.nDecks), deck.Shuffle(*rand.New(rand.NewSource(time.Now().UnixNano()))))
			shuffled = true
		}
		bet(g, ai, shuffled)
		deal(g)
		if Blackjack(g.dealer...) {
			endRound(g, ai)
			continue
		}

		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(*g.currentHand()))
			copy(hand, *g.currentHand())
			move := ai.Play(hand, g.dealer[0])
			err := move(g)
			switch err {
			case errBust:
				MoveStand(g)
			case nil:
				// noop
			default:
				panic(err)
			}
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}
		endRound(g, ai)
	}
	return g.balance
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player[g.handIdx].cards
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't any player's turn.")
	}
}

var (
	errBust = errors.New("hand score exceeded 21")
)

type Move func(*Game) error

func MoveHit(g *Game) error {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = g.deck[0], g.deck[1:]
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		return errBust
	}
	return nil
}

func MoveDouble(g *Game) error {
	if len(g.player) != 2 {
		return errors.New("can only double on a hand with 2 cards")
	}
	g.playerBet *= 2
	MoveHit(g)
	return MoveStand(g)
}

func MoveSplit(g *Game) error {
	cards := g.currentHand()
	if len(*cards) != 2 {
		return errors.New("you can only split with two cards in your cards")
	}
	if (*cards)[0].Rank != (*cards)[1].Rank {
		return errors.New("both cards must have the same rank to split")
	}
	g.player = append(g.player, hand{
		cards: []deck.Card{(*cards)[1]},
		bet:   g.player[g.handIdx].bet,
	})
	g.player[g.handIdx].cards = (*cards)[:1]
	return nil
}

func MoveStand(g *Game) error {
	if g.state == stateDealerTurn {
		g.state++
		return nil
	}
	if g.state == statePlayerTurn {
		g.handIdx++
		if g.handIdx == len(g.player) {
			g.state++
		}
		return nil
	}
	return errors.New("invalid state")
}

// Score will take in a hand of cards and return the best blackjack scores
// possible with the hand.
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}

	for _, c := range hand {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, and we are changing it to be worth 11
			// 11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

// Soft returns true if Ace is being counted as 11 points.
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

// Blackjack returns true if a hand is a blackjack, that is if there are two
// cards in the hand and they sum up to 21.
func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
}

func endRound(g *Game, ai AI) {
	dScore, dBlkjk := Score(g.dealer...), Blackjack(g.dealer...)
	allHands := make([][]deck.Card, len(g.player))
	for hi, hand := range g.player {
		allHands[hi] = hand.cards
		pScore, pBlkjk := Score(hand.cards...), Blackjack(hand.cards...)
		winnings := hand.bet
		switch {
		case pBlkjk && dBlkjk:
			winnings = 0
		case dBlkjk:
			winnings = -winnings
		case pBlkjk:
			winnings = int(float64(winnings) * g.blackjackPayout)
		case pScore > 21:
			fmt.Println("You busted")
			winnings = -winnings
		case dScore > 21:
			// win
		case pScore > dScore:
			// win
		case dScore > pScore:
			fmt.Println("You lose")
			winnings = -winnings
		case dScore == pScore:
			fmt.Println("Draw game")
			winnings = 0
		}
		g.balance += winnings
	}
	ai.Results(allHands, g.dealer)
	g.player = nil
	g.dealer = nil
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	g.playerBet = bet
}

func deal(g *Game) {
	playerHand := make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = g.deck[0], g.deck[1:]
		playerHand = append(playerHand, card)
		card, g.deck = g.deck[0], g.deck[1:]
		g.dealer = append(g.dealer, card)
	}
	g.player = []hand{
		{
			cards: playerHand,
			bet:   g.playerBet,
		},
	}
	g.state = statePlayerTurn
}

func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		if int(c.Rank) > 10 {
			// Jack, Queen, King are counted as 10
			c.Rank = 10
		}
		score += int(c.Rank)
	}
	return score
}
