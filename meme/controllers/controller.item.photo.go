package controllers

import (
	"net/http"
	db_queries "shared-library/database/queries/queries-meme"
	utils "shared-library/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func PhotoItemsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	checkMethod := utils.HttpMethodSet(http.MethodGet, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	// Load enviroments
	utils.EnvLoader()
	// Check count param in url parameters
	count := mux.Vars(r)["count"]
	if count == "" {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Count is needed for get the  records of texts")
	}

	// Parse the `count` variable as an integer.
	countInt, err := strconv.Atoi(count)

	// Check if the `count` variable could be parsed as an integer.
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Count must be an integer")
		return
	}

	rows, err := db_queries.PhotoGetByCount(countInt)
	if err != nil {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Error: "+err.Error())
		return
	}

	utils.HandleSuccess(w, rows)
	return
}
