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

	status := utils.CheckWebsocketConnection(r)
	if status == false {
		utils.HandleErrorWS(conn, "This connection is not web socket")
		return
	}

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

	status = queries_account.AccountAvaliableViaId(accountId)
	if status == false {
		utils.HandleErrorWS(conn, "Account could not found")
		return
	}

	/*

		@ Checked account have a session and account is available?
		@ Now, we check room available and account is owner of the room?

	*/

	room_link := mux.Vars(r)["room_link"]

	status = queries_game.RoomAvailable(room_link)
	if status == false {
		utils.HandleErrorWS(conn, "Room not found: "+room_link)
		return
	}

	status = queries_account.AccountOwnerTheRoom(int(accountId))

	if status == false {
		utils.HandleErrorWS(conn, "Unauthorized")
		return
	}

	/*

		@ Now, we getting room via link for check account count in the room

	*/

	room := queries_game.GetRoomByLink(room_link)
	roomId := room["id"].(int)

	howAcc := queries_account.AccountCountInRoom(roomId)
	if howAcc == 0 {
		utils.HandleErrorWS(conn, "Min require 4 account")
		return
	}

	/*

		@ Now, we getting accounts in the room.
		@ Before, Create account cache for each account in the room.

	*/

	/*

		@ Creating Room Cache

	*/

	var accounts = queries_account.GetAccountsInRoom(roomId)
	for i := 0; i < len(accounts); i++ {
		err := service_redis.CreateAccountCacheService(accounts[i])
		if err != nil {
			utils.HandleErrorWS(conn, err.Error())
			return
		}
	}

	/*
		@ Creating Room Cache
	*/

	err := service_redis.CreateRoomCacheService(float64(roomId), room_link, howAcc, int(accountId))
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	/*
		@ Creating Round Cache
	*/

	err = service_redis.CreateRoundCacheService(float64(roomId), room_link)
	if err != nil {
		utils.HandleErrorWS(conn, err.Error())
		return
	}

	/*

		@ Finally, we started a game session
		@ Because those specifier of the started game

	*/

	utils.HandleSuccessWS(conn, map[string]interface{}{})
	return
}
