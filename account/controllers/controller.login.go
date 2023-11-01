package controllers

import (
	"net/http"
	accountQueries "shared-library/database/queries/queries-account"
	"shared-library/interceptors"
	myjwt "shared-library/jwt"
	accountTypes "shared-library/types/account"
	"shared-library/utils"
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
	id, err := accountQueries.IsAccountValidated(formData)
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	// Account not found
	if id == 0 {
		utils.HandleError(w, http.StatusUnauthorized, "Account not founded")
		return
	}

	// Create token
	token, err := myjwt.CreateJWT(id)
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	// Create cookie
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   315000101, //time.Now().Add(24 * time.Hour), // change the time,
		HttpOnly: true,
		Secure:   true,
	}

	// Publish the token
	http.SetCookie(w, cookie)
	utils.HandleSuccess(w, map[string]interface{}{})
	return
}
