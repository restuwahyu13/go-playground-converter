package validate

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	helper "github.com/restuwahyu13/go-playground-converter/helpers"
)

type Errors struct {
	Msg   string `json:"msg"`
	Param string `json:"param"`
	Tag   string `json:"tag"`
}

func check(s interface{}, v *validator.Validate) error {
	typeof := reflect.TypeOf(s)

	if typeof.Kind() == reflect.Pointer {
		return helper.Exception("err_format_input_not_pointer", typeof.Kind().String())
	} else {
		if typeof.Kind() != reflect.Struct {
			return helper.Exception("err_format_input_not_struct", typeof.Kind().String())
		}
	}

	if res, err := helper.KeyExist(s); err != nil || res == 0 {
		return helper.Exception("err_format_input_empty_struct", s)
	}

	return nil
}

func format(err error, trans ut.Translator) ([]Errors, error) {
	errResult := Errors{}
	errResults := make([]Errors, 0, len(err.(validator.ValidationErrors)))

	if !reflect.ValueOf(err).CanConvert(reflect.TypeOf(err)) {
		return nil, helper.Exception("err_format_not_err_validator", nil)
	}

	for _, e := range err.(validator.ValidationErrors) {
		errResult.Param = e.StructField()
		errResult.Msg = e.Translate(trans)
		errResult.Tag = e.ActualTag()

		errResults = append(errResults, errResult)
	}

	return errResults, nil
}

func translator(err error, s interface{}, v *validator.Validate) ([]Errors, error) {
	translatorEng := en.New()
	universalTranslator := ut.New(translatorEng, translatorEng)

	translator, found := universalTranslator.GetTranslator("en")
	if !found {
		return nil, helper.Exception("err_format_transalator_not_found", nil)
	}

	if err := en_translations.RegisterDefaultTranslations(v, translator); err != nil {
		return nil, err
	}

	return format(err, translator)
}
