package services

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Send message.
func SendMessage(message string, conn *websocket.Conn) error {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Println(err)
		panic("Error, whe write message")

	}
	return nil
}

// Receive message.
func ReceiveMessage(conn *websocket.Conn) string {
	_, p, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Could not read message:", err)
		return ""
	}

	return string(p)
}
