package helper

import (
	"testing"
)

func TestKeyExist(action *testing.T) {
	action.Run("Should be TestKeyExist - key not exist", func(t *testing.T) {
		data := "hello wordl"
		_, err := KeyExist(data)

		if err == nil {
			t.FailNow()
		}
	})

	action.Run("Should be TestKeyExist - key is exist", func(t *testing.T) {
		data := struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}{}

		_, err := KeyExist(data)

		if err != nil {
			t.FailNow()
		}
	})
}
