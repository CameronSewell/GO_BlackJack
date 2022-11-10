package dealer

import (
	"log"
	"main/cards"
)

// Defines a dealer with a total, their
// hand, and the deck to use for the game
type Dealer struct {
	Total int
	hand  []cards.Card
	deck  cards.Deck
}

// Tells the dealer to draw a card from the deck
// and put it in their own hand
func drawCard(dealer Dealer) {
	card, err := cards.HitDeck(dealer.deck)
	if err != nil {
		log.Fatal(err)
	}
	dealer.hand = append(dealer.hand, *card)
	dealer.Total += card.Value
}

// Creates a new dealer struct and gives them
// The two starting cards of their hand
func NewDealer() Dealer {

	dealer := Dealer{
		Total: 0,
		hand:  make([]cards.Card, 3),
		deck:  cards.NewDeck(),
	}

	for i := 0; i < 2; i++ {
		drawCard(dealer)
	}

	return dealer
}

// Deal a card from the deck
func DealCard(dealer Dealer) (*cards.Card, error) {
	return cards.HitDeck(dealer.deck)
}

// Get the number of cards in the current hand
func GetDealerHandCount(dealer Dealer) int {
	return len(dealer.hand)
}

// Get the first card in the dealer's hand (meant to
// be the only visible hand)
func GetDealerVisibleCard(dealer Dealer) cards.Card {
	return dealer.hand[0]
}

// Get the total value of the dealer's hand
func GetDealerHandTotal(dealer Dealer) int {
	return dealer.Total
}

// Makes the dealer hit the deck until their total
// reaches or breaks 17. When that happens, the dealer stops
// Taking hits
func DealerPlay(dealer Dealer) {
	for dealer.Total < 17 {
		drawCard(dealer)
	}
}
