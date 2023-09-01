package utils

import (
	"log"

	godotenv "github.com/joho/godotenv"
)

func EnvLoader() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

}
