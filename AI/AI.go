package main

import (
	"log"
	"main/cards"
)

type AI struct {
	Total int
	cards []cards.Card
}

func NewAI() AI {
	return AI{
		Total: 0,
		cards: make([]cards.Card, 3),
	}
}

func Play(ai *AI, d *cards.Deck) {
	for ai.Total < 17 {
		card, err := cards.HitDeck(*d)
		if err != nil {
			log.Fatal(err)
		}
		ai.cards = append(ai.cards, *card)
		ai.Total += card.Value

	}
}
