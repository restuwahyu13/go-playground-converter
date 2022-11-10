package gpc

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validation request from struct field
func Validator(s interface{}) interface{} {
	if reflect.TypeOf(s).Kind().String() != "struct" {
		panic(fmt.Errorf("Validator value not supported, because %v is not struct", reflect.TypeOf(s).Kind().String()))
	} else if reflect.ValueOf(s).IsZero() {
		panic(fmt.Errorf("Validator value can't be empty struct %v", s))
	}

	val := validator.New()
	err := val.Struct(s)

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	if err := en_translations.RegisterDefaultTranslations(val, trans); err != nil {
		panic(err)
	}

	if err == nil {
		return nil
	}

	return bindError(err, trans)
}

// binding error response into validator
func bindError(err error, trans ut.Translator) interface{} {
	errRes := make(map[string][]map[string]interface{})

	for _, e := range err.(validator.ValidationErrors) {
		errResult := make(map[string]interface{})
		errResult["param"] = e.StructField()
		errResult["msg"] = e.Translate(trans)
		errResult["tag"] = e.ActualTag()
		errRes["errors"] = append(errRes["errors"], errResult)
	}

	return errRes
}
