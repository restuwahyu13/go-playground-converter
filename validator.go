package gpc

import (
	"fmt"
	"log"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type errorResponse struct {
	Errors interface{} `json:"errors"`
}

// Validation request for struct field
func Validator(s interface{}) errorResponse {
	dataType(s)

	val := validator.New()
	err := val.Struct(s)

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	trans_err := en_translations.RegisterDefaultTranslations(val, trans)

	if trans_err != nil {
		defer log.Fatalf("RegisterDefaultTranslations Error: %v", trans_err)
		panic(fmt.Sprintf("RegisterDefaultTranslations Error: %v", trans_err))
	}

	return bindError(err, trans)
}

/**
 * @description -> binding error response into validator
 */

func bindError(err error, trans ut.Translator) errorResponse {
	errCollection := make(map[string][]map[string]interface{})
	errors := errorResponse{}

	if err == nil {
		return errors
	}

	for _, e := range err.(validator.ValidationErrors) {

		errResult := make(map[string]interface{})
		errResult["param"] = e.StructField()
		errResult["msg"] = e.Translate(trans)
		errResult["tag"] = e.Tag()

		errCollection["errors"] = append(errCollection["errors"], errResult)
		errors.Errors = errCollection["errors"]
	}

	if errors.Errors == nil {
		errors.Errors = errCollection
	}

	return errors
}

func dataType(typeData interface{}) {
	var x1 interface{}
	var x2 struct{}

	switch reflect.TypeOf(typeData) {

	case reflect.TypeOf(0):
		defer log.Fatalf("Validator value not supported, because %v is not struct type:", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	case reflect.TypeOf("hello"):
		defer log.Fatalf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	case reflect.TypeOf([]int{0}):
		defer log.Fatalf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	case reflect.TypeOf([]string{"hello"}):
		defer log.Fatalf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	case reflect.TypeOf(map[string]interface{}{"name": "john doe"}):
		defer log.Fatalf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	case reflect.TypeOf([]map[string]interface{}{{"name": "john doe"}}):
		defer log.Fatalf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	case reflect.TypeOf(x1):
		defer log.Fatalf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	case reflect.TypeOf(x2):
		defer log.Fatalf("Validator value not supported, because %v is not struct type", reflect.TypeOf(typeData))
		panic("Validator value not supported, because is not struct type")

	default:
		return
	}
}
