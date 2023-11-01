package controllers

import (
	service_redis "game/services"
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	queries_game "shared-library/database/queries/queries-game"
	utils "shared-library/utils"
	websocket_connect "shared-library/websocket-connection"

	"github.com/gorilla/mux"
)

func StartPlayController(w http.ResponseWriter, r *http.Request) {
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
	status = queries_account.AccountAvaliableViaId(accountId)
	if status == false {
		utils.HandleErrorWS(conn, "Account could not found")
		return
	}

	// Presence of room
	status = queries_game.RoomAvailable(room_link)
	if status == false {
		utils.HandleErrorWS(conn, "Room not found: "+room_link)
		return
	}

	// If user owner of the room
	status = queries_account.AccountOwnerTheRoom(int(accountId))

	// if account is does not owner of the room
	if status == false {
		utils.HandleErrorWS(conn, "Unauthorized")
		return
	}

	// Get room via link
	room := queries_game.GetRoomByLink(room_link)
	roomId := room["id"].(int)

	// How much there user in the room?
	howAcc := queries_account.AccountCountInRoom(roomId)

	if howAcc == 0 {
		utils.HandleErrorWS(conn, "Min require 4 account")
		return
	}

	// Get accounts in the room
	var accounts = queries_account.GetAccountsInRoom(roomId)

	// Create a  account game cache
	for i := 0; i < len(accounts); i++ {
		err := service_redis.CreateAccountCacheService(accounts[i])
		if err != nil {
			utils.HandleErrorWS(conn, err.Error())
			return
		}
	}

	// Create a room game cache
	err := service_redis.CreateRoomCacheService(float64(roomId), room_link, howAcc, int(accountId))
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	// Create a round game cache
	err = service_redis.CreateRoundCacheService(float64(roomId), room_link)
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	// Start game
	utils.HandleSuccessWS(conn, map[string]interface{}{})
	return
}
