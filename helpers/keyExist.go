package helper

import (
	"encoding/json"
	"reflect"
)

func KeyExist(input interface{}) (int, error) {
	if reflect.TypeOf(input).Kind() != reflect.Struct {
		return -1, Exception("key_exist_input_invalid", input)
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
