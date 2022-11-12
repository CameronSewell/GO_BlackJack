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
	bet()
}

type superPLayer struct
{
	hand Card
	//money dollars
}

func stand()
{

}
func hit()
{

}
func quit()
{

}
func bet()