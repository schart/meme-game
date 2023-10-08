package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func ParameterChecker(body url.Values, structure interface{}) error {
	structValue := reflect.ValueOf(structure)
	structType := structValue.Type()

	// compare structure fields with expect key of data in body
	for i := 0; i < structType.NumField(); i++ {

		fieldName := structType.Field(i).Name
		body := strings.TrimSpace(body.Get(fieldName))
		if body == "" {
			return fmt.Errorf("Error: Need field named the %v", fieldName)
		}
	}

	return nil
}
