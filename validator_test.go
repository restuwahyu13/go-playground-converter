package gpc

import (
	"testing"
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
	action.Run("Should be TestValidator - email is not valid", func(t *testing.T) {
		payload := Login{Email: "johndoe@#gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Email must be a valid email address":
				assert(t, v["msg"], "Email must be a valid email address")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator - password must be greater than 7", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Password must be greater than 7 characters in length":
				assert(t, v["msg"], "Password must be greater than 7 characters in length")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with single gpc - email and password is empty", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "", Password: ""}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with single gpc - email and password not valid", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "johndoe#gmail.com", Password: "qwert12"}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with single gpc - email is not empty", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "email tidak boleh kosong":
				assert(t, v["msg"], "email tidak boleh kosong")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with single gpc - password is not empty", func(t *testing.T) {
		payload := LoginSingleGpc{Email: "johndoe@gmail.com", Password: ""}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "password tidak boleh kosong":
				assert(t, v["msg"], "password tidak boleh kosong")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with multi gpc - email and password is empty", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "", Password: ""}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with multi gpc - email and password not valid", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "johndoe#gmail.com", Password: "qwert12"}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with multi gpc - email not valid", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "johndoe#gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "email format tidak valid":
				assert(t, v["msg"], "email format tidak valid")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with multi gpc - password must be greater than 7", func(t *testing.T) {
		payload := LoginMultiGpc{Email: "johndoe@gmail.com", Password: "qwerty"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "password harus lebih besar dari 7":
				assert(t, v["msg"], "password harus lebih besar dari 7")

			default:
				t.FailNow()
			}
		}
	})
}
