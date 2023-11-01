package controllers

import (
	"net/http"
	queries_game "shared-library/database/queries/queries-game"
	"shared-library/interceptors"
	utils "shared-library/utils"
)

func GetAllRoomsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodGet, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	// Check access key
	presenceOfKey := interceptors.AccessKeyCheck(w, r)
	if presenceOfKey == false {
		utils.HandleError(w, http.StatusBadRequest, "Error: Required access key, check your 'access key'")
		return
	}

	room := queries_game.GetAllRoom()
	utils.HandleSuccess(w, map[string]interface{}{"room": room})
	return
}
