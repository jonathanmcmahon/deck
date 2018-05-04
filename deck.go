// Package deck provides a standard deck of playing cards.
package deck

import (
	"bytes"
	"errors"
	"fmt"
	mrand "math/rand"
	"time"
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
	return fmt.Sprintf("%s%s", rs, ss)
}

// A Deck is a deck of playing cards.
type Deck struct {
	cards       []Card
	cardOrder   []int // the value at each location indexes into cards
	currentCard int   // the current card (indexes into cardOrder)
	rng         *mrand.Rand
}

// New creates a new unshuffled deck comprised of nDecks of 52 cards and nJokers,
// minus any Ranks listed in omitRanks.
func New(nDecks int, nJokers int, omitRanks []Rank) Deck {
	var c = []Card{}

	// Build map of ranks to include
	var includeRanks = make(map[Rank]bool)
	for _, r := range ranks {
		includeRanks[r] = true
	}
	for _, r := range omitRanks {
		includeRanks[r] = false
	}

	// Create a multi-deck deck if needed
	for dcount := 0; dcount < nDecks; dcount++ {
		// Add one card for each suit/rank combination
		i := 0
		for _, s := range suits {
			for _, r := range ranks {
				if includeRanks[r] {
					c = append(c, Card{s, r})
					i++
				}
			}
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

	// Add RNG
	rsrc := mrand.NewSource(time.Now().UnixNano())
	rng := mrand.New(rsrc)

	deck := Deck{c, indices, 0, rng}

	return deck
}

// Shuffle shuffles the cards randomly and resets the deck.
func (d *Deck) Shuffle() {
	i := d.rng.Perm(len(d.cards))

	d.cardOrder = i
}

func (d Deck) String() string {
	var buffer bytes.Buffer

	for _, c := range d.cardOrder {
		buffer.WriteString(d.cards[c].String())
		buffer.WriteString(" ")
	}

	return fmt.Sprintf(buffer.String())
}

func (d *Deck) DrawCard() (Card, error) {
	var card Card
	currIndex := d.currentCard

	if currIndex < len(d.cards) {
		d.currentCard++
		card = d.cards[d.cardOrder[currIndex]]
		return card, nil
	} else {
		return card, errors.New("Deck is exhausted")
	}
}

func compare(a, b Deck) bool {
	if len(a.cardOrder) != len(b.cardOrder) {
		return false
	}
	for i := 0; i < len(a.cardOrder); i++ {
		if a.cardOrder[i] != b.cardOrder[i] {
			return false
		}
	}
	if a.currentCard != b.currentCard {
		return false
	}
	return true
}
