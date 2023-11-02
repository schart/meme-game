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

	statusToken := interceptors.TokenCheck(w, r)
	if statusToken == true {
		utils.HandleError(w, http.StatusUnauthorized, "You already have a session")
		return
	}

	formData := utils.ParsedBodyGet(w, r)

	err := utils.ParameterChecker(formData, accountTypes.AccountLogin{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	/*
		@ Checked body, session and authenticated parameter
		@ Now, we check if the account is verified or unverified and finally we create and publish the jwt
	*/

	status := accountQueries.AccountVerified(formData)
	if status == false {
		utils.HandleError(w, http.StatusUnauthorized, "Account not founded")
		return
	}

	account := accountQueries.GetAccountViaUsername(formData.Get("Username"))
	id := account["id"].(int)

	token, err := myjwt.CreateJWT(id)
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   315000101, //time.Now().Add(24 * time.Hour), // change the time,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	utils.HandleSuccess(w, map[string]interface{}{})
	return
}
