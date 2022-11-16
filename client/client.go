package client

import (
	"fmt"
	"log"
	"main/message"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost"
	url := "ws://localhost:12345/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	var data message.T
	data.Msg = "Hello World"
	data.Count = 0
	if err := websocket.JSON.Send(ws, data); err != nil {
		log.Fatal(err)
	}
	if err = websocket.JSON.Receive(ws, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s %d.\n", data.Msg, data.Count)
}
