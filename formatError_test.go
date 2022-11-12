package gpc

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFormatError(action *testing.T) {
	action.Run("Should be TestFormatError - value is not struct", func(t *testing.T) {
		payload := "hello world"

		_, err := Validator(payload)
		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("validator value not supported, because %v is not struct", reflect.TypeOf(payload).Kind().String()))
		}
	})

	action.Run("Should be TestFormatError - is empty struct", func(t *testing.T) {
		payload := struct{}{}
		_, err := Validator(payload)

		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("validator value can't be empty struct %v", payload))
		}
	})

	action.Run("Should be TestFormatError - not empty value", func(t *testing.T) {
		payload := Login{Email: "", Password: ""}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Email is a required field":
				assert(t, v["msg"], "Email is a required field")

			case "Password is a required field":
				assert(t, v["msg"], "Password is a required field")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestFormatError - success", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)
		assert(t, res, nil)
	})
}
