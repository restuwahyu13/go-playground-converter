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
	action.Run("Should be TestValidator - error is not empty", func(t *testing.T) {
		var (
			payload  Login = Login{Email: "johndoe@#gmail.com", Password: "qwerty12"}
			res, err       = Validator(payload)
		)

		if err != nil {
			t.FailNow()
		}

		if res == nil {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator - error is empty", func(t *testing.T) {
		var (
			payload  Login = Login{Email: "johndoe@gmail.com", Password: "qwerty12"}
			res, err       = Validator(payload)
		)

		if err != nil {
			t.FailNow()
		}

		if res != nil {
			t.FailNow()
		}
	})
}
