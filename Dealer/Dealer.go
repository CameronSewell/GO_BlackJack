package dealer

import (
	"log"
	"main/cards"
	"main/guistate"
	"time"
)

// Defines a dealer with a total, their
// hand, and the deck to use for the game
type Dealer struct {
	Hand cards.Hand
	deck cards.Deck
}

// Creates a new dealer struct and gives them
// The two starting cards of their hand
func NewDealer() Dealer {
	dealer := Dealer{
		Hand: cards.NewHand(),
		deck: cards.NewDeck(),
	}
	dealer.deck.ShuffleDeck()
	return dealer
}

// Deal two cards to the given hand to start the game
func (d *Dealer) DealStartingHand(hand *cards.Hand, isPlayer bool) {
	for i := 0; i < 2; i++ {
		hand.AddCard(d.DealCard(), isPlayer)
	}
}

// Dealers a card from the dealer's remaining deck
func (d *Dealer) DealCard() cards.Card {
	card, err := d.deck.HitDeck()
	if err != nil {
		log.Fatal(err)
	}
	return *card
}

// Makes the dealer hit the deck until their total
// reaches or breaks 17. When that happens, the dealer stops
// Taking hits
func (d *Dealer) DealerPlay() {
	for d.Hand.GetHandTotal() < 17 {
		d.Hand.AddCard(d.DealCard(), false)
		guistate.SetCards(d.Hand, guistate.DealerHand, true)
		time.Sleep(time.Second)
	}
}
