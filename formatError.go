package gpc

import (
	"reflect"
	"regexp"
	"strings"
	"sync"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type FormatError struct {
	Errors []FormatErrorMetadata
}

type FormatErrorMetadata struct {
	Msg   string `json:"msg"`
	Param string `json:"param"`
	Tag   string `json:"tag"`
}

func formatError(err error, trans ut.Translator, customMessage interface{}) (*FormatError, error) {
	var (
		mutex       *sync.RWMutex = new(sync.RWMutex)
		errResponse *FormatError  = new(FormatError)
	)

	for i, e := range err.(validator.ValidationErrors) {
		errResult := new(FormatErrorMetadata)
		errResult.Param = e.StructField()

		if _, ok := reflect.TypeOf(customMessage).Field(i).Tag.Lookup("gpc"); !ok {
			errResult.Msg = e.Translate(trans)
		} else {
			strucField, _ := reflect.TypeOf(customMessage).FieldByName(e.StructField())
			structTags := strucField.Tag.Get("gpc")

			regexTag := regexp.MustCompile(`=+[\w].*`)
			regexVal := regexp.MustCompile(`[\w]+=`)
			tags := strings.Split(structTags, ",")

			for j, v := range tags {
				replacedTag := regexTag.ReplaceAllString(tags[j], "")
				if replacedTag == e.ActualTag() {
					errResult.Msg = regexVal.ReplaceAllString(v, "")
				}
			}
		}

		mutex.RLock()
		defer mutex.RUnlock()

		errResult.Tag = e.ActualTag()
		errResponse.Errors = append(errResponse.Errors, *errResult)
	}

	return errResponse, nil
}

// func formatError(err error, trans ut.Translator, customMessage interface{}) (*FormatError, error) {
// 	errs := err.(validator.ValidationErrors)
// 	errRes := &FormatError{
// 		Errors: make([]FormatErrorMetadata, 0, len(errs)),
// 	}

// 	regexTag := regexp.MustCompile(`=+[\w].*`)
// 	regexVal := regexp.MustCompile(`[\w]+=`)

// 	for _, e := range errs {
// 		errResult := FormatErrorMetadata{
// 			Param: e.StructField(),
// 		}

// 		strucField, ok := reflect.TypeOf(customMessage).FieldByName(e.StructField())
// 		if !ok {
// 			errResult.Msg = e.Translate(trans)
// 		} else {
// 			structTags := strucField.Tag.Get("gpc")
// 			tags := strings.Split(structTags, ",")

// 			for _, v := range tags {
// 				replacedTag := regexTag.ReplaceAllString(v, "")
// 				if replacedTag == e.ActualTag() {
// 					errResult.Msg = regexVal.ReplaceAllString(v, "")
// 					break
// 				}
// 			}
// 		}

// 		errResult.Tag = e.ActualTag()
// 		errRes.Errors = append(errRes.Errors, errResult)
// 	}

// 	fmt.Println(errRes)

// 	return errRes, nil
// }
