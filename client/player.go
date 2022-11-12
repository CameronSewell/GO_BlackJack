package main

import (
	"fmt"
	"log"
	"main/message"
	"cards.go"
	"golang.org/x/net/websocket"
)

type player interface
{
	hit()
	stand()
	quit()
	bet(int)
}

type superPLayer struct
{
	hand Card
	//money dollars
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
	data.Msg = "call hit()"
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