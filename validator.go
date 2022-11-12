package gpc

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/restuwahyu13/go-playground-converter/helpers"
)

// Validation request from struct field
func Validator(s interface{}) (interface{}, error) {
	if reflect.TypeOf(s).Kind().String() != "struct" {
		return nil, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(s).Kind().String())
	} else if res, err := helpers.KeyExist(s); err != nil || res == 0 {
		return nil, fmt.Errorf("validator value can't be empty struct %v", s)
	}

	val := validator.New()
	err := val.Struct(s)

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	if err := en_translations.RegisterDefaultTranslations(val, trans); err != nil {
		return nil, err
	}

	if err == nil {
		return nil, err
	}

	return helpers.FormatError(err, trans, s)
}
