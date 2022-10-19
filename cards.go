package main

/**
 * Defines what makes up a card including value, suit, and symbol. Including some helper methods
 * @author Kyle Shultz
 */
import (
	"fmt"
)

const (
	CARD_ONE = iota
	CARD_TWO
	CARD_THREE
	CARD_FOUR
	CARD_FIVE
	CARD_SIX
	CARD_SEVEN
	CARD_EIGHT
	CARD_NINE
	CARD_TEN
	CARD_JACK
	CARD_QUEEN
	CARD_KING
	CARD_ACE //TO-DO: We will need a way to address alternate values for Aces
)

/* These are unicode values that represent the suits of a typical card deck */
const (
	SUIT_SPADES   = '\u2660'
	SUIT_HEARTS   = '\u2665'
	SUIT_DIAMONDS = '\u2666'
	SUIT_CLUBS    = '\u2663'
)

// Represents a single card in a hand.
type Card struct {
	// The symbol of the card.
	Symbol int

	// The suit of the card.
	Suit rune

	// The value of the card.
	Value int
}

// Formats a playing card
func (card Card) String() string {
	var symbol string

	switch card.Symbol {
	default:
		symbol = fmt.Sprintf("%d", card.Value)
	case CARD_JACK:
		symbol = "J"
	case CARD_QUEEN:
		symbol = "Q"
	case CARD_KING:
		symbol = "K"
	case CARD_ACE:
		symbol = "A"
	}

	return fmt.Sprintf("%s%c", symbol, card.Suit)
}

// Creates a new card with the given symbol and suit.
func NewCard(symbol int, suit rune) Card {
	card := Card{}
	card.Symbol = symbol
	card.Suit = suit

	var value int

	// Determine the value of the card
	switch symbol {
	default:
		value = symbol + 1
	case CARD_JACK, CARD_QUEEN, CARD_KING:
		value = 10
	case CARD_ACE:
		value = 11
	}

	card.Value = value

	return card
}
