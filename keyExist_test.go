package gpc

import "testing"

func TestKeyExist(action *testing.T) {
	action.Run("Should be TestKeyExist - key not exist", func(t *testing.T) {
		data := struct{}{}
		res, err := keyExist(data)
		if err != nil {
			assert(t, res, -1)
		}
	})

	action.Run("Should be TestKeyExist - key is exist", func(t *testing.T) {
		data := struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}{}

		res, err := keyExist(data)
		if err != nil {
			t.FailNow()
		}

		assert(t, res, 2)
	})
}
