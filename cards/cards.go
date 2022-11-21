package cards

/**
 * Defines what makes up a card including value, suit, and symbol. Including some helper methods
 * @author Kyle Shultz
 */
import (
	"fmt"
)

type CardValue int

// Define constants for creating a deck
const CardValues int = 14
const Suits int = 4

const (
	CARD_ONE CardValue = iota
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
	SUIT_SPADES   = "Spades"
	SUIT_HEARTS   = "Hearts"
	SUIT_DIAMONDS = "Diamonds"
	SUIT_CLUBS    = "Clubs"
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
		break
	case 1:
		suit = SUIT_HEARTS
		break
	case 2:
		suit = SUIT_DIAMONDS
		break
	case 3:
		suit = SUIT_CLUBS
		break
	}
	return suit
}

// Formats a playing card
func (card Card) String() string {
	var symbol string

	switch CardValue(card.Symbol) {
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
