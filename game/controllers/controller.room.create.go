package controllers

import (
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	queries_game "shared-library/database/queries/queries-game"
	utils "shared-library/utils"

	"github.com/gofrs/uuid"
)

func RoomCreateController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

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

	status := queries_account.AccountAvaliableViaId(accountId)
	if status == false {
		utils.HandleError(w, http.StatusBadRequest, "Account could not found")
		return
	}

	/*
		@ Checked account session and presence of wihch
		@ Now, we checking account joined a room?
	*/

	status = queries_account.AccountHaveTheRoom(accountId)
	if status == true {
		utils.HandleError(w, http.StatusBadRequest, "You already joined a room")
		return
	}

	/*
		@ Finally, we creating room link as id and turn it
	*/

	id, _ := uuid.NewV4()
	err := queries_game.RoomInsert(accountId, id)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSuccess(w, map[string]interface{}{"id": id.String()})
	return
}
