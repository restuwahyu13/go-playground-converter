package gpc

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func keyExist(input interface{}) (int, error) {
	if reflect.TypeOf(input).Kind() != reflect.Struct {
		return -1, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(input).Kind().String())
	}

	received := make(map[string]interface{})

	stringify, err := json.Marshal(input)
	if err != nil {
		return -1, err
	}

	if err := json.Unmarshal(stringify, &received); err != nil {
		return -1, err
	}

	return len(received), nil
}
