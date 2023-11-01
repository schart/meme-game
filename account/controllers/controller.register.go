package controllers

import (
	"net/http"
	"shared-library/database/queries/queries-account"
	accountQueries "shared-library/database/queries/queries-account"
	myJwt "shared-library/jwt"
	accountTypes "shared-library/types/account"
	"shared-library/utils"
)

func AccountRegisterController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	formData := utils.ParsedBodyGet(w, r)

	err := utils.ParameterChecker(formData, accountTypes.AccountRegister{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	/*
		@ We checked body and parsed it
		@ We must do check account, because db will does  declare account duplicate in otherwise

	*/

	status := queries.IsAccountThereUsername(formData.Get("Username"))
	if status == true {
		utils.HandleError(w, http.StatusExpectationFailed, "Username is already using")
		return
	}

	/*
		@ Finally we can create account, create token with its informations
		@ And publish token as cookie!
	*/

	id, err := accountQueries.AccountInsert(formData.Get("Username"), formData.Get("Password"))
	if err != nil {
		utils.HandleError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	token, err := myJwt.CreateJWT(id)
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

	// Publish the token
	http.SetCookie(w, cookie)
	utils.HandleSuccess(w, map[string]interface{}{})
	return
}
