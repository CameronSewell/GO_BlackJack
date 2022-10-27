package main

import (
	"main/message"
	"net/http"

	"golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
	var data message.T
	websocket.JSON.Receive(ws, &data)
	data.Count += 5
	websocket.JSON.Send(ws, data)
}

func main() {
	http.Handle("/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("Listen And Serve: " + err.Error())
	}
}
