package controllers

import (
	"fmt"
	service_redis "game/services"
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	queries_game "shared-library/database/queries/queries-game"
	utils "shared-library/utils"
	websocket_connect "shared-library/websocket-connection"

	"github.com/gorilla/mux"
)

func StartGameController(w http.ResponseWriter, r *http.Request) {
	conn := websocket_connect.Connect(w, r)

	// check connection is websocket conn?
	status := utils.CheckWebsocketConnection(r)
	if status == false {
		utils.HandleErrorWS(conn, "This connection is not web socket")
		return
	}

	room_link := mux.Vars(r)["room_link"]

	/*
		// Get token
		token, err := r.Cookie("token")
		if err != nil {
			utils.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("token room: ", token)

		// Check jwt
		statusToken := interceptor.ValidateJWT(token)
		if statusToken == false {
			utils.HandleError(w, http.StatusUnauthorized, "Required login")
			return
		}
	*/

	accountId := float64(1)

	// Presence of account
	account, err := queries_account.IsThereAccount(accountId)
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	// account could not found
	if account == nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	// Presence of room
	roomId, err := queries_game.IsThereRoom(room_link)
	fmt.Println(roomId, err)
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	// If not any record
	if roomId == 0 {
		utils.HandleErrorWS(conn, "Room not found: "+room_link)
		return
	}

	// If user owner of the room
	status, err = queries_account.IsAccountOwnerRoom(int(accountId))
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	// if account is does not owner of the room
	if status == false {
		utils.HandleErrorWS(conn, "Unauthorized")
		return
	}

	// How much there user in the room?
	howAcc, err := queries_account.CheckMinAccountInRoom(roomId)
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	if howAcc == 0 {
		utils.HandleErrorWS(conn, "Min require 4 account")
		return
	}

	err = service_redis.CacheInitForGame(accountId)
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	// Start game
	utils.HandleSuccessWS(conn, []string{})
	return
}

/*
func NextRoundController(w http.ResponseWriter, r *http.Request) {

}*/
