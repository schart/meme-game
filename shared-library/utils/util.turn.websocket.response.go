package utils

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleErrorWS(conn *websocket.Conn, message string) {
	response := NewResponse(false, http.StatusExpectationFailed, message, map[string]interface{}{})
	conn.WriteJSON(response)

}

func HandleSuccessWS(conn *websocket.Conn, params map[string]interface{}) {
	response := NewResponse(true, http.StatusOK, "Successfuly proccess", params)
	conn.WriteJSON(response)
}
