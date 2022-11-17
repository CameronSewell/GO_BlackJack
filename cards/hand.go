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
func GetHandCount(h Hand) int {
	return len(h.cards)
}

// Get the hand total
func GetHandTotal(h Hand) int {
	return h.total
}

// Return true if the given hand is a blackjack
func IsBlackjack(h Hand) bool {
	return h.total == Blackjack
}

// Return true if the given hand's total is over 21
func IsBust(h Hand) bool {
	return h.total > Blackjack
}

// Get first card of the hand
func GetFirstCard(h Hand) Card {
	return h.cards[0]
}

// Add a card to the given hand
func AddCard(card Card, h Hand) Hand {
	h.cards = append(h.cards, card)
	h.total += card.Value
	return h
}

// Get all the cards of the hand (ONLY CALL AT END OF GAME)
func GetCards(h Hand) []Card {
	return h.cards
}
