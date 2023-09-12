package utils

import (
	"net/http"
)

func HttpMethodSet(expectMethod string, r *http.Request) bool {
	if r.Method != expectMethod {
		return false
	}
	return true
}
