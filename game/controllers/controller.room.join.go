package controllers

import (
	"fmt"
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	queries_game "shared-library/database/queries/queries-game"
	types_game "shared-library/types/game"
	utils "shared-library/utils"
)

func JoinRoomController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	/*// Get token
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

	/*// Get account id in jwt
	claims, err := myjwt.DecodeJWT(token)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err.Error())
	}

	accountId := claims["accountId"].(float64)
	fmt.Println("Acc id: ", accountId)
	*/

	accountId := float64(2)

	status := queries_account.AccountAvaliableViaId(accountId)
	if status == false {
		utils.HandleError(w, http.StatusBadRequest, "Account could not found")
		return
	}

	/*
		@ Checked body, session and authenticated parameter and presence of account
		@ Now, we parse the body and check the parameter, we also check the available room, has the account joined a room?
	*/

	formData := utils.ParsedBodyGet(w, r)

	err := utils.ParameterChecker(formData, types_game.JoinRoom{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	roomLink := formData.Get("RoomLink")

	status = queries_game.RoomAvailable(roomLink)
	if status == false {
		utils.HandleError(w, http.StatusBadRequest, "Room not found: "+roomLink)
		return
	}

	status = queries_account.AccountHaveTheRoom(accountId)
	if status == true {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, "Just one room, you can create/join")
		return
	}

	/*

		@ Finally, join room

	*/

	room := queries_game.GetRoomByLink(roomLink)
	roomId := room["id"].(int)

	err = queries_game.RoomJoin(accountId, roomId, false)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSuccess(w, map[string]interface{}{})
	return
}
