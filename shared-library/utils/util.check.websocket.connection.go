package utils

import (
	"fmt"
	"net/http"
)

func CheckWebsocketConnection(r *http.Request) bool {
	// Check if the headers indicate that this is a WebSocket connection request.
	if r.Header.Get("Upgrade") == "websocket" &&
		r.Header.Get("Connection") == "Upgrade" &&
		r.Header.Get("Sec-WebSocket-Key") != "" {
		fmt.Println("This is a WebSocket connection request.")
		return true
	} else {
		return false
	}
}
