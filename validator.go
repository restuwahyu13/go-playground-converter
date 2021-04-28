package gpc

import (
	"github.com/go-playground/validator/v10"
)

type response struct {
    Message string `json:"msg"`
    Param string `json:"param"`
    Value interface{} `json:"value"`
}

type Validators interface {
    Validator(s interface{}) map[string]interface{}
}

type validators struct {
    validate *validator.Validate
}

func NewValidator(validate *validator.Validate) *validators{
    return &validators{validate: validate}
}

func (v *validators) Validator(s *interface{}) map[string]interface{} {

    v.validate = validator.New()

    err := v.validate.Struct(s)
    errObject := make(map[string]interface{})

    if err != nil {
        for _, errResult := range err.(validator.ValidationErrors) {
            switch errResult.ActualTag() {
            case "email":
            errObject[errResult.StructField()] = response{
                Message: "email is not valid",
                Param: errResult.StructField(),
                Value: errResult.Value(),
                }
            case "required":
            errObject[errResult.StructField()] = response{
                Message: "is required",
                Value: errResult.Value(),
                Param: errResult.StructField(),
            }
        }
    }
}
    return errObject
}