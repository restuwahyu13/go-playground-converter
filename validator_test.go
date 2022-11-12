package gpc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
)

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gt=7"`
}

type LoginSingleGpc struct {
	Email    string `validate:"required,email" gpc:"required=email tidak boleh kosong"`
	Password string `validate:"required,gt=7" gpc:"required=password tidak boleh kosong"`
}

type LoginMultiGpc struct {
	Email    string `validate:"required,email" gpc:"required=email tidak boleh kosong,email=email format tidak valid"`
	Password string `validate:"required,gt=7" gpc:"required=password tidak boleh kosong,gt=password harus lebih besar dari 7"`
}

func TestValidator(action *testing.T) {
	action.Run("Should be validator - value is not struct", func(t *testing.T) {
		payload := "hello world"

		_, err := Validator(payload)
		if err != nil {
			assert.Equal(t, err.Error(), fmt.Sprintf("validator value not supported, because %v is not struct", reflect.TypeOf(payload).Kind().String()))
		}
	})

	action.Run("Should be validator - is empty struct", func(t *testing.T) {
		payload := struct{}{}
		_, err := Validator(payload)

		if err != nil {
			assert.Equal(t, err.Error(), fmt.Sprintf("validator value can't be empty struct %v", payload))
		}
	})

	action.Run("Should be validator - is not empty value", func(t *testing.T) {

		payload := Login{Email: "", Password: ""}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
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
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Email must be a valid email address":
				assert.Equal(t, v["msg"], "Email must be a valid email address")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator - password must be greater than 7", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Password must be greater than 7 characters in length":
				assert.Equal(t, v["msg"], "Password must be greater than 7 characters in length")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator with single gpc - email and password is empty", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "", Password: ""}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert.Equal(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be validator with single gpc - email and password not valid", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "johndoe#gmail.com", Password: "qwert12"}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert.Equal(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be validator with single gpc - email is not empty", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "email tidak boleh kosong":
				assert.Equal(t, v["msg"], "email tidak boleh kosong")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator with single gpc - password is not empty", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "johndoe@gmail.com", Password: ""}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "password tidak boleh kosong":
				assert.Equal(t, v["msg"], "password tidak boleh kosong")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator with multi gpc - email and password is empty", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "", Password: ""}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert.Equal(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be validator with multi gpc - email and password not valid", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "johndoe#gmail.com", Password: "qwert12"}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert.Equal(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be validator with multi gpc - email not valid", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "johndoe#gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "email format tidak valid":
				assert.Equal(t, v["msg"], "email format tidak valid")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator with multi gpc - password must be greater than 7", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "johndoe@gmail.com", Password: "qwerty"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "password harus lebih besar dari 7":
				assert.Equal(t, v["msg"], "password harus lebih besar dari 7")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be validator - success", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)
		assert.Equal(t, res, nil)
	})
}
