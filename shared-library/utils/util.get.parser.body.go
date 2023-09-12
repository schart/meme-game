package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ParsedBodyGet(w http.ResponseWriter, r *http.Request) url.Values {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("Error: when read body %s", err)
	}

	formData, err := url.ParseQuery(string(body))
	return formData
}
