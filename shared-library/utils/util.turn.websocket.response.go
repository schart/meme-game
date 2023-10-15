package utils

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleErrorWS(conn *websocket.Conn, message string) {
	response := NewResponse(false, http.StatusExpectationFailed, message, []string{})
	conn.WriteJSON(response)

}

func HandleSuccessWS(conn *websocket.Conn, params []string) {
	response := NewResponse(true, http.StatusOK, "Successfuly proccess", params)
	conn.WriteJSON(response)
}
