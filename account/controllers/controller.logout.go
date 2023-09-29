package controllers

import (
	"net/http"
	"shared-library/interceptors"
	"shared-library/utils"
)

func AccountLogoutController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Check method
	checkMethod := utils.HttpMethodSet(http.MethodGet, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	// Check token, You are must be logged in.
	statusToken := interceptors.TokenCheck(w, r)
	if statusToken != true {
		utils.HandleError(w, http.StatusUnauthorized, "You are already have not session")
		return
	}

	// Create cookie
	cookie := &http.Cookie{
		Name:  "token",
		Value: "",
		// Expires: <-time.Tick(0), // change the time
	}

	// Publish the token
	http.SetCookie(w, cookie)
	utils.HandleSuccess(w, []string{})
	return
}
