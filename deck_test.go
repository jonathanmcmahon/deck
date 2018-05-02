package deck

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

	if c.String() != "A♥" {
		t.Error("String conversion for ace of hearts was incorrect: ", c.String())
	}

	c = Card{Rank: King, Suit: Diamonds}

	if c.String() != "K♦" {
		t.Error("String conversion for king of diamonds was incorrect: ", c.String())
	}

}

func TestDeck(t *testing.T) {
	d := New(1, 2, []Rank{})

	if len(d.cards) != 54 {
		t.Error("Deck is not 54 cards; it is ", len(d.cards))
	}

	if len(d.cards) != len(d.cardOrder) {
		t.Error("Card count does not match deck index")
	}
}

func TestDeckShuffle(t *testing.T) {
	d := New(1, 2, []Rank{})

	// Shuffle cards
	pre := d.cardOrder
	d.Shuffle()
	post := d.cardOrder

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

func TestCustomJokerCount(t *testing.T) {
	nJokers := 4
	d := New(1, nJokers, []Rank{})

	// Shuffle cards
	d.Shuffle()

	// Look for difference in card ordering post-shuffle
	jokerCount := 0
	for i := 0; i < len(d.cards); i++ {
		if d.cards[i].Suit == Joker {
			jokerCount++
		}
	}
	if jokerCount != nJokers {
		t.Error("There should be", nJokers, "Jokers; there were", jokerCount)
	}
}

func TestMultipleDeck(t *testing.T) {
	nJokers := 2
	d := New(2, nJokers, []Rank{})

	nCards := 52 + 52 + nJokers

	if len(d.cards) != nCards {
		t.Error("Deck is not", nCards, "cards; it is ", len(d.cards))
	}

	if len(d.cards) != len(d.cardOrder) {
		t.Error("Card count does not match deck index")
	}

}

func TestOmitRanks(t *testing.T) {
	omitRanks := []Rank{Two, Three}

	d := New(1, 2, omitRanks)

	for _, v := range d.cards {
		for _, r := range omitRanks {
			if v.Rank == r {
				t.Error("Deck contains a card of a rank that should be omitted:", v)
			}
		}
	}
}

func TestPrintDeck(t *testing.T) {

	d := New(1, 2, []Rank{})

	fmt.Println("Printing deck: ")
	fmt.Println(d)
}

func TestDrawCard(t *testing.T) {
	d := New(1, 2, []Rank{})

	if d.currentCard != 0 {
		t.Error("Starting card is not top card:", d.currentCard)
	}

	for i := 0; i < 10; i++ {
		_, err := d.DrawCard()
		if err != nil {
			t.Error("Error while drawing card:", err)
		}
	}

	if d.currentCard != 10 {
		t.Error("Current card index is incorrect; should be 10 but it is", d.currentCard)
	}
}

func TestExhaustDeck(t *testing.T) {

	d := New(1, 2, []Rank{})
	deckLength := len(d.cards)

	for i := 0; i < deckLength; i++ {
		_, err := d.DrawCard()
		if err != nil {
			t.Error("Error while drawing card:", err)
		}
	}
	// This draw should fail
	_, err := d.DrawCard()
	if err == nil {
		t.Error("Deck should have been exhausted but was not.")
	}
}
