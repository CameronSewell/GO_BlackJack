package cards

/**
 * Defines what makes up a card including value, suit, and symbol. Including some helper methods
 * @author Kyle Shultz
 */
import (
	"fmt"
	"strconv"
)

type CardValue int

// Define constants for creating a deck
const CardValues int = 13
const Suits int = 4

const (
	CARD_ACE CardValue = iota
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
)

/* These are unicode values that represent the suits of a typical card deck */
const (
	SUIT_SPADES   = "spades"
	SUIT_HEARTS   = "hearts"
	SUIT_DIAMONDS = "diamonds"
	SUIT_CLUBS    = "clubs"
)

// Represents a single card in a hand.
type Card struct {
	// The symbol of the card.
	Symbol int

	// The suit of the card.
	Suit string

	// The value of the card.
	Value int
}

func getSuit(Suit int) string {
	var suit string
	switch Suit {
	default:
	case 0:
		suit = SUIT_SPADES

	case 1:
		suit = SUIT_HEARTS

	case 2:
		suit = SUIT_DIAMONDS

	case 3:
		suit = SUIT_CLUBS

	}
	return suit
}

// Formats a playing card
func (card Card) String() string {
	var symbol string

	switch CardValue(card.Symbol) {
	default:
		symbol = strconv.Itoa(card.Value)
	case CARD_JACK:
		symbol = "jack"
	case CARD_QUEEN:
		symbol = "queen"
	case CARD_KING:
		symbol = "king"
	case CARD_ACE:
		symbol = "ace"
	}

	return fmt.Sprintf("%s_%s", card.Suit, symbol)
}

// Creates a new card with the given symbol and suit.
func NewCard(symbol int, suit string) Card {
	card := Card{}
	card.Symbol = symbol
	card.Suit = suit

	var value int

	// Determine the value of the card
	switch CardValue(symbol) {
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
