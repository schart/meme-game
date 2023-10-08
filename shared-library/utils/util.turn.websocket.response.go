package utils

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleErrorWS(conn *websocket.Conn, method any, message string) {

	response := NewResponse(false, http.StatusFailedDependency, message, []string{})
	conn.WriteJSON(response)
}

func HandleSuccessWS(conn *websocket.Conn, method any) {
	response := NewResponse(true, http.StatusOK, "Successfuly proccess", []string{})
	conn.WriteJSON(response)
}
