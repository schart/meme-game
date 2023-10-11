package controllers

import (
	"fmt"
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	queries_game "shared-library/database/queries/queries-game"
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

	_ = queries_account.GetRoomOfAccount(accountId)

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
