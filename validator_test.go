package gpc

import (
	"sync"
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
	action.Run("Should be TestValidator - with error", func(t *testing.T) {
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

	action.Run("Should be TestValidator - without error", func(t *testing.T) {
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

	action.Run("Should be TestValidator - with error use gorutine", func(t *testing.T) {
		var (
			wg         *sync.WaitGroup   = new(sync.WaitGroup)
			errorsChan chan *FormatError = make(chan *FormatError)
		)

		wg.Add(1)

		go func() {
			wg.Done()

			res, err := Validator(Login{Email: "johndoe@#gmail.com", Password: "qwerty12"})
			if err != nil {
				t.FailNow()
			}

			errorsChan <- res
		}()

		wg.Wait()
		errors := <-errorsChan

		if len(errors.Errors) < 1 {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator - large validation", func(t *testing.T) {

		for i := 0; i < 100000; i++ {

			res, err := Validator(Login{Email: "johndoe@#gmail.com", Password: "qwerty12"})

			if err != nil {
				t.FailNow()
			}

			if res == nil {
				t.FailNow()
			}
		}
	})
}
