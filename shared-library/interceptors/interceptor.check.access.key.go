package interceptors

import (
	"net/http"
	"os"
	"shared-library/utils"
)

func AccessKeyCheck(w http.ResponseWriter, r *http.Request) bool {
	utils.EnvLoader()

	var key string = r.Header.Get("ACCESS_KEY")

	if os.Getenv("ACCESS_KEY") == "" {
		return false
	}

	if key == os.Getenv("ACCESS_KEY") {
		return true
	}

	return false

}

func TokenCheck(w http.ResponseWriter, r *http.Request) bool {
	utils.EnvLoader()

	// Get token
	token, err := r.Cookie("token")
	if err != nil {
		return false
	}

	// If cookie value is empty then turn the false
	if token.Value == "" {
		return false
	}
	// Otherwise
	return true
}
