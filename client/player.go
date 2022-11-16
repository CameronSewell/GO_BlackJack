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
	log.Println("stand func called")
	return player.hand
}
func hit()
{
	log.Println("hit func called")
	var Card card = cards.HitDeck(dealer.deck)
	player.Total += card.value
	return player.hand = append(card)
}
func quit()
{
	log.Println("quit func called")
	os.Exit(0)
}
func bet(int amount)
{
	log.Println("bet func called")
	player.money = player.money - amount
	//add amount to the pot
}
func getTotal()
{
	return player.total
}
func split()
{
	return log.Println("spit func called")
	//
}
func double()
{
	log.Println("double func called")
	return player.bet(player.buyIn)
}