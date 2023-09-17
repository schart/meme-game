package jwt

import (
	"fmt"
	"os"

	utils "shared-library/utils"

	"github.com/dgrijalva/jwt-go"
)

func TokenCreate(username string) (string, error) {
	// Load env(s).
	utils.EnvLoader()

	// Call secret_key
	secret_key := os.Getenv("SECRET_KEY_JWT")
	if secret_key == "" {
		return "", fmt.Errorf("Error when call secret key")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["expire"] = int64(30.44 * 24 * 3600)

	// Validate token
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", fmt.Errorf("Error when create:", err)
	}

	fmt.Println("token created:", tokenString)
	return tokenString, nil
}
