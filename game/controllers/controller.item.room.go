package controllers

import (
	"net/http"
	queries_game "shared-library/database/queries/queries-game"
	"shared-library/interceptors"
	utils "shared-library/utils"
)

func GetAllRoomsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	checkMethod := utils.HttpMethodSet(http.MethodGet, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	presenceOfKey := interceptors.AccessKeyCheck(w, r)
	if presenceOfKey == false {
		utils.HandleError(w, http.StatusBadRequest, "Error: Required access key, check your 'access key'")
		return
	}

	/*
		@ Checked access key.
		@ Finally we getting all rooms and turn these as reponse
	*/

	room := queries_game.GetAllRoom()
	utils.HandleSuccess(w, map[string]interface{}{"room": room})
	return
}
