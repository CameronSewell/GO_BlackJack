package Game

import (
	"log"
	"main/cards"
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

	return dealer
}

// Deal two cards to the given hand to start the game
func DealStartingHand(dealer Dealer, hand cards.Hand) cards.Hand {
	for i := 0; i < 2; i++ {
		hand = cards.AddCard(DealCard(dealer), hand)
	}
	return hand
}

// Dealer deals a card from the remaining deck
func DealCard(dealer Dealer) cards.Card {
	card, err := cards.HitDeck(dealer.deck)
	if err != nil {
		log.Fatal(err)
	}
	return *card
}

// Makes the dealer hit the deck until their total
// reaches or breaks 17. When that happens, the dealer stops
// Taking hits
func DealerPlay(dealer Dealer) {
	for cards.GetHandTotal(dealer.Hand) < 17 {
		dealer.Hand = cards.AddCard(DealCard(dealer), dealer.Hand)
	}
}
