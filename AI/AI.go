package main

import (
	"main/cards"
)

type AI struct {
	Total int
	cards []cards.Card
}

func newAI() *AI {
	return &AI{
		Total: 0,
		cards: make([]cards.Card, 3),
	}
}

func keepHitting(ai *AI) {
	for ai.Total < 17 {
		ai.cards = append(ai.cards, cards.NewCard(cards.CARD_ONE, '\u2663'))
		ai.Total += ai.cards[len(ai.cards)-1].Value

	}
}
