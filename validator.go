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
	dataType(s)

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

func dataType(typeData interface{}) {
	var x1 interface{}
	var x2 struct{}

	switch reflect.TypeOf(typeData) {
	case reflect.TypeOf(0):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type:", reflect.TypeOf(typeData)))

	case reflect.TypeOf("hello"):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData)))

	case reflect.TypeOf([]int{0}):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData)))

	case reflect.TypeOf([]string{"hello"}):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData)))

	case reflect.TypeOf(map[string]interface{}{"name": "john doe"}):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData)))

	case reflect.TypeOf([]map[string]interface{}{{"name": "john doe"}}):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData)))

	case reflect.TypeOf(x1):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData)))

	case reflect.TypeOf(x2):
		panic(fmt.Errorf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData)))
	}
}
