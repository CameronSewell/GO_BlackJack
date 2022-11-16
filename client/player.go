package client

import (
	"fmt"
	"log"
	"main/message"
	"main/cards"
	"main/Dealer"
	"golang.org/x/net/websocket"
	"os"
)

type Player interface
{
	newPlayer()
	hit()
	stand()
	quit()
	bet(int)

	var Total int
	var hand []cards.cards
	var money float32
	var buyIn float32
}

func newPlayer() Player {

	player := Player
	{
		Total: 0,
		hand:  make([]cards.Card, 3),
		money: 10.0,
		buyIn 1.5,
	}

	for i := 0; i < 2; i++
	{
		player.hand = append(Dealer.DealCard())
	}

	return p
}

func stand()
{
	return player.hand
}
func hit()
{
	var Card card = cards.HitDeck(dealer.deck)
	player.Total += card.value
	return player.hand = append(card)
}
func quit()
{
	os.Exit(0)
}
func bet(int amount)
{
	player.money = player.money - amount
	//add amount to the pot
}
func getTotal()
{
	return player.total
}
func split()
{
	//
}
func double()
{
	player.bet(player.buyIn)
}