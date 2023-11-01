package controllers

import (
	"net/http"
	"shared-library/interceptors"
	"shared-library/utils"
)

func AccountLogoutController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	checkMethod := utils.HttpMethodSet(http.MethodGet, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	statusToken := interceptors.TokenCheck(w, r)
	if statusToken != true {
		utils.HandleError(w, http.StatusUnauthorized, "You are already have not session")
		return
	}

	/*
		@ Checked token for token there for do logout
		@ Now, we doing logout by deleting the cookie of the in session
	*/

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   315000101, //time.Now().Add(24 * time.Hour), // change the time,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	utils.HandleSuccess(w, map[string]interface{}{})
	return
}
