package gpc

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type response struct {
    Message interface{} `json:"message"`
    Value interface{} `json:"value"`
    Param string `json:"param"`
    Tag string `json:"tag"`
}

type validators interface {
    validator(s interface{}, config map[int]interface{}) map[string]interface{}
}

type validate struct {
    v2 *validator.Validate
}

func NewValidator(v2 *validator.Validate) *validate {
    return &validate{v2: v2}
}

func (v *validate) validator(s interface{}, config map[int]interface{}) map[string]interface{} {

errObject := make(map[string]interface{})

for _, value := range config {
    encode, _ := json.Marshal(value)
    json.Unmarshal(encode, &value)
    mapping := value.(map[string]interface{})

    v.v2 = validator.New()
    err := v.v2.Struct(s)

    if err != nil {
        for _, errResult := range err.(validator.ValidationErrors) {
            switch {
            case mapping["Tag"] == errResult.ActualTag() && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                    }
                }
            }
        }
    }
    return errObject
}