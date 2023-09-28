package controllers

import (
	"net/http"
	accountQueries "shared-library/database/queries/queries-account"
	"shared-library/interceptors"
	myJwt "shared-library/jwt"
	accountTypes "shared-library/types/account"
	"shared-library/utils"
	"time"
)

func AccountLoginController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	// Check token
	statusToken := interceptors.TokenCheck(w, r)
	if statusToken == true {
		utils.HandleError(w, http.StatusUnauthorized, "You already have a session")
		return
	}

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

	// Create token
	token, err := myJwt.TokenCreate(formData.Get("Username"))
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	// Create cookie
	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour), // change the time
	}

	// Publish the token
	http.SetCookie(w, cookie)
	utils.HandleSuccess(w, []string{})
}
