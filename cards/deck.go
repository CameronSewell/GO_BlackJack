package cards

import (
	"errors"
	"math/rand"
	"time"
)

const deckSize int = 52

type Deck struct {
	cards [deckSize]Card
	index int
}

// Make a new card deck
func NewDeck() Deck {
	d := Deck{}
	d.index = deckSize - 1
	for j := 0; j < Suits; j++ {
		for i := 0; i < CardValues; i++ {
			d.cards[i] = NewCard(i, getSuit(j))
		}
	}
	return d
}

func ShuffleDeck(d Deck) Deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d.cards {
		n := r.Intn(len(d.cards) - 1)
		d.cards[i], d.cards[n] = d.cards[n], d.cards[i]
	}
	return d
}

func HitDeck(d Deck) (*Card, error) {
	if d.index == 0 {
		return nil, errors.New("hit last card in deck")
	}
	c := &d.cards[d.index]
	d.index = d.index - 1
	return c, nil
}
