package utils

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func Generator() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error generate uuid")
	}

	return id
}
