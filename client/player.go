package client

import (
	"fmt"
	"log"
	"main/message"
	"main/cards"
	"main/Dealer"
	"cards.go"
	"golang.org/x/net/websocket"
)

type player interface
{
	hit()
	stand()
	quit()
	bet(int)

	var Total int
	var hand []cards.cards
	var money float32
}

func stand()
{
	var data message.T
	data.Msg = "call stand()"
	if err := websocket.JSON.Send(ws, data); err != nil {
		log.Fatal(err)
	}
	if err = websocket.JSON.Receive(ws, &data); err != nil {
		log.Fatal(err)
	}
}
func hit()
{
	var data message.T
	data.Msg = "dealer.drawCard()"
	if err := websocket.JSON.Send(ws, data); err != nil {
		log.Fatal(err)
	}
	if err = websocket.JSON.Receive(ws, &data); err != nil {
		log.Fatal(err)
	}
}
func quit()
{

}
func bet(int amount)
{
	var data message.T
	data.Msg = "call bet(amount)"
	if err := websocket.JSON.Send(ws, data); err != nil {
		log.Fatal(err)
	}
	if err = websocket.JSON.Receive(ws, &data); err != nil {
		log.Fatal(err)
	}
}