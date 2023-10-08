package utils

import (
	"encoding/json"
	"fmt"
)

func JsonToMap(jsonData string) map[string]interface{} {
	// JSON convert to map
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	// Print data the in map
	fmt.Println("data:", data)

	return data
}
