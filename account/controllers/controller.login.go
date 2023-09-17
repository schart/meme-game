package controllers

import (
	"fmt"
	"net/http"
	accountQueries "shared-library/database/queries/queries-account"
	accountTypes "shared-library/types/account"
	"shared-library/utils"
)

func AccountLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}
	fmt.Println("Test cookie: ", r.Cookies())

	// Check token
	// here

	// Get parsed body
	formData := utils.ParsedBodyGet(w, r)

	// Check body
	err := utils.ParameterChecker(formData, accountTypes.AccountLogin{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Check account
	res, err := accountQueries.AccountOneGet(formData)
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	if res == nil {
		utils.HandleError(w, http.StatusUnauthorized, "Account not founded")
		return
	}

	utils.HandleSuccess(w)
}
