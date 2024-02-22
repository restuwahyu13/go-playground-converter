package gpc

import (
	"testing"
)

func TestFormatError(action *testing.T) {
	action.Run("Should be TestFormatError - value is not struct", func(t *testing.T) {
		var (
			payload string = "hello world"
			_, err         = Validator(payload)
		)

		if err == nil {
			t.FailNow()
		}
	})

	action.Run("Should be TestFormatError - is empty struct", func(t *testing.T) {
		var (
			payload struct{} = struct{}{}
			_, err           = Validator(payload)
		)

		if err == nil {
			t.FailNow()
		}
	})

	action.Run("Should be TestFormatError - not empty value", func(t *testing.T) {
		var (
			payload  = Login{Email: "", Password: ""}
			res, err = Validator(payload)
		)

		if err != nil {
			t.FailNow()
		}

		if res == nil {
			t.FailNow()
		}
	})

	action.Run("Should be TestFormatError - success", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty12"}
		res, err := Validator(payload)

		if err != nil {
			t.FailNow()
		}

		if res != nil {
			t.FailNow()
		}
	})
}
