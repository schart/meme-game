package controllers

import (
	"net/http"
	db_queries "shared-library/database/queries/queries-meme"
	utils "shared-library/utils"
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
	count := r.URL.Query().Get("count")
	if count == "" {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Count is needed for get the  records of texts")
	}

	rows, err := db_queries.PhotoGetByCount(count)
	if err != nil {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Error: "+err.Error())
		return
	}
	utils.HandleSuccess(w, rows)
	return
}
