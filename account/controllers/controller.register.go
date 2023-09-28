package controllers

import (
	"net/http"
	accountQueries "shared-library/database/queries/queries-account"
	myJwt "shared-library/jwt"
	accountTypes "shared-library/types/account"
	"shared-library/utils"
	"time"
)

func AccountRegisterController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	// get parsed body
	formData := utils.ParsedBodyGet(w, r)

	// check body
	err := utils.ParameterChecker(formData, accountTypes.AccountRegister{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create account
	err = accountQueries.AccountInsert(formData.Get("Username"), formData.Get("Password"))
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
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
