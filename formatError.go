package gpc

import (
	"reflect"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func formatError(err error, trans ut.Translator, customMessage interface{}) (interface{}, error) {
	errRes := make(map[string][]map[string]interface{})
	tags := []string{}

	for i, e := range err.(validator.ValidationErrors) {

		errResult := make(map[string]interface{})
		errResult["param"] = e.StructField()

		if _, ok := reflect.TypeOf(customMessage).Field(i).Tag.Lookup("gpc"); !ok {
			errResult["msg"] = e.Translate(trans)
		} else {
			strucField, _ := reflect.TypeOf(customMessage).FieldByName(e.StructField())
			structTags := strucField.Tag.Get("gpc")

			regexTag := regexp.MustCompile(`=+[\w].*`)
			regexVal := regexp.MustCompile(`[\w]+=`)
			strArr := strings.Split(structTags, ",")
			tags = append(tags, mergeSlice(strArr)...)

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
