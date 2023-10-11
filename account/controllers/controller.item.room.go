package controllers

import (
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	utils "shared-library/utils"
)

func GetRoomAccountController(w http.ResponseWriter, r *http.Request) {
	// -> Start http conn
	// This http connection 'll may be change with ws conn
	w.Header().Set("Content-Type", "text/plain")

	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodGet, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}
	// -> End http conn

	// Check session
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

	accountId := float64(1)

	rooms := queries_account.GetRoomOfAccount(accountId)
	utils.HandleSuccess(w, rooms)
	return
}
