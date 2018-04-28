package cards

import (
	"fmt"
	"testing"
)

func TestCards(t *testing.T) {
	fmt.Println(Card{Rank: Ace, Suit: Hearts})
	fmt.Println(Card{Rank: Two, Suit: Spades})
	fmt.Println(Card{Rank: Nine, Suit: Diamonds})
	fmt.Println(Card{Rank: Jack, Suit: Clubs})
	fmt.Println(Card{Suit: Joker})

	c := Card{Rank: Two, Suit: Joker}

	if c.Rank != 2 {
		t.Error("Wrong value for two.")
	}

	if c.String() != "Joker" {
		t.Error("String conversion for joker was incorrect: ", c.String())
	}

	c = Card{Rank: Ace, Suit: Hearts}

	if c.Rank != 14 {
		t.Error("Wrong value for ace.")
	}

	if c.String() != "A ♥" {
		t.Error("String conversion for ace of hearts was incorrect: ", c.String())
	}

	c = Card{Rank: King, Suit: Diamonds}

	if c.String() != "K ♦" {
		t.Error("String conversion for king of diamonds was incorrect: ", c.String())
	}

}

func TestDeck(t *testing.T) {
	d := newDeck()

	if len(d.cards) != 54 {
		t.Error("Deck is not 54 cards; it is ", len(d.cards))
	}

	if len(d.cards) != len(d.deckIndex) {
		t.Error("Card count does not match deck index")
	}

	fmt.Println(d.deckIndex)
}

func TestDeckShuffle(t *testing.T) {
	d := newDeck()

	// Shuffle cards
	pre := d.deckIndex
	d.Shuffle()
	post := d.deckIndex

	// Look for difference in card ordering post-shuffle
	diff := false
	for i := 0; i < len(pre); i++ {
		if pre[i] != post[i] {
			diff = true
			break
		}
	}
	if !diff {
		t.Error("Card order is the same after shuffling")
	}
}
