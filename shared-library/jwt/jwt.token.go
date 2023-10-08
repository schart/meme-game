package jwt

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"shared-library/utils"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(accountId int) (string, error) {
	// ENV Loader
	utils.EnvLoader()

	// Get SECRET_KEY_JWT
	secretKey := os.Getenv("SECRET_KEY_JWT")
	if secretKey == "" {
		return "", fmt.Errorf("Error when calling the secret key")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["accountId"] = accountId
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println("Token created: ", tokenStr)
	return tokenStr, nil
}

func DecodeJWT(token *http.Cookie) (jwt.MapClaims, error) {
	// ENV Loader
	utils.EnvLoader()

	// Get SECRET_KEY_JWT
	secretKey := os.Getenv("SECRET_KEY_JWT")
	if secretKey == "" {
		return nil, fmt.Errorf("Error when calling the secret key")
	}

	if token != nil {
		token, err := jwt.Parse(token.Value, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unauthorized token")
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			return nil, fmt.Errorf("Unauthorized token")

		}

		if token.Valid {
			claims, _ := token.Claims.(jwt.MapClaims)
			return claims, nil
		}

	} else {
		return nil, fmt.Errorf("Unauthorized Account")

	}

	return nil, nil
}
