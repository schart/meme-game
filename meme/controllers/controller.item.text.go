package controllers

import (
	"net/http"
	db_queries "shared-library/database/queries/queries-meme"
	utils "shared-library/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func TextItemsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	checkMethod := utils.HttpMethodSet(http.MethodGet, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	utils.EnvLoader()

	count := mux.Vars(r)["count"]
	if count == "" {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Count is needed for get the  records of texts")
	}

	countInt, err := strconv.Atoi(count)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Count must be an integer")
		return
	}

	/*

	  @ We taken params and converted needed data type to use
	  @ Finally, taken texts in the database according to count

	*/

	texts := db_queries.GetText(countInt)

	utils.HandleSuccess(w, map[string]interface{}{"texts": texts})
	return
}
