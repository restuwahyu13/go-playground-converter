package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/restuwahyu13/go-playground-converter/helpers"
)

type Messages struct {
	Param string `json:"param"`
	Tag   string `json:"tag"`
	Msg   string `json:"msg"`
}

type ErrorsResponse struct {
	Errors []Messages
}

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

	return formatError(err, trans, s)
}

func formatError(err error, trans ut.Translator, customMessage interface{}) (interface{}, error) {
	errRes := make(map[string][]map[string]interface{})
	tags := []string{}

	for i, e := range err.(validator.ValidationErrors) {

		errResult := make(map[string]interface{})
		errResult["param"] = e.StructField()

		// x, _ := reflect.TypeOf(customMessage).FieldByName(e.StructField())
		// d := x.Tag.Get("gpc")

		// fmt.Println(d)

		if _, ok := reflect.TypeOf(customMessage).Field(i).Tag.Lookup("gpc"); !ok {
			errResult["msg"] = e.Translate(trans)
		} else {
			// structTags := reflect.TypeOf(customMessage).Field(i).Tag.Get("gpc")
			strucField, _ := reflect.TypeOf(customMessage).FieldByName(e.StructField())
			structTags := strucField.Tag.Get("gpc")

			regexTag := regexp.MustCompile(`=+[\w].*`)
			regexVal := regexp.MustCompile(`[\w]+=`)
			strArr := strings.Split(structTags, ",")
			tags = append(tags, helpers.MergeSlice(strArr)...)

			for j, v := range tags {
				if ok := regexTag.ReplaceAllString(tags[j], ""); ok == e.ActualTag() {
					errResult["msg"] = regexVal.ReplaceAllString(v, "")
					tags = append(tags, "")
				}
			}
		}

		errResult["tag"] = e.ActualTag()
		errRes["errors"] = append(errRes["errors"], errResult)
	}

	return errRes, nil
}
