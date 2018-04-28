// Package cards provides a standard deck of playing cards.
package cards

import (
	"fmt"
	"math/rand"
)

// Rank is used to represent the rank of a card (e.g. King).
type Rank uint8

// Suit is used to represent the suit of a card (e.g. spades).
type Suit uint8

// These constants are used to represent the various suit values of playing cards.
const (
	// _             = iota
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker
)

// These constants are used to represent the various rank values of playing cards.
const (
	_        = iota
	Two Rank = iota + 1
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
	Ace
)

// A Card is a single playing card.
type Card struct {
	Suit
	Rank
}

var suits = [...]Suit{Spades, Diamonds, Clubs, Hearts}
var ranks = [...]Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

var suitStr = []string{"♠", "♦", "♣", "♥"}
var rankStr = []string{"_", "_", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// String produces a pretty string representation for a card.
func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	ss := suitStr[c.Suit]
	rs := rankStr[c.Rank]
	return fmt.Sprintf("%s %s", rs, ss)
}

// A Deck is a deck of playing cards.
type Deck struct {
	cards     []Card
	deckIndex []int // the order of the cards in the deck
}

// New creates a new unshuffled deck of 54 cards.
func New(nJokers int) Deck {
	var c = []Card{}

	// Add one card for each suit/rank combination
	i := 0
	for _, s := range suits {
		for _, r := range ranks {
			c = append(c, Card{s, r})
			i++
		}
	}
	// Add jokers
	for i := 0; i < nJokers; i++ {
		c = append(c, Card{Joker, 0})
	}

	// Initially the deck is unshuffled
	var indices = []int{}
	for i := 0; i < len(c); i++ {
		indices = append(indices, i)
	}

	deck := Deck{c, indices}

	return deck
}

// Shuffle shuffles the cards randomly.
func (d *Deck) Shuffle() {
	i := rand.Perm(54)

	d.deckIndex = i
}
