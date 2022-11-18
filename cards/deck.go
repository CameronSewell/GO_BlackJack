package cards

import (
	"errors"
	"math/rand"
	"time"
)

// Define the deck size as 52
const deckSize int = 52

// Card deck definition
type Deck struct {
	cards [deckSize]Card
	index int
}

// Number of cards left
func (d *Deck) CardsLeft() int {
	return d.index + 1
}

// Make a 52 new card deck with all 4
// suits and 14 cards in each suit
func NewDeck() Deck {
	d := Deck{}
	d.index = deckSize - 1
	//Go over all possible suit combinations
	for j := 0; j < Suits; j++ {
		//Go over all possible card values
		for i := 0; i < CardValues; i++ {
			d.cards[i] = NewCard(i, getSuit(j))
		}
	}
	return d
}

// Randomly shuffle the deck
func (d *Deck) ShuffleDeck() {
	//Initialize randomizer
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//Randomly shuffle all cards
	for i := range d.cards {
		//Generate the index to swap
		n := r.Intn(len(d.cards) - 1)
		//Swap the card with the card at the generated
		//index
		d.cards[i], d.cards[n] = d.cards[n], d.cards[i]
	}
}

// Hit the deck and decrease the number of remaining cards
func (d *Deck) HitDeck() (*Card, error) {
	//Throw error if no cards left
	if d.CardsLeft() == 0 {
		return nil, errors.New("hit last card in deck")
	}
	c := &d.cards[d.index]
	d.index = d.index - 1
	return c, nil
}
