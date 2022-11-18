package cards

var Blackjack int = 21

type Hand struct {
	total int
	cards []Card
}

// Create an empty hand
func NewHand() Hand {
	h := Hand{
		total: 0,
		cards: make([]Card, 1),
	}
	return h
}

// Get the count of the hand
func (h *Hand) GetHandCount() int {
	return len(h.cards)
}

// Get the hand total
func (h *Hand) GetHandTotal() int {
	return h.total
}

// Return true if the given hand is a blackjack
func (h *Hand) IsBlackjack() bool {
	return h.total == Blackjack
}

// Return true if the given hand's total is over 21
func (h *Hand) IsBust() bool {
	return h.total > Blackjack
}

// Get first card of the hand
func (h *Hand) GetFirstCard() Card {
	return h.cards[0]
}

// Add a card to the given hand
func (h *Hand) AddCard(card Card) {
	h.cards = append(h.cards, card)
	h.total += card.Value
}

// Get all the cards of the hand (ONLY CALL AT END OF GAME)
func (h *Hand) GetCards() []Card {
	return h.cards
}
