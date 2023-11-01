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
	// -> Start http conn
	// This http connection 'll may be change with ws conn
	w.Header().Set("Content-Type", "text/plain")

	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}
	// -> End http conn

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

	// Parse body
	formData := utils.ParsedBodyGet(w, r)

	// Check body
	err := utils.ParameterChecker(formData, types_game.JoinRoom{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	roomLink := formData.Get("RoomLink")

	/*// Get account id in jwt
	claims, err := myjwt.DecodeJWT(token)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err.Error())
	}

	accountId := claims["accountId"].(float64)
	fmt.Println("Acc id: ", accountId)
	*/

	accountId := float64(2)

	// Presence of account
	status := queries_account.AccountAvaliableViaId(accountId)
	if status == false {
		utils.HandleError(w, http.StatusBadRequest, "Account could not found")
		return
	}

	// Presence of room
	status = queries_game.RoomAvailable(roomLink)
	if status == false {
		utils.HandleError(w, http.StatusBadRequest, "Room not found: "+roomLink)
		return
	}

	// Check user have a room or joined a room
	status = queries_account.AccountHaveTheRoom(accountId)
	if status == true {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, "Just one room, you can create/join")
		return
	}

	room := queries_game.GetRoomByLink(roomLink)
	roomId := room["id"].(int)

	// Join room
	err = queries_game.RoomJoin(accountId, roomId, false)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSuccess(w, map[string]interface{}{})
	return
}
