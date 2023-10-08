package interceptors

import (
	"fmt"
	"net/http"
	"os"
	utils "shared-library/utils"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(cookie *http.Cookie) bool {
	// ENV Loader
	utils.EnvLoader()

	// Get SECRET_KEY_JWT
	secretKey := os.Getenv("SECRET_KEY_JWT")
	if secretKey == "" {
		fmt.Errorf("Secret key is undefined")
		return false
	}

	// Convert the secret key to []byte
	secretKeyBytes := []byte(secretKey)

	if cookie != nil {
		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				fmt.Println("Unauthorized account.")
			}
			return secretKeyBytes, nil
		})

		if err != nil {
			fmt.Println("Unauthorized: ", err.Error())
			return false
		}

		if token.Valid {
			fmt.Println("Token is valid.")
			return true
		}

	} else {
		return false
	}

	return false
}
