package controllers

import (
	"net/http"
	accountQueries "shared-library/database/queries/queries-account"
	myJwt "shared-library/jwt"
	accountTypes "shared-library/types/account"
	"shared-library/utils"
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
	id, err := accountQueries.AccountInsert(formData.Get("Username"), formData.Get("Password"))
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	// Create token
	token, err := myJwt.CreateJWT(id)
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
	utils.HandleSuccess(w, []string{})
	return
}
