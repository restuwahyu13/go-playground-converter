package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func TestValidator(action *testing.T) {
	action.Run("Should be validator - is not empty value", func(t *testing.T) {

		payload := Login{Email: "", Password: ""}
		err := Validator(payload)

		for _, v := range err.Errors.([]map[string]interface{}) {

			switch v["msg"] {

			case "Email is a required field":
				assert.Equal(t, v["msg"], "Email is a required field")

			case "Password is a required field":
				assert.Equal(t, v["msg"], "Password is a required field")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator - email is not valid", func(t *testing.T) {

		payload := Login{Email: "johndoe@#gmail.com", Password: "qwerty12"}
		err := Validator(payload)

		for _, v := range err.Errors.([]map[string]interface{}) {

			switch v["msg"] {

			case "Email must be a valid email address":
				assert.Equal(t, v["msg"], "Email must be a valid email address")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator - success", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty12"}
		err := Validator(payload)

		assert.Equal(t, err.Errors, nil)
	})
}
