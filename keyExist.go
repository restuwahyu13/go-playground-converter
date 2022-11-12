package gpc

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func keyExist(input interface{}) (int, error) {

	if reflect.TypeOf(input).Kind().String() != "struct" {
		return -1, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(input).Kind().String())
	}

	received := make(map[string]interface{})
	mapsArr := []string{}

	stringify, err := json.Marshal(&input)

	if err != nil {
		return -1, err
	}

	if err := json.Unmarshal(stringify, &received); err != nil {
		return -1, err
	}

	for i, _ := range received {
		mapsArr = append(mapsArr, i)
	}

	return len(mapsArr), nil
}
