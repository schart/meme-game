package controllers

import (
	"fmt"
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	queries_game "shared-library/database/queries/queries-game"
	types_game "shared-library/types/game"
	utils "shared-library/utils"

	"github.com/gofrs/uuid"
)

func RoomCreateController(w http.ResponseWriter, r *http.Request) {
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

	/*
		// Get token
		token, err := r.Cookie("token")
		if err != nil {
			utils.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Check logged status
		statusToken := interceptor.ValidateJWT(token)
		if statusToken == false {
			utils.HandleError(w, http.StatusUnauthorized, "Required login")
			return
		}

		// Get account id in jwt
		claims, err := myjwt.DecodeJWT(token)
		if err != nil {
			utils.HandleError(w, http.StatusUnauthorized, err.Error())
		}

		accountId := claims["accountId"].(float64) // assertion to float64
	*/

	accountId := float64(1) //  id 1 name heja

	// Presence of Account
	account, err := queries_account.IsThereAccount(accountId)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Owner could not found
	if (account == nil) == true {
		utils.HandleError(w, http.StatusBadRequest, "Account could not found")
		return
	}

	// Check user have a room or joined a room
	statusRoom, err := queries_account.IsHaveARoomAccount(accountId)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// If joined a room
	if statusRoom == true {
		utils.HandleError(w, http.StatusBadRequest, "You already joined a room")
		return
	}

	// Create id for link
	id, _ := uuid.NewV4()

	// Create room
	err = queries_game.RoomInsert(accountId, id)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSuccess(w, []string{})
	return
}

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
	account, err := queries_account.IsThereAccount(accountId)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// account could not found
	if account == nil {
		utils.HandleError(w, http.StatusBadRequest, "Account could not found")
		return
	}

	// Presence of room
	room, err := queries_game.IsThereRoom(formData.Get("RoomId"))
	if err != nil {
		utils.HandleError(w, http.StatusBadGateway, err.Error())
		return
	}

	// If not any record
	if room == false {
		utils.HandleError(w, http.StatusBadGateway, "Room not found: "+formData.Get("RoomId"))
		return
	}

	// Check user have a room or joined a room
	statusRoom, err := queries_account.IsHaveARoomAccount(accountId)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// If joined a room
	if statusRoom == true {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, "Just one room, you can create/join")
		return
	}

	// Join room
	err = queries_game.RoomJoin(accountId, formData.Get("RoomId"), false)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSuccess(w, []string{})
	return
}
