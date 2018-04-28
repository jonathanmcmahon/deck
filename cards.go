// A deck of cards
package cards

import (
	"fmt"
	"math/rand"
)

type Rank uint8
type Suit uint8

const (
	// _             = iota
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
	Joker
)

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

type Card struct {
	Suit
	Rank
}

var suits = [...]Suit{Spades, Hearts, Clubs, Diamonds}
var ranks = [...]Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

var suitStr = []string{"♠", "♥", "♣", "♦"}
var rankStr = []string{"_", "_", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// Pretty string representation for a card
func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	ss := suitStr[c.Suit]
	rs := rankStr[c.Rank]
	return fmt.Sprintf("%s %s", rs, ss)
}

// A deck of cards
type Deck struct {
	cards     []Card
	deckIndex []int // the order of the cards in the deck
}

// Create a new unshuffled deck
func newDeck() Deck {
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
	c = append(c, Card{Joker, 0})
	c = append(c, Card{Joker, 0})

	// Initially the deck is unshuffled
	var indices = []int{}
	for i := 0; i < len(c); i++ {
		indices = append(indices, i)
	}

	deck := Deck{c, indices}

	return deck
}

// Shuffle the cards
func (d *Deck) Shuffle() {
	i := rand.Perm(54)

	d.deckIndex = i
}

// func (d *Deck) DrawCard() Card {
// 	c, d. = d[0], d[1:]
// 	return c
// }
